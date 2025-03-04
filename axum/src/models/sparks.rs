use serde::{Deserialize, Serialize};
use time::format_description::well_known::Rfc3339;
use time::OffsetDateTime;


#[serde_with::serde_as]
#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct SparkBody {
    pub id: String,
    pub title: String,
    pub markdown: String,
    pub forge_id: String,
    pub owner_id: String,
    #[serde_as(as = "Rfc3339")]
    pub created_at: OffsetDateTime,
    #[serde_as(as = "Rfc3339")]
    pub updated_at: OffsetDateTime,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct SparkCreateResponse {
    pub id: String,
}


#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct SparkCreateBody {
    pub title: String,
    pub markdown: String,
}

#[derive(Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct SparkUpdate {
    pub title: Option<String>,
    pub markdown: Option<String>,
}
