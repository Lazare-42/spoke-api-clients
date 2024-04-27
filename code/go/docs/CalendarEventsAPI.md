# \CalendarEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalendarCalendarEventsGetEvent**](CalendarEventsAPI.md#CalendarCalendarEventsGetEvent) | **Get** /calendar_events/{uuid} | 
[**CalendarCalendarEventsListEvents**](CalendarEventsAPI.md#CalendarCalendarEventsListEvents) | **Get** /calendar_events | 
[**CalendarCalendarEventsScheduleRecordEvent**](CalendarEventsAPI.md#CalendarCalendarEventsScheduleRecordEvent) | **Post** /calendar_events/{uuid}/bot | 
[**CalendarCalendarEventsUnscheduleRecordEvent**](CalendarEventsAPI.md#CalendarCalendarEventsUnscheduleRecordEvent) | **Delete** /calendar_events/{uuid}/bot | 



## CalendarCalendarEventsGetEvent

> Event CalendarCalendarEventsGetEvent(ctx, uuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarEventsAPI.CalendarCalendarEventsGetEvent(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsAPI.CalendarCalendarEventsGetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalendarCalendarEventsGetEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `CalendarEventsAPI.CalendarCalendarEventsGetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalendarCalendarEventsGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalendarCalendarEventsListEvents

> []Event CalendarCalendarEventsListEvents(ctx).CalendarId(calendarId).Offset(offset).Limit(limit).Cursor(cursor).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	calendarId := int32(56) // int32 | 
	offset := int32(56) // int32 | 
	limit := int32(56) // int32 | 
	cursor := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarEventsAPI.CalendarCalendarEventsListEvents(context.Background()).CalendarId(calendarId).Offset(offset).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsAPI.CalendarCalendarEventsListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalendarCalendarEventsListEvents`: []Event
	fmt.Fprintf(os.Stdout, "Response from `CalendarEventsAPI.CalendarCalendarEventsListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalendarCalendarEventsListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calendarId** | **int32** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **cursor** | **int32** |  | 

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalendarCalendarEventsScheduleRecordEvent

> Event CalendarCalendarEventsScheduleRecordEvent(ctx, uuid).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarEventsAPI.CalendarCalendarEventsScheduleRecordEvent(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsAPI.CalendarCalendarEventsScheduleRecordEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalendarCalendarEventsScheduleRecordEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `CalendarEventsAPI.CalendarCalendarEventsScheduleRecordEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalendarCalendarEventsScheduleRecordEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalendarCalendarEventsUnscheduleRecordEvent

> Event CalendarCalendarEventsUnscheduleRecordEvent(ctx, uuid).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "uuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarEventsAPI.CalendarCalendarEventsUnscheduleRecordEvent(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsAPI.CalendarCalendarEventsUnscheduleRecordEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalendarCalendarEventsUnscheduleRecordEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `CalendarEventsAPI.CalendarCalendarEventsUnscheduleRecordEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalendarCalendarEventsUnscheduleRecordEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

