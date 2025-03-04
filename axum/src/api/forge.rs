use axum::{extract::{Path, State}, Json};
use hyper::StatusCode;
use crate::{auth::jwt::Claims, db::init::AppState, models::forges::{ForgeBody, ForgeCreateBody, ForgeUpdate, ForgeAccessBody, ForgeAccessCreateBody, ForgeAccessRole}};
use super::error::APIError;

pub async fn fetch_forge(
    State(state): State<AppState>,
    Path(forge_id): Path<String>,
    claims: Claims,
) -> Result<Json<ForgeBody>, APIError> {
    // Check if user is owner or has access to the forge
    let forge = sqlx::query_as!(
        ForgeBody,
        r#"
        SELECT f.* FROM forges f
        WHERE f.id = $1 AND (
            f.owner_id = $2 
            OR EXISTS (
                SELECT 1 FROM forge_access fa 
                WHERE fa.forge_id = f.id AND fa.user_id = $2
            )
        )
        "#,
        forge_id,
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
    // Get all forges where user is owner or has access
    let forge = sqlx::query_as!(
        ForgeBody,
        r#"
        SELECT f.* FROM forges f
        WHERE f.owner_id = $1 
        OR EXISTS (
            SELECT 1 FROM forge_access fa 
            WHERE fa.forge_id = f.id AND fa.user_id = $1
        )
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
    let forge_id = nanoid::nanoid!();
    
    let _forge = sqlx::query!(
        r#"
        INSERT INTO forges (id, title, description, owner_id)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT DO NOTHING
        "#,
        forge_id,
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
    Path(forge_id): Path<String>
) -> Result<StatusCode, APIError> {
    // Only the owner can delete a forge
    let _forge = sqlx::query!(
        r#"
        DELETE FROM forges WHERE id = $1 AND owner_id = $2
        "#,
        forge_id,
        claims.sub,
    )
    .execute(&state.db)
    .await?;
    
    Ok(StatusCode::NO_CONTENT)
}

pub async fn update_forge(
    State(state): State<AppState>,
    Path(forge_id): Path<String>,
    claims: Claims,
    Json(payload): Json<ForgeUpdate>,
) -> Result<StatusCode, APIError> {
    // Check if user is owner or has admin access to update
    let result = sqlx::query!(
        r#"
        UPDATE forges
        SET
            title = COALESCE($1, title),
            description = COALESCE($2, description)
        WHERE id = $3 AND (
            owner_id = $4
            OR EXISTS (
                SELECT 1 FROM forge_access 
                WHERE forge_id = $3 AND user_id = $4 AND access_role = 'admin'
            )
        )
        "#,
        payload.title,
        payload.description,
        forge_id,
        claims.sub
    )
    .execute(&state.db)
    .await?;
    
    if result.rows_affected() == 0 {
        return Err(APIError::AccessDenied);
    }
    
    Ok(StatusCode::OK)
}

// New function to add a user to a forge
pub async fn add_forge_access(
    State(state): State<AppState>,
    Path(forge_id): Path<String>,
    claims: Claims,
    Json(payload): Json<ForgeAccessCreateBody>,
) -> Result<StatusCode, APIError> {
    // Check if the requester is the owner or has admin access
    let forge = sqlx::query!(
        r#"
        SELECT owner_id FROM forges
        WHERE id = $1 AND (
            owner_id = $2
            OR EXISTS (
                SELECT 1 FROM forge_access 
                WHERE forge_id = $1 AND user_id = $2 AND access_role = 'admin'
            )
        )
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if forge.is_none() {
        return Err(APIError::AccessDenied);
    }
    
    // Add the user to forge_access
    let _access = sqlx::query!(
        r#"
        INSERT INTO forge_access (id, forge_id, user_id, access_role, added_by)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (forge_id, user_id) DO UPDATE
        SET access_role = $4, added_by = $5, updated_at = NOW()
        "#,
        nanoid::nanoid!(),
        forge_id,
        payload.user_id,
        payload.access_role.to_string(),
        claims.sub,
    )
    .execute(&state.db)
    .await?;
    
    Ok(StatusCode::CREATED)
}

// Function to remove a user's access to a forge
pub async fn delete_forge_access(
    State(state): State<AppState>,
    Path((forge_id, user_id)): Path<(String, String)>,
    claims: Claims,
) -> Result<StatusCode, APIError> {
    // Check if the requester is the owner or has admin access
    let forge = sqlx::query!(
        r#"
        SELECT owner_id FROM forges
        WHERE id = $1 AND (
            owner_id = $2
            OR EXISTS (
                SELECT 1 FROM forge_access 
                WHERE forge_id = $1 AND user_id = $2 AND access_role = 'admin'
            )
        )
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if forge.is_none() {
        return Err(APIError::AccessDenied);
    }
    
    // Remove the user's access
    let _result = sqlx::query!(
        r#"
        DELETE FROM forge_access 
        WHERE forge_id = $1 AND user_id = $2
        "#,
        forge_id,
        user_id
    )
    .execute(&state.db)
    .await?;
    
    Ok(StatusCode::NO_CONTENT)
}

// Function to list all users with access to a forge
pub async fn list_forge_access(
    State(state): State<AppState>,
    Path(forge_id): Path<String>,
    claims: Claims,
) -> Result<Json<Vec<ForgeAccessBody>>, APIError> {
    // Check if the requester has access to the forge
    let forge = sqlx::query!(
        r#"
        SELECT 1 AS exists FROM forges
        WHERE id = $1 AND (
            owner_id = $2
            OR EXISTS (
                SELECT 1 FROM forge_access 
                WHERE forge_id = $1 AND user_id = $2
            )
        )
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if forge.is_none() {
        return Err(APIError::AccessDenied);
    }
    
    // Get all users with access to the forge
    let access_list = sqlx::query_as!(
        ForgeAccessBody,
        r#"
        SELECT fa.*, u.name as user_name, u.email as user_email
        FROM forge_access fa
        JOIN users u ON fa.user_id = u.id
        WHERE fa.forge_id = $1
        "#,
        forge_id
    )
    .fetch_all(&state.db)
    .await?;
    
    Ok(Json(access_list))
}

