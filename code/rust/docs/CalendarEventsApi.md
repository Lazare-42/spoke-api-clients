# \CalendarEventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendar_calendar_events_get_event**](CalendarEventsApi.md#calendar_calendar_events_get_event) | **GET** /calendar_events/{uuid} | 
[**calendar_calendar_events_list_events**](CalendarEventsApi.md#calendar_calendar_events_list_events) | **GET** /calendar_events | 
[**calendar_calendar_events_schedule_record_event**](CalendarEventsApi.md#calendar_calendar_events_schedule_record_event) | **POST** /calendar_events/{uuid}/bot | 
[**calendar_calendar_events_unschedule_record_event**](CalendarEventsApi.md#calendar_calendar_events_unschedule_record_event) | **DELETE** /calendar_events/{uuid}/bot | 



## calendar_calendar_events_get_event

> models::Event calendar_calendar_events_get_event(uuid)


get event with uuid

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**uuid** | **String** |  | [required] |

### Return type

[**models::Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendar_calendar_events_list_events

> Vec<models::Event> calendar_calendar_events_list_events(calendar_id, offset, limit, cursor)


list all events of a calendar

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**calendar_id** | **i32** |  | [required] |
**offset** | **i32** |  | [required] |
**limit** | **i32** |  | [required] |
**cursor** | Option<**i32**> |  |  |

### Return type

[**Vec<models::Event>**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendar_calendar_events_schedule_record_event

> models::Event calendar_calendar_events_schedule_record_event(uuid)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**uuid** | **String** |  | [required] |

### Return type

[**models::Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendar_calendar_events_unschedule_record_event

> models::Event calendar_calendar_events_unschedule_record_event(uuid)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**uuid** | **String** |  | [required] |

### Return type

[**models::Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

