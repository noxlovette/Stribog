use axum::{extract::{Path, State}, Json};
use hyper::StatusCode;
use crate::{auth::jwt::Claims, db::init::AppState, models::sparks::{SparkBody, SparkCreateBody, SparkUpdate}};
use super::error::APIError;


pub async fn fetch_spark(
    State(state): State<AppState>,
    Path(id): Path<String>,
    claims: Claims,
) -> Result<Json<SparkBody>, APIError> {
    let spark = sqlx::query_as!(
        SparkBody,
        r#"
        SELECT * FROM sparks
        WHERE id = $1 AND owner_id = $2
        "#,
        id,
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
    let spark = sqlx::query_as!(
        SparkBody,
        r#"
        SELECT * FROM sparks
        WHERE owner_id = $1 AND forge_id = $2
        "#,
        claims.sub,
        forge_id
    )
    .fetch_all(&state.db)
    .await?;

    Ok(Json(spark))
}

pub async fn create_spark(
    State(state): State<AppState>,
    claims: Claims,
    Path(forge_id): Path<String>,
    Json(payload): Json<SparkCreateBody>,
) -> Result<StatusCode, APIError> {

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

pub async fn update_spark(
    State(state): State<AppState>,
    Path(spark_id): Path<String>,
    claims: Claims,
    Json(payload): Json<SparkUpdate>,
) -> Result<StatusCode, APIError> {
    let _forge = sqlx::query!(
        r#"
        UPDATE sparks 
         SET 
            title = COALESCE($1, title),
            markdown = COALESCE($2, markdown)
         WHERE id = $3 AND owner_id = $4
         "#,
        payload.title,
        payload.markdown,
        spark_id,
        claims.sub
    )
    .execute(&state.db)
    .await?;

    Ok(StatusCode::OK)
}