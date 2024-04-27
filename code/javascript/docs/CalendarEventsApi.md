# SpokeHandler.CalendarEventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendarCalendarEventsGetEvent**](CalendarEventsApi.md#calendarCalendarEventsGetEvent) | **GET** /calendar_events/{uuid} | 
[**calendarCalendarEventsListEvents**](CalendarEventsApi.md#calendarCalendarEventsListEvents) | **GET** /calendar_events | 
[**calendarCalendarEventsScheduleRecordEvent**](CalendarEventsApi.md#calendarCalendarEventsScheduleRecordEvent) | **POST** /calendar_events/{uuid}/bot | 
[**calendarCalendarEventsUnscheduleRecordEvent**](CalendarEventsApi.md#calendarCalendarEventsUnscheduleRecordEvent) | **DELETE** /calendar_events/{uuid}/bot | 



## calendarCalendarEventsGetEvent

> Event calendarCalendarEventsGetEvent(uuid)



get event with uuid

### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarEventsApi();
let uuid = "uuid_example"; // String | 
apiInstance.calendarCalendarEventsGetEvent(uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **String**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## calendarCalendarEventsListEvents

> [Event] calendarCalendarEventsListEvents(calendarId, offset, limit, opts)



list all events of a calendar

### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarEventsApi();
let calendarId = 56; // Number | 
let offset = 56; // Number | 
let limit = 56; // Number | 
let opts = {
  'cursor': 56 // Number | 
};
apiInstance.calendarCalendarEventsListEvents(calendarId, offset, limit, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calendarId** | **Number**|  | 
 **offset** | **Number**|  | 
 **limit** | **Number**|  | 
 **cursor** | **Number**|  | [optional] 

### Return type

[**[Event]**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## calendarCalendarEventsScheduleRecordEvent

> Event calendarCalendarEventsScheduleRecordEvent(uuid)



### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarEventsApi();
let uuid = "uuid_example"; // String | 
apiInstance.calendarCalendarEventsScheduleRecordEvent(uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **String**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## calendarCalendarEventsUnscheduleRecordEvent

> Event calendarCalendarEventsUnscheduleRecordEvent(uuid)



### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarEventsApi();
let uuid = "uuid_example"; // String | 
apiInstance.calendarCalendarEventsUnscheduleRecordEvent(uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **String**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

