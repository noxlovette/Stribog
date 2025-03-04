use crate::api::forge;
use crate::db::init::AppState;
use axum::routing::get;
use axum::Router;

pub fn forge_routes() -> Router<AppState> {
    Router::new()
        .route("/", get(forge::list_forge).post(forge::create_forge))
        .route(
            "/{forge_id}",
            get(forge::fetch_forge)
                .patch(forge::update_forge)
                .delete(forge::delete_forge),
        )
        .route(
            "/{forge_id}/access",
            get(forge::list_forge_access).post(forge::add_forge_access).delete(forge::delete_forge_access)
        )
}
 