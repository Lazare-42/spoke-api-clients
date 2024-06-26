/*
 * spoke_handler
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.2.0
 * 
 * Generated by: https://openapi-generator.tech
 */

use crate::models;

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateCalendarParams {
    #[serde(rename = "oauth_client_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub oauth_client_id: Option<Option<String>>,
    #[serde(rename = "oauth_client_secret", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub oauth_client_secret: Option<Option<String>>,
    #[serde(rename = "oauth_refresh_token")]
    pub oauth_refresh_token: String,
    #[serde(rename = "platform")]
    pub platform: models::Provider,
    #[serde(rename = "calendar_ids_to_sync", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub calendar_ids_to_sync: Option<Option<Vec<String>>>,
}

impl CreateCalendarParams {
    pub fn new(oauth_refresh_token: String, platform: models::Provider) -> CreateCalendarParams {
        CreateCalendarParams {
            oauth_client_id: None,
            oauth_client_secret: None,
            oauth_refresh_token,
            platform,
            calendar_ids_to_sync: None,
        }
    }
}

