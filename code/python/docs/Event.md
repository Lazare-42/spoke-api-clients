# Event


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | 
**google_id** | **str** |  | 
**name** | **str** |  | 
**account_id** | **int** |  | 
**meeting_url** | **str** |  | 
**start_time** | [**SystemTime**](SystemTime.md) |  | 
**calendar_id** | **int** |  | 
**attendees** | **object** |  | 
**record** | **bool** |  | 
**is_organizer** | **bool** |  | 
**end_time** | [**SystemTime**](SystemTime.md) |  | 
**status** | [**EventStatus**](EventStatus.md) |  | 
**error** | **str** |  | [optional] 
**session_id** | **str** |  | [optional] 
**project_id** | **int** |  | [optional] 
**refuse_recording_by** | **str** |  | [optional] 
**delete_token** | **str** |  | 
**recurring_event_id** | **str** |  | [optional] 
**is_recurring** | **bool** |  | 
**agenda_id** | **int** |  | [optional] 
**uuid** | **str** |  | 

## Example

```python
from openapi_client.models.event import Event

# TODO update the JSON string below
json = "{}"
# create an instance of Event from a JSON string
event_instance = Event.from_json(json)
# print the JSON string representation of the object
print(Event.to_json())

# convert the object into a dict
event_dict = event_instance.to_dict()
# create an instance of Event from a dict
event_from_dict = Event.from_dict(event_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


