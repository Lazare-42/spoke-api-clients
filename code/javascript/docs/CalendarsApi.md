# SpokeHandler.CalendarsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendarCalendarsCreateCalendar**](CalendarsApi.md#calendarCalendarsCreateCalendar) | **POST** /calendars | 
[**calendarCalendarsDeleteCalendar**](CalendarsApi.md#calendarCalendarsDeleteCalendar) | **DELETE** /calendars/{uuid} | 
[**calendarCalendarsGetCalendar**](CalendarsApi.md#calendarCalendarsGetCalendar) | **GET** /calendars/{uuid} | 
[**calendarCalendarsListCalendars**](CalendarsApi.md#calendarCalendarsListCalendars) | **GET** /calendars | 
[**calendarCalendarsUpdateCalendar**](CalendarsApi.md#calendarCalendarsUpdateCalendar) | **PATCH** /calendars/{uuid} | 



## calendarCalendarsCreateCalendar

> calendarCalendarsCreateCalendar(createCalendarParams)



synchronize all events of a new calendar

### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarsApi();
let createCalendarParams = new SpokeHandler.CreateCalendarParams(); // CreateCalendarParams | 
apiInstance.calendarCalendarsCreateCalendar(createCalendarParams, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCalendarParams** | [**CreateCalendarParams**](CreateCalendarParams.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## calendarCalendarsDeleteCalendar

> calendarCalendarsDeleteCalendar(uuid)



### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarsApi();
let uuid = "uuid_example"; // String | 
apiInstance.calendarCalendarsDeleteCalendar(uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **String**|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## calendarCalendarsGetCalendar

> calendarCalendarsGetCalendar(uuid)



### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarsApi();
let uuid = "uuid_example"; // String | 
apiInstance.calendarCalendarsGetCalendar(uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **String**|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## calendarCalendarsListCalendars

> calendarCalendarsListCalendars()



### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarsApi();
apiInstance.calendarCalendarsListCalendars((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## calendarCalendarsUpdateCalendar

> calendarCalendarsUpdateCalendar(uuid)



### Example

```javascript
import SpokeHandler from 'spoke_handler';

let apiInstance = new SpokeHandler.CalendarsApi();
let uuid = "uuid_example"; // String | 
apiInstance.calendarCalendarsUpdateCalendar(uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **String**|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

