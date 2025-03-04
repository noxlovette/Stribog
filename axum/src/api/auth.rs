use crate::auth::error::AuthError;
use crate::auth::helpers::verify_password;
use crate::auth::helpers::{generate_refresh_token, generate_token, hash_password};
use crate::auth::jwt::RefreshClaims;
use crate::db::init::AppState;
use crate::api::error::APIError;
use crate::models::users::{
    AuthBody, AuthPayload, SignUpBody, SignUpPayload, User
};
use axum::extract::Json;
use axum::extract::State;
use axum::response::Response;
use nanoid::nanoid;
use validator::Validate;

pub async fn signup(
    State(state): State<AppState>,
    Json(payload): Json<SignUpPayload>,
) -> Result<Json<SignUpBody>, APIError> {
    tracing::info!("Creating user");
    if payload.username.is_empty() || payload.pass.is_empty() {
        return Err(APIError::InvalidCredentials);
    }
    payload.validate().map_err(|e| {
        eprintln!("{:?}", e);
        APIError::InvalidCredentials
    })?;
    

    let SignUpPayload {
        name,
        username,
        email,
        pass,
    } = payload;

    let db = &state.db;
    let hashed_password = hash_password(&pass)?;
    let id = nanoid!();

    sqlx::query!(
        r#"
            INSERT INTO users (name, username, email, pass, verified, id)
            VALUES ($1, $2, $3, $4, false, $5)
        "#,
        name,
        username,
        email,
        hashed_password,
        id
    )
    .execute(db)
    .await
    .map_err(|e| match e {
        sqlx::Error::Database(dbe) if dbe.constraint() == Some("user_username_key") => {
            APIError::AlreadyExists("Username already taken".into())
        }
        sqlx::Error::Database(dbe) if dbe.constraint() == Some("user_email_key") => {
            APIError::AlreadyExists("Email already registered".into())
        }
        _ => APIError::Database(e),
    })?;

    Ok(Json(SignUpBody { id }))
}

pub async fn authorize(
    State(state): State<AppState>,
    Json(payload): Json<AuthPayload>,
) -> Result<Response, APIError> {

    if payload.username.is_empty() || payload.pass.is_empty() {
        return Err(APIError::InvalidCredentials);
    }
    payload.validate().map_err(|e| {
        eprintln!("{:?}", e);
        APIError::InvalidCredentials
    })?;

    let user = sqlx::query_as!(
        User,
        r#"
        SELECT username, email, id, name, pass, verified
        FROM users
        WHERE username = $1
        "#,
        payload.username
    )
    .fetch_optional(&state.db)
    .await
    .map_err(|e| {
        eprintln!("{:?}", e);
        AuthError::WrongCredentials
    })?
    .ok_or_else(|| APIError::NotFound("User not found".into()))?;

    if !verify_password(&user.pass, &payload.pass)? {
        return Err(APIError::AuthenticationFailed);
    }

    let token = generate_token(&user)?;
    let refresh_token = generate_refresh_token(&user)?;

    Ok(AuthBody::into_response(token, refresh_token))
}

pub async fn refresh(
    State(state): State<AppState>,
    claims: RefreshClaims,
) -> Result<Response, APIError> {
    let user = sqlx::query_as!(
        User,
        r#"
        SELECT username, email, id, name, pass, verified 
        FROM users
        WHERE id = $1
        "#,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await
    .map_err(|e| {
        eprintln!("{:?}", e);
        APIError::NotFound("User not Found".into())
    })?
    .ok_or(AuthError::UserNotFound)?;

    let token = generate_token(&user)?;
    Ok(AuthBody::into_refresh(token))
}
