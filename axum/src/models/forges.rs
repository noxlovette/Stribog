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

// Define access roles as an enum
#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "lowercase")]
pub enum ForgeAccessRole {
    Viewer,
    Editor,
    Admin,
}

impl ToString for ForgeAccessRole {
    fn to_string(&self) -> String {
        match self {
            ForgeAccessRole::Viewer => "viewer".to_string(),
            ForgeAccessRole::Editor => "editor".to_string(),
            ForgeAccessRole::Admin => "admin".to_string(),
        }
    }
}

// Model for creating forge access
#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct ForgeAccessCreateBody {
    pub user_id: String,
    pub access_role: ForgeAccessRole,
}

// Model for returning forge access details
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct ForgeAccessBody {
    pub id: String,
    pub forge_id: String,
    pub user_id: String,
    pub user_name: String,
    pub user_email: String,
    pub access_role: String,
    pub added_by: String,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,
}