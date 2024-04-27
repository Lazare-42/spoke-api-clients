# CalendarEventsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendarCalendarEventsGetEvent**](CalendarEventsApi.md#calendarCalendarEventsGetEvent) | **GET** /calendar_events/{uuid} | 
[**calendarCalendarEventsListEvents**](CalendarEventsApi.md#calendarCalendarEventsListEvents) | **GET** /calendar_events | 
[**calendarCalendarEventsScheduleRecordEvent**](CalendarEventsApi.md#calendarCalendarEventsScheduleRecordEvent) | **POST** /calendar_events/{uuid}/bot | 
[**calendarCalendarEventsUnscheduleRecordEvent**](CalendarEventsApi.md#calendarCalendarEventsUnscheduleRecordEvent) | **DELETE** /calendar_events/{uuid}/bot | 



## calendarCalendarEventsGetEvent



get event with uuid

### Example

```bash
 calendarCalendarEventsGetEvent uuid=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | [default to null]

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendarCalendarEventsListEvents



list all events of a calendar

### Example

```bash
 calendarCalendarEventsListEvents  calendar_id=value  offset=value  limit=value  cursor=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calendarId** | **integer** |  | [default to null]
 **offset** | **integer** |  | [default to null]
 **limit** | **integer** |  | [default to null]
 **cursor** | **integer** |  | [optional] [default to null]

### Return type

[**array[Event]**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendarCalendarEventsScheduleRecordEvent



### Example

```bash
 calendarCalendarEventsScheduleRecordEvent uuid=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | [default to null]

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendarCalendarEventsUnscheduleRecordEvent



### Example

```bash
 calendarCalendarEventsUnscheduleRecordEvent uuid=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | [default to null]

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

