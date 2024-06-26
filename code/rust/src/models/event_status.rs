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

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum EventStatus {
    #[serde(rename = "PrepareRecording")]
    PrepareRecording,
    #[serde(rename = "Recording")]
    Recording,
    #[serde(rename = "Error")]
    Error,
    #[serde(rename = "None")]
    None,

}

impl ToString for EventStatus {
    fn to_string(&self) -> String {
        match self {
            Self::PrepareRecording => String::from("PrepareRecording"),
            Self::Recording => String::from("Recording"),
            Self::Error => String::from("Error"),
            Self::None => String::from("None"),
        }
    }
}

impl Default for EventStatus {
    fn default() -> EventStatus {
        Self::PrepareRecording
    }
}

