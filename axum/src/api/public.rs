use axum::{extract::{Path, State}, Json};
use crate::{db::init::AppState, models::sparks::SparkBody};
use super::error::APIError;


// Public endpoint to fetch sparks by API key in URL
pub async fn fetch_public_sparks(
    State(state): State<AppState>,
    Path((api_key, forge_id)): Path<(String, String)>,
) -> Result<Json<Vec<SparkBody>>, APIError> {
    // Verify API key is valid and active
    let api_key_record = sqlx::query!(
        r#"
        SELECT id, forge_id 
        FROM api_keys
        WHERE id = $1 AND forge_id = $2 AND is_active = true
        "#,
        api_key,
        forge_id
    )
    .fetch_optional(&state.db)
    .await?;
    
    if api_key_record.is_none() {
        return Err(APIError::AuthenticationFailed);
    }
    
    // Update last_used_at timestamp
    sqlx::query!(
        r#"
        UPDATE api_keys
        SET last_used_at = NOW()
        WHERE id = $1
        "#,
        api_key
    )
    .execute(&state.db)
    .await?;
    
    // Fetch the sparks
    let sparks = sqlx::query_as!(
        SparkBody,
        r#"
        SELECT id, forge_id, title, markdown, owner_id, created_at, updated_at
        FROM sparks
        WHERE forge_id = $1
        ORDER BY updated_at DESC
        "#,
        forge_id
    )
    .fetch_all(&state.db)
    .await?;
    
    Ok(Json(sparks))
}

// Public endpoint to fetch a specific spark by ID using API key from URL
pub async fn fetch_public_spark(
    State(state): State<AppState>,
    Path((api_key, forge_id, spark_id)): Path<(String, String, String)>,
) -> Result<Json<SparkBody>, APIError> {
    // Verify API key is valid and active
    let api_key_record = sqlx::query!(
        r#"
        SELECT id, forge_id 
        FROM api_keys
        WHERE id = $1 AND forge_id = $2 AND is_active = true
        "#,
        api_key,
        forge_id
    )
    .fetch_optional(&state.db)
    .await?;
    
    if api_key_record.is_none() {
        return Err(APIError::AuthenticationFailed);
    }
    
    // Update last_used_at timestamp
    sqlx::query!(
        r#"
        UPDATE api_keys
        SET last_used_at = NOW()
        WHERE id = $1
        "#,
        api_key
    )
    .execute(&state.db)
    .await?;
    
    // Fetch the specific spark
    let spark = sqlx::query_as!(
        SparkBody,
        r#"
        SELECT id, forge_id, title, markdown, owner_id, created_at, updated_at
        FROM sparks
        WHERE forge_id = $1 AND id = $2
        "#,
        forge_id,
        spark_id
    )
    .fetch_optional(&state.db)
    .await?;
    
    match spark {
        Some(s) => Ok(Json(s)),
        None => Err(APIError::NotFound("Spark not found".into())),
    }
}