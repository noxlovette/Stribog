use crate::api::keys;
use crate::db::init::AppState;
use axum::routing::{get, put};
use axum::Router;

pub fn key_routes() -> Router<AppState> {
    Router::new().route(
        "/{forge_id}",
        get(keys::list_api_keys)
            .post(keys::create_api_key)
    )
    .route("/{forge_id}/{key_id}", put(keys::toggle_api_key).delete(keys::delete_api_key))
}
