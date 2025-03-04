use serde::{Deserialize, Serialize};
use time::format_description::well_known::Rfc3339;
use time::OffsetDateTime;

#[serde_with::serde_as]
#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct ForgeBody {
    pub id: String,
    pub title: String,
    pub description: Option<String>,
    pub owner_id: String,
    #[serde_as(as = "Rfc3339")]
    pub created_at: OffsetDateTime,
    #[serde_as(as = "Rfc3339")]
    pub updated_at: OffsetDateTime,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct ForgeCreateResponse {
    pub id: String,
}


#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct ForgeCreateBody {
    pub title: String,
    pub description: Option<String>
}

#[derive(Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct ForgeUpdate {
    pub id: Option<String>,
    pub title: Option<String>,
    pub description: Option<String>,
}
