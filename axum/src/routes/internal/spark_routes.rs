use crate::api::spark;
use crate::db::init::AppState;
use axum::routing::get;
use axum::Router;

pub fn spark_routes() -> Router<AppState> {
    Router::new()
        .route("/{forge_id}", get(spark::list_spark).post(spark::create_spark))
        .route(
            "/s/{spark_id}",
            get(spark::fetch_spark)
                .patch(spark::update_spark)
                .delete(spark::delete_spark),
        )
}
 