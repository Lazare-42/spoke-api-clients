# SystemTime


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**secs_since_epoch** | **int** |  | 
**nanos_since_epoch** | **int** |  | 

## Example

```python
from openapi_client.models.system_time import SystemTime

# TODO update the JSON string below
json = "{}"
# create an instance of SystemTime from a JSON string
system_time_instance = SystemTime.from_json(json)
# print the JSON string representation of the object
print(SystemTime.to_json())

# convert the object into a dict
system_time_dict = system_time_instance.to_dict()
# create an instance of SystemTime from a dict
system_time_from_dict = SystemTime.from_dict(system_time_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


