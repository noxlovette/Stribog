use axum::{extract::{Path, State}, Json};
use hyper::StatusCode;
use crate::{auth::jwt::Claims, db::init::AppState, models::sparks::{SparkBody, SparkCreateBody, SparkUpdate}};
use super::error::APIError;


// Spark handlers updated to consider forge access

pub async fn fetch_spark(
    State(state): State<AppState>,
    Path(spark_id): Path<String>,
    claims: Claims,
) -> Result<Json<SparkBody>, APIError> {
    let spark = sqlx::query_as!(
        SparkBody,
        r#"
        SELECT s.* FROM sparks s
        JOIN forges f ON s.forge_id = f.id
        WHERE s.id = $1 AND (
            s.owner_id = $2
            OR f.owner_id = $2
            OR EXISTS (
                SELECT 1 FROM forge_access fa 
                WHERE fa.forge_id = s.forge_id AND fa.user_id = $2
            )
        )
        "#,
        spark_id,
        claims.sub
    )
    .fetch_one(&state.db)
    .await?;
    
    Ok(Json(spark))
}

pub async fn list_spark(
    State(state): State<AppState>,
    claims: Claims,
    Path(forge_id): Path<String>,
) -> Result<Json<Vec<SparkBody>>, APIError> {
    // Check if user has access to this forge
    let forge_access = sqlx::query!(
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
    
    if forge_access.is_none() {
        return Err(APIError::AccessDenied);
    }
    
    let sparks = sqlx::query_as!(
        SparkBody,
        r#"
        SELECT * FROM sparks
        WHERE forge_id = $1
        "#,
        forge_id
    )
    .fetch_all(&state.db)
    .await?;
    
    Ok(Json(sparks))
}

pub async fn create_spark(
    State(state): State<AppState>,
    claims: Claims,
    Path(forge_id): Path<String>,
    Json(payload): Json<SparkCreateBody>,
) -> Result<StatusCode, APIError> {
    // Check if user can create in this forge (owner, admin, or editor)
    let can_create = sqlx::query!(
        r#"
        SELECT 1 as exists FROM forges
        WHERE id = $1 AND (
            owner_id = $2
            OR EXISTS (
                SELECT 1 FROM forge_access 
                WHERE forge_id = $1 AND user_id = $2 
                AND access_role IN ('admin', 'editor')
            )
        )
        "#,
        forge_id,
        claims.sub
    )
    .fetch_optional(&state.db)
    .await?;
    
    if can_create.is_none() {
        return Err(APIError::AccessDenied);
    }
    
    let _sparks = sqlx::query!(
        r#"
        INSERT INTO sparks (id, title, markdown, forge_id, owner_id)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT DO NOTHING
        "#,
        nanoid::nanoid!(),
        payload.title,
        payload.markdown,
        forge_id,
        claims.sub,
    )
    .execute(&state.db)
    .await?;
    
    Ok(StatusCode::CREATED)
}

pub async fn delete_spark(
    State(state): State<AppState>,
    claims: Claims,
    Path(spark_id): Path<String>
) -> Result<StatusCode, APIError> {
    // Allow deletion if: spark owner, forge owner, or admin
    let result = sqlx::query!(
        r#"
        DELETE FROM sparks 
        WHERE id = $1 AND (
            owner_id = $2
            OR EXISTS (
                SELECT 1 FROM forges f
                WHERE f.id = sparks.forge_id AND f.owner_id = $2
            )
            OR EXISTS (
                SELECT 1 FROM forge_access fa
                WHERE fa.forge_id = sparks.forge_id 
                AND fa.user_id = $2 
                AND fa.access_role = 'admin'
            )
        )
        "#,
        spark_id,
        claims.sub,
    )
    .execute(&state.db)
    .await?;
    
    if result.rows_affected() == 0 {
        return Err(APIError::AccessDenied);
    }
    
    Ok(StatusCode::NO_CONTENT)
}

pub async fn update_spark(
    State(state): State<AppState>,
    Path(spark_id): Path<String>,
    claims: Claims,
    Json(payload): Json<SparkUpdate>,
) -> Result<StatusCode, APIError> {
    // Allow update if: spark owner, forge owner, admin, or editor
    let result = sqlx::query!(
        r#"
        UPDATE sparks
        SET
            title = COALESCE($1, title),
            markdown = COALESCE($2, markdown)
        WHERE id = $3 AND (
            owner_id = $4
            OR EXISTS (
                SELECT 1 FROM forges f
                WHERE f.id = sparks.forge_id AND f.owner_id = $4
            )
            OR EXISTS (
                SELECT 1 FROM forge_access fa
                WHERE fa.forge_id = sparks.forge_id 
                AND fa.user_id = $4 
                AND fa.access_role IN ('admin', 'editor')
            )
        )
        "#,
        payload.title,
        payload.markdown,
        spark_id,
        claims.sub
    )
    .execute(&state.db)
    .await?;
    
    if result.rows_affected() == 0 {
        return Err(APIError::AccessDenied);
    }
    
    Ok(StatusCode::OK)
}