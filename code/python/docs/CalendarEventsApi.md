# openapi_client.CalendarEventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendar_calendar_events_get_event**](CalendarEventsApi.md#calendar_calendar_events_get_event) | **GET** /calendar_events/{uuid} | 
[**calendar_calendar_events_list_events**](CalendarEventsApi.md#calendar_calendar_events_list_events) | **GET** /calendar_events | 
[**calendar_calendar_events_schedule_record_event**](CalendarEventsApi.md#calendar_calendar_events_schedule_record_event) | **POST** /calendar_events/{uuid}/bot | 
[**calendar_calendar_events_unschedule_record_event**](CalendarEventsApi.md#calendar_calendar_events_unschedule_record_event) | **DELETE** /calendar_events/{uuid}/bot | 


# **calendar_calendar_events_get_event**
> Event calendar_calendar_events_get_event(uuid)



get event with uuid

### Example


```python
import openapi_client
from openapi_client.models.event import Event
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CalendarEventsApi(api_client)
    uuid = 'uuid_example' # str | 

    try:
        api_response = api_instance.calendar_calendar_events_get_event(uuid)
        print("The response of CalendarEventsApi->calendar_calendar_events_get_event:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CalendarEventsApi->calendar_calendar_events_get_event: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **str**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **calendar_calendar_events_list_events**
> List[Event] calendar_calendar_events_list_events(calendar_id, offset, limit, cursor=cursor)



list all events of a calendar

### Example


```python
import openapi_client
from openapi_client.models.event import Event
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CalendarEventsApi(api_client)
    calendar_id = 56 # int | 
    offset = 56 # int | 
    limit = 56 # int | 
    cursor = 56 # int |  (optional)

    try:
        api_response = api_instance.calendar_calendar_events_list_events(calendar_id, offset, limit, cursor=cursor)
        print("The response of CalendarEventsApi->calendar_calendar_events_list_events:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CalendarEventsApi->calendar_calendar_events_list_events: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calendar_id** | **int**|  | 
 **offset** | **int**|  | 
 **limit** | **int**|  | 
 **cursor** | **int**|  | [optional] 

### Return type

[**List[Event]**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **calendar_calendar_events_schedule_record_event**
> Event calendar_calendar_events_schedule_record_event(uuid)



### Example


```python
import openapi_client
from openapi_client.models.event import Event
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CalendarEventsApi(api_client)
    uuid = 'uuid_example' # str | 

    try:
        api_response = api_instance.calendar_calendar_events_schedule_record_event(uuid)
        print("The response of CalendarEventsApi->calendar_calendar_events_schedule_record_event:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CalendarEventsApi->calendar_calendar_events_schedule_record_event: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **str**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **calendar_calendar_events_unschedule_record_event**
> Event calendar_calendar_events_unschedule_record_event(uuid)



### Example


```python
import openapi_client
from openapi_client.models.event import Event
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CalendarEventsApi(api_client)
    uuid = 'uuid_example' # str | 

    try:
        api_response = api_instance.calendar_calendar_events_unschedule_record_event(uuid)
        print("The response of CalendarEventsApi->calendar_calendar_events_unschedule_record_event:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CalendarEventsApi->calendar_calendar_events_unschedule_record_event: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **str**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

