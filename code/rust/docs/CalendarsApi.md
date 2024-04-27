# \CalendarsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendar_calendars_create_calendar**](CalendarsApi.md#calendar_calendars_create_calendar) | **POST** /calendars | 
[**calendar_calendars_delete_calendar**](CalendarsApi.md#calendar_calendars_delete_calendar) | **DELETE** /calendars/{uuid} | 
[**calendar_calendars_get_calendar**](CalendarsApi.md#calendar_calendars_get_calendar) | **GET** /calendars/{uuid} | 
[**calendar_calendars_list_calendars**](CalendarsApi.md#calendar_calendars_list_calendars) | **GET** /calendars | 
[**calendar_calendars_update_calendar**](CalendarsApi.md#calendar_calendars_update_calendar) | **PATCH** /calendars/{uuid} | 



## calendar_calendars_create_calendar

> calendar_calendars_create_calendar(create_calendar_params)


synchronize all events of a new calendar

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_calendar_params** | [**CreateCalendarParams**](CreateCalendarParams.md) |  | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendar_calendars_delete_calendar

> calendar_calendars_delete_calendar(uuid)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**uuid** | **String** |  | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendar_calendars_get_calendar

> calendar_calendars_get_calendar(uuid)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**uuid** | **String** |  | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendar_calendars_list_calendars

> calendar_calendars_list_calendars()


### Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendar_calendars_update_calendar

> calendar_calendars_update_calendar(uuid)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**uuid** | **String** |  | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

