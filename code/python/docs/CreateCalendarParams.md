# CreateCalendarParams


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**oauth_client_id** | **str** |  | [optional] 
**oauth_client_secret** | **str** |  | [optional] 
**oauth_refresh_token** | **str** |  | 
**platform** | [**Provider**](Provider.md) |  | 
**calendar_ids_to_sync** | **List[str]** |  | [optional] 

## Example

```python
from openapi_client.models.create_calendar_params import CreateCalendarParams

# TODO update the JSON string below
json = "{}"
# create an instance of CreateCalendarParams from a JSON string
create_calendar_params_instance = CreateCalendarParams.from_json(json)
# print the JSON string representation of the object
print(CreateCalendarParams.to_json())

# convert the object into a dict
create_calendar_params_dict = create_calendar_params_instance.to_dict()
# create an instance of CreateCalendarParams from a dict
create_calendar_params_from_dict = CreateCalendarParams.from_dict(create_calendar_params_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


