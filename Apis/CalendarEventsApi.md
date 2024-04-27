# CalendarEventsApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**calendarCalendarEventsGetEvent**](CalendarEventsApi.md#calendarCalendarEventsGetEvent) | **GET** /calendar_events/{uuid} |  |
| [**calendarCalendarEventsListEvents**](CalendarEventsApi.md#calendarCalendarEventsListEvents) | **GET** /calendar_events |  |
| [**calendarCalendarEventsScheduleRecordEvent**](CalendarEventsApi.md#calendarCalendarEventsScheduleRecordEvent) | **POST** /calendar_events/{uuid}/bot |  |
| [**calendarCalendarEventsUnscheduleRecordEvent**](CalendarEventsApi.md#calendarCalendarEventsUnscheduleRecordEvent) | **DELETE** /calendar_events/{uuid}/bot |  |


<a name="calendarCalendarEventsGetEvent"></a>
# **calendarCalendarEventsGetEvent**
> Event calendarCalendarEventsGetEvent(uuid)



    get event with uuid

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **uuid** | **String**|  | [default to null] |

### Return type

[**Event**](../Models/Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="calendarCalendarEventsListEvents"></a>
# **calendarCalendarEventsListEvents**
> List calendarCalendarEventsListEvents(calendar\_id, offset, limit, cursor)



    list all events of a calendar

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **calendar\_id** | **Integer**|  | [default to null] |
| **offset** | **Integer**|  | [default to null] |
| **limit** | **Integer**|  | [default to null] |
| **cursor** | **Integer**|  | [optional] [default to null] |

### Return type

[**List**](../Models/Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="calendarCalendarEventsScheduleRecordEvent"></a>
# **calendarCalendarEventsScheduleRecordEvent**
> Event calendarCalendarEventsScheduleRecordEvent(uuid)



### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **uuid** | **String**|  | [default to null] |

### Return type

[**Event**](../Models/Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="calendarCalendarEventsUnscheduleRecordEvent"></a>
# **calendarCalendarEventsUnscheduleRecordEvent**
> Event calendarCalendarEventsUnscheduleRecordEvent(uuid)



### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **uuid** | **String**|  | [default to null] |

### Return type

[**Event**](../Models/Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

