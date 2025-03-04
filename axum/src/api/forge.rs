use axum::{extract::{Path, State}, Json};
use hyper::StatusCode;
use crate::{auth::jwt::Claims, db::init::AppState, models::forges::{ForgeBody, ForgeCreateBody, ForgeUpdate}};
use super::error::APIError;


pub async fn fetch_forge(
    State(state): State<AppState>,
    Path(id): Path<String>,
    claims: Claims,
) -> Result<Json<ForgeBody>, APIError> {
    let forge = sqlx::query_as!(
        ForgeBody,
        r#"
        SELECT * FROM forges
        WHERE id = $1 AND owner_id = $2
        "#,
        id,
        claims.sub
    )
    .fetch_one(&state.db)
    .await?;

    Ok(Json(forge))
}


pub async fn list_forge(
    State(state): State<AppState>,
    claims: Claims,
) -> Result<Json<Vec<ForgeBody>>, APIError> {
    let forge = sqlx::query_as!(
        ForgeBody,
        r#"
        SELECT * FROM forges
        WHERE owner_id = $1
        "#,
        claims.sub
    )
    .fetch_all(&state.db)
    .await?;

    Ok(Json(forge))
}

pub async fn create_forge(
    State(state): State<AppState>,
    claims: Claims,
    Json(payload): Json<ForgeCreateBody>,
) -> Result<StatusCode, APIError> {

    let _forge = sqlx::query!(
        r#"
        INSERT INTO forges (id, title, description, owner_id) 
         VALUES ($1, $2, $3, $4)
         ON CONFLICT DO NOTHING
         "#,
        nanoid::nanoid!(),
        payload.title,
        payload.description,
        claims.sub,
    )
    .execute(&state.db)
    .await?;

    Ok(StatusCode::CREATED)
}

pub async fn delete_forge(
    State(state): State<AppState>,
    claims: Claims,
    Path(id): Path<String>
) -> Result<StatusCode, APIError> {
    let _forge = sqlx::query!(
        r#"
        DELETE FROM forges WHERE id = $1 AND owner_id = $2
         "#,
        id,
        claims.sub,
    )
    .execute(&state.db)
    .await?;

    Ok(StatusCode::NO_CONTENT)
}

pub async fn update_forge(
    State(state): State<AppState>,
    Path(id): Path<String>,
    claims: Claims,
    Json(payload): Json<ForgeUpdate>,
) -> Result<StatusCode, APIError> {
    let _forge = sqlx::query!(
        r#"
        UPDATE forges 
         SET 
            title = COALESCE($1, title),
            description = COALESCE($2, description)
         WHERE id = $3 AND owner_id = $4
         "#,
        payload.title,
        payload.description,
        id,
        claims.sub
    )
    .execute(&state.db)
    .await?;

    Ok(StatusCode::OK)
}