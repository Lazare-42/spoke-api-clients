# openapi_client.CalendarsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendar_calendars_create_calendar**](CalendarsApi.md#calendar_calendars_create_calendar) | **POST** /calendars | 
[**calendar_calendars_delete_calendar**](CalendarsApi.md#calendar_calendars_delete_calendar) | **DELETE** /calendars/{uuid} | 
[**calendar_calendars_get_calendar**](CalendarsApi.md#calendar_calendars_get_calendar) | **GET** /calendars/{uuid} | 
[**calendar_calendars_list_calendars**](CalendarsApi.md#calendar_calendars_list_calendars) | **GET** /calendars | 
[**calendar_calendars_update_calendar**](CalendarsApi.md#calendar_calendars_update_calendar) | **PATCH** /calendars/{uuid} | 


# **calendar_calendars_create_calendar**
> calendar_calendars_create_calendar(create_calendar_params)



synchronize all events of a new calendar

### Example


```python
import openapi_client
from openapi_client.models.create_calendar_params import CreateCalendarParams
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
    api_instance = openapi_client.CalendarsApi(api_client)
    create_calendar_params = openapi_client.CreateCalendarParams() # CreateCalendarParams | 

    try:
        api_instance.calendar_calendars_create_calendar(create_calendar_params)
    except Exception as e:
        print("Exception when calling CalendarsApi->calendar_calendars_create_calendar: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_calendar_params** | [**CreateCalendarParams**](CreateCalendarParams.md)|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **calendar_calendars_delete_calendar**
> calendar_calendars_delete_calendar(uuid)



### Example


```python
import openapi_client
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
    api_instance = openapi_client.CalendarsApi(api_client)
    uuid = 'uuid_example' # str | 

    try:
        api_instance.calendar_calendars_delete_calendar(uuid)
    except Exception as e:
        print("Exception when calling CalendarsApi->calendar_calendars_delete_calendar: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **str**|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **calendar_calendars_get_calendar**
> calendar_calendars_get_calendar(uuid)



### Example


```python
import openapi_client
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
    api_instance = openapi_client.CalendarsApi(api_client)
    uuid = 'uuid_example' # str | 

    try:
        api_instance.calendar_calendars_get_calendar(uuid)
    except Exception as e:
        print("Exception when calling CalendarsApi->calendar_calendars_get_calendar: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **str**|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **calendar_calendars_list_calendars**
> calendar_calendars_list_calendars()



### Example


```python
import openapi_client
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
    api_instance = openapi_client.CalendarsApi(api_client)

    try:
        api_instance.calendar_calendars_list_calendars()
    except Exception as e:
        print("Exception when calling CalendarsApi->calendar_calendars_list_calendars: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **calendar_calendars_update_calendar**
> calendar_calendars_update_calendar(uuid)



### Example


```python
import openapi_client
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
    api_instance = openapi_client.CalendarsApi(api_client)
    uuid = 'uuid_example' # str | 

    try:
        api_instance.calendar_calendars_update_calendar(uuid)
    except Exception as e:
        print("Exception when calling CalendarsApi->calendar_calendars_update_calendar: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **str**|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

