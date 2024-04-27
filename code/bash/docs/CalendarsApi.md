# CalendarsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**calendarCalendarsCreateCalendar**](CalendarsApi.md#calendarCalendarsCreateCalendar) | **POST** /calendars | 
[**calendarCalendarsDeleteCalendar**](CalendarsApi.md#calendarCalendarsDeleteCalendar) | **DELETE** /calendars/{uuid} | 
[**calendarCalendarsGetCalendar**](CalendarsApi.md#calendarCalendarsGetCalendar) | **GET** /calendars/{uuid} | 
[**calendarCalendarsListCalendars**](CalendarsApi.md#calendarCalendarsListCalendars) | **GET** /calendars | 
[**calendarCalendarsUpdateCalendar**](CalendarsApi.md#calendarCalendarsUpdateCalendar) | **PATCH** /calendars/{uuid} | 



## calendarCalendarsCreateCalendar



synchronize all events of a new calendar

### Example

```bash
 calendarCalendarsCreateCalendar
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCalendarParams** | [**CreateCalendarParams**](CreateCalendarParams.md) |  |

### Return type

(empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not Applicable

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendarCalendarsDeleteCalendar



### Example

```bash
 calendarCalendarsDeleteCalendar uuid=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | [default to null]

### Return type

(empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: Not Applicable

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendarCalendarsGetCalendar



### Example

```bash
 calendarCalendarsGetCalendar uuid=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | [default to null]

### Return type

(empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: Not Applicable

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendarCalendarsListCalendars



### Example

```bash
 calendarCalendarsListCalendars
```

### Parameters

This endpoint does not need any parameter.

### Return type

(empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: Not Applicable

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## calendarCalendarsUpdateCalendar



### Example

```bash
 calendarCalendarsUpdateCalendar uuid=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | [default to null]

### Return type

(empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: Not Applicable

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

