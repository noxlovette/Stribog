
use serde::{Deserialize, Serialize};
use time::format_description::well_known::Rfc3339;
use time::OffsetDateTime;

#[serde_with::serde_as]
#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct KeyBody {
    pub id: String,
    pub title: String,
    pub forge_id: String,
    pub is_active: bool,
    #[serde_as(as = "Rfc3339")]
    pub created_at: OffsetDateTime,
    #[serde_as(as = "Option<Rfc3339>")]
    pub last_used_at: Option<OffsetDateTime>,
}


#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct KeyCreateBody {
    pub title: String
}

#[derive(Deserialize, Debug)]
#[serde(rename_all = "camelCase")]
pub struct KeyUpdate {
    pub id: Option<String>,
    pub title: Option<String>,
    pub is_active: Option<bool>,
}
