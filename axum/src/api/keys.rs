use axum::{extract::{Path, State}, Json};
use hyper::StatusCode;
use crate::{auth::jwt::Claims, db::init::AppState, models::api_keys::{KeyBody, KeyCreateBody}};
use super::error::APIError;
use nanoid::nanoid;

// Create a new API key for a forge
pub async fn create_api_key(
    State(state): State<AppState>,
    claims: Claims,
    Path(forge_id): Path<String>,
    Json(payload): Json<KeyCreateBody>,
) -> Result<Json<KeyBody>, APIError> {
    // First verify the user owns the forge
    let forge = sqlx::query!(
        r#"
        SELECT id FROM forges
        WHERE id = $1 AND owner_id = $2
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if forge.is_none() {
        return Err(APIError::NotFound("Forge not found or unauthorized".into()));
    }
    
    // Generate API key with nanoid
    let key_id = nanoid!();
    
    let new_key = sqlx::query_as!(
        KeyBody,
        r#"
        INSERT INTO api_keys (id, forge_id, title, is_active)
        VALUES ($1, $2, $3, true)
        RETURNING id, forge_id, title, is_active, created_at, last_used_at
        "#,
        key_id,
        forge_id,
        payload.title
    )
    .fetch_one(&state.db)
    .await?;
    
    Ok(Json(new_key))
}

// List all API keys for a forge
pub async fn list_api_keys(
    State(state): State<AppState>,
    claims: Claims,
    Path(forge_id): Path<String>,
) -> Result<Json<Vec<KeyBody>>, APIError> {
    // Verify the user owns the forge
    let forge = sqlx::query!(
        r#"
        SELECT id FROM forges
        WHERE id = $1 AND owner_id = $2
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if forge.is_none() {
        return Err(APIError::NotFound("Forge not found or unauthorized".into()));
    }
    
    let keys = sqlx::query_as!(
        KeyBody,
        r#"
        SELECT id, forge_id, title, is_active, created_at, last_used_at
        FROM api_keys
        WHERE forge_id = $1
        "#,
        forge_id
    )
    .fetch_all(&state.db)
    .await?;
    
    Ok(Json(keys))
}

// Delete an API key
pub async fn delete_api_key(
    State(state): State<AppState>,
    claims: Claims,
    Path((forge_id, key_id)): Path<(String, String)>,
) -> Result<StatusCode, APIError> {
    // Verify the user owns the forge
    let forge = sqlx::query!(
        r#"
        SELECT id FROM forges
        WHERE id = $1 AND owner_id = $2
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if forge.is_none() {
        return Err(APIError::NotFound("Forge not found or unauthorized".into()));
    }
    
    sqlx::query!(
        r#"
        DELETE FROM api_keys 
        WHERE id = $1 AND forge_id = $2
        "#,
        key_id,
        forge_id
    )
    .execute(&state.db)
    .await?;
    
    Ok(StatusCode::NO_CONTENT)
}

// Toggle activation status
pub async fn toggle_api_key(
    State(state): State<AppState>,
    claims: Claims,
    Path((forge_id, key_id)): Path<(String, String)>,
) -> Result<Json<KeyBody>, APIError> {
    // Verify the user owns the forge
    let forge = sqlx::query!(
        r#"
        SELECT id FROM forges
        WHERE id = $1 AND owner_id = $2
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if forge.is_none() {
        return Err(APIError::NotFound("Forge not found or unauthorized".into()));
    }
    
    let updated_key = sqlx::query_as!(
        KeyBody,
        r#"
        UPDATE api_keys
        SET is_active = NOT is_active
        WHERE id = $1 AND forge_id = $2
        RETURNING id, forge_id, title, is_active, created_at, last_used_at
        "#,
        key_id,
        forge_id
    )
    .fetch_one(&state.db)
    .await?;
    
    Ok(Json(updated_key))
}