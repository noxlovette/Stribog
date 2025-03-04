use crate::api::public;
use crate::db::init::AppState;
use axum::routing::get;
use axum::Router;

pub fn public_routes() -> Router<AppState> {
    Router::new()
        .route(
            "/{api_key}/{forge_id}/{spark_id}",
            get(public::fetch_public_spark)
        )
}
