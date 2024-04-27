# CalendarsApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**calendarCalendarsCreateCalendar**](CalendarsApi.md#calendarCalendarsCreateCalendar) | **POST** /calendars |  |
| [**calendarCalendarsDeleteCalendar**](CalendarsApi.md#calendarCalendarsDeleteCalendar) | **DELETE** /calendars/{uuid} |  |
| [**calendarCalendarsGetCalendar**](CalendarsApi.md#calendarCalendarsGetCalendar) | **GET** /calendars/{uuid} |  |
| [**calendarCalendarsListCalendars**](CalendarsApi.md#calendarCalendarsListCalendars) | **GET** /calendars |  |
| [**calendarCalendarsUpdateCalendar**](CalendarsApi.md#calendarCalendarsUpdateCalendar) | **PATCH** /calendars/{uuid} |  |


<a name="calendarCalendarsCreateCalendar"></a>
# **calendarCalendarsCreateCalendar**
> calendarCalendarsCreateCalendar(CreateCalendarParams)



    synchronize all events of a new calendar

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateCalendarParams** | [**CreateCalendarParams**](../Models/CreateCalendarParams.md)|  | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="calendarCalendarsDeleteCalendar"></a>
# **calendarCalendarsDeleteCalendar**
> calendarCalendarsDeleteCalendar(uuid)



### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **uuid** | **String**|  | [default to null] |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="calendarCalendarsGetCalendar"></a>
# **calendarCalendarsGetCalendar**
> calendarCalendarsGetCalendar(uuid)



### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **uuid** | **String**|  | [default to null] |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="calendarCalendarsListCalendars"></a>
# **calendarCalendarsListCalendars**
> calendarCalendarsListCalendars()



### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="calendarCalendarsUpdateCalendar"></a>
# **calendarCalendarsUpdateCalendar**
> calendarCalendarsUpdateCalendar(uuid)



### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **uuid** | **String**|  | [default to null] |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

