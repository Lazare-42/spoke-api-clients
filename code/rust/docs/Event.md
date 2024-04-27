# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **i32** |  | 
**google_id** | **String** |  | 
**name** | **String** |  | 
**account_id** | **i32** |  | 
**meeting_url** | **String** |  | 
**start_time** | [**models::SystemTime**](SystemTime.md) |  | 
**calendar_id** | **i32** |  | 
**attendees** | Option<[**serde_json::Value**](.md)> |  | 
**record** | **bool** |  | 
**is_organizer** | **bool** |  | 
**end_time** | [**models::SystemTime**](SystemTime.md) |  | 
**status** | [**models::EventStatus**](EventStatus.md) |  | 
**error** | Option<**String**> |  | [optional]
**session_id** | Option<**String**> |  | [optional]
**project_id** | Option<**i32**> |  | [optional]
**refuse_recording_by** | Option<**String**> |  | [optional]
**delete_token** | **String** |  | 
**recurring_event_id** | Option<**String**> |  | [optional]
**is_recurring** | **bool** |  | 
**agenda_id** | Option<**i32**> |  | [optional]
**uuid** | [**uuid::Uuid**](uuid::Uuid.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


