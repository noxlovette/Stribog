use axum::{
    http::{HeaderName, HeaderValue, Method, Request, StatusCode},
    response::IntoResponse,
    Router,
};
use stribog::db::init::{init_db, AppState};
use stribog::tools::logging::init_logging;
use stribog::tools::middleware::api_key::validate_api_key;
use std::{env, sync::Arc};
use std::time::Duration;
use tower::ServiceBuilder;
use tower_governor::{governor::GovernorConfigBuilder, GovernorLayer};
use tower_http::{
    cors::{Any, CorsLayer},
    limit::RequestBodyLimitLayer,
    request_id::{MakeRequestUuid, PropagateRequestIdLayer, SetRequestIdLayer},
    timeout::TimeoutLayer,
    trace::TraceLayer,
};
use tracing::{error, info, info_span};

const REQUEST_ID_HEADER: &str = "x-request-id";

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Initialize logging
    init_logging().await;
    
    // Read environment variables
    let internal_cors_origins = env::var("INTERNAL_CORS_ORIGINS")
        .unwrap_or_else(|_| "http://localhost:5173".to_string());
    
    let public_cors_origins = env::var("PUBLIC_CORS_ORIGINS")
        .unwrap_or_else(|_| "*".to_string());
    
    let rate_limit_requests = env::var("RATE_LIMIT_REQUESTS")
        .unwrap_or_else(|_| "100".to_string())
        .parse::<u64>()
        .unwrap_or(100);
    
    let rate_limit_period_secs = env::var("RATE_LIMIT_PERIOD_SECS")
        .unwrap_or_else(|_| "60".to_string())
        .parse::<u64>()
        .unwrap_or(60);
    
    let timeout_secs = env::var("REQUEST_TIMEOUT_SECS")
        .unwrap_or_else(|_| "10".to_string())
        .parse::<u64>()
        .unwrap_or(10);
    
    info!("Starting server with configuration:");
    info!("Internal CORS origins: {}", internal_cors_origins);
    info!("Public CORS origins: {}", public_cors_origins);
    info!("Rate limit: {} requests per {} seconds", rate_limit_requests, rate_limit_period_secs);
    info!("Request timeout: {} seconds", timeout_secs);
    
    // Initialize database connection
    let state = AppState {
        db: init_db().await?,
    };

    // Parse CORS origins
    let internal_origins = parse_cors_origins(&internal_cors_origins);
    let public_origins = parse_cors_origins(&public_cors_origins);
    

    let governor_conf = Arc::new(
        GovernorConfigBuilder::default()
            .per_second(rate_limit_period_secs)
            .burst_size(rate_limit_requests as u32)
            .finish()
            .unwrap()
    );
    
    // Internal routes with API key authentication
    let internal_routes = Router::new()
        .nest("/lesson", stribog::routes::internal::spark_routes::spark_routes())
        .nest("/forge", stribog::routes::internal::forge_routes::forge_routes())
        .nest("/user", stribog::routes::internal::user_routes::user_routes())
        .nest("/auth", stribog::routes::internal::auth_routes::auth_routes())
        .layer(axum::middleware::from_fn(validate_api_key))
        .layer(
            CorsLayer::new()
                .allow_origin(internal_origins)
                .allow_methods([
                    Method::GET,
                    Method::POST,
                    Method::DELETE,
                    Method::PUT,
                    Method::PATCH,
                ])
                .allow_headers(Any),
        );

    // Public routes with rate limiting
    let public_routes = Router::new()
        .nest("/cdn", stribog::routes::external::spark::public_routes())
        .layer(GovernorLayer { config: governor_conf })
        .layer(
            CorsLayer::new()
                .allow_origin(public_origins)
                .allow_methods([Method::GET])
                .allow_headers(Any),
        );

    // Combine routes and apply middleware
    let app = Router::new()
        .merge(public_routes)
        .merge(internal_routes)
        .fallback(handler_404)
        .with_state(state)
        .layer(
            ServiceBuilder::new()
                .layer(SetRequestIdLayer::new(
                    HeaderName::from_static(REQUEST_ID_HEADER),
                    MakeRequestUuid,
                ))
                .layer(
                    TraceLayer::new_for_http().make_span_with(|request: &Request<_>| {
                        let request_id = request.headers().get(REQUEST_ID_HEADER);
                        match request_id {
                            Some(request_id) => info_span!(
                                "http_request",
                                request_id = ?request_id
                            ),
                            None => {
                                error!("could not extract request_id");
                                info_span!("http_request")
                            }
                        }
                    }),
                )
                .layer(PropagateRequestIdLayer::new(HeaderName::from_static(
                    REQUEST_ID_HEADER,
                )))
                .layer(RequestBodyLimitLayer::new(1024 * 1024 * 50)) // 50MB limit
                .layer(TimeoutLayer::new(Duration::from_secs(timeout_secs))),
        );

    // Start server
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await?;
    info!("Server listening on {}", listener.local_addr()?);
    axum::serve(listener, app).await?;
    Ok(())
}

// Helper function to parse CORS origins from environment variable
fn parse_cors_origins(origins_str: &str) -> tower_http::cors::AllowOrigin {
    if origins_str == "*" {
        return tower_http::cors::AllowOrigin::any();
    }

    let origins: Vec<HeaderValue> = origins_str
        .split(',')
        .filter_map(|origin| origin.trim().parse().ok())
        .collect();

    if origins.is_empty() {
        tower_http::cors::AllowOrigin::any()
    } else {
        tower_http::cors::AllowOrigin::list(origins)
    }
}

async fn handler_404() -> impl IntoResponse {
    (StatusCode::NOT_FOUND, "Nothing to see here")
}