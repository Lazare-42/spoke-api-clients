/*
spoke_handler

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CalendarEventsAPIService CalendarEventsAPI service
type CalendarEventsAPIService service

type ApiCalendarCalendarEventsGetEventRequest struct {
	ctx context.Context
	ApiService *CalendarEventsAPIService
	uuid string
}

func (r ApiCalendarCalendarEventsGetEventRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.CalendarCalendarEventsGetEventExecute(r)
}

/*
CalendarCalendarEventsGetEvent Method for CalendarCalendarEventsGetEvent

get event with uuid

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiCalendarCalendarEventsGetEventRequest
*/
func (a *CalendarEventsAPIService) CalendarCalendarEventsGetEvent(ctx context.Context, uuid string) ApiCalendarCalendarEventsGetEventRequest {
	return ApiCalendarCalendarEventsGetEventRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return Event
func (a *CalendarEventsAPIService) CalendarCalendarEventsGetEventExecute(r ApiCalendarCalendarEventsGetEventRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsAPIService.CalendarCalendarEventsGetEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar_events/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCalendarCalendarEventsListEventsRequest struct {
	ctx context.Context
	ApiService *CalendarEventsAPIService
	calendarId *int32
	offset *int32
	limit *int32
	cursor *int32
}

func (r ApiCalendarCalendarEventsListEventsRequest) CalendarId(calendarId int32) ApiCalendarCalendarEventsListEventsRequest {
	r.calendarId = &calendarId
	return r
}

func (r ApiCalendarCalendarEventsListEventsRequest) Offset(offset int32) ApiCalendarCalendarEventsListEventsRequest {
	r.offset = &offset
	return r
}

func (r ApiCalendarCalendarEventsListEventsRequest) Limit(limit int32) ApiCalendarCalendarEventsListEventsRequest {
	r.limit = &limit
	return r
}

func (r ApiCalendarCalendarEventsListEventsRequest) Cursor(cursor int32) ApiCalendarCalendarEventsListEventsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiCalendarCalendarEventsListEventsRequest) Execute() ([]Event, *http.Response, error) {
	return r.ApiService.CalendarCalendarEventsListEventsExecute(r)
}

/*
CalendarCalendarEventsListEvents Method for CalendarCalendarEventsListEvents

list all events of a calendar

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCalendarCalendarEventsListEventsRequest
*/
func (a *CalendarEventsAPIService) CalendarCalendarEventsListEvents(ctx context.Context) ApiCalendarCalendarEventsListEventsRequest {
	return ApiCalendarCalendarEventsListEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Event
func (a *CalendarEventsAPIService) CalendarCalendarEventsListEventsExecute(r ApiCalendarCalendarEventsListEventsRequest) ([]Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsAPIService.CalendarCalendarEventsListEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar_events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.calendarId == nil {
		return localVarReturnValue, nil, reportError("calendarId is required and must be specified")
	}
	if r.offset == nil {
		return localVarReturnValue, nil, reportError("offset is required and must be specified")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "calendar_id", r.calendarId, "")
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCalendarCalendarEventsScheduleRecordEventRequest struct {
	ctx context.Context
	ApiService *CalendarEventsAPIService
	uuid string
}

func (r ApiCalendarCalendarEventsScheduleRecordEventRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.CalendarCalendarEventsScheduleRecordEventExecute(r)
}

/*
CalendarCalendarEventsScheduleRecordEvent Method for CalendarCalendarEventsScheduleRecordEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiCalendarCalendarEventsScheduleRecordEventRequest
*/
func (a *CalendarEventsAPIService) CalendarCalendarEventsScheduleRecordEvent(ctx context.Context, uuid string) ApiCalendarCalendarEventsScheduleRecordEventRequest {
	return ApiCalendarCalendarEventsScheduleRecordEventRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return Event
func (a *CalendarEventsAPIService) CalendarCalendarEventsScheduleRecordEventExecute(r ApiCalendarCalendarEventsScheduleRecordEventRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsAPIService.CalendarCalendarEventsScheduleRecordEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar_events/{uuid}/bot"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCalendarCalendarEventsUnscheduleRecordEventRequest struct {
	ctx context.Context
	ApiService *CalendarEventsAPIService
	uuid string
}

func (r ApiCalendarCalendarEventsUnscheduleRecordEventRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.CalendarCalendarEventsUnscheduleRecordEventExecute(r)
}

/*
CalendarCalendarEventsUnscheduleRecordEvent Method for CalendarCalendarEventsUnscheduleRecordEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid
 @return ApiCalendarCalendarEventsUnscheduleRecordEventRequest
*/
func (a *CalendarEventsAPIService) CalendarCalendarEventsUnscheduleRecordEvent(ctx context.Context, uuid string) ApiCalendarCalendarEventsUnscheduleRecordEventRequest {
	return ApiCalendarCalendarEventsUnscheduleRecordEventRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return Event
func (a *CalendarEventsAPIService) CalendarCalendarEventsUnscheduleRecordEventExecute(r ApiCalendarCalendarEventsUnscheduleRecordEventRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsAPIService.CalendarCalendarEventsUnscheduleRecordEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar_events/{uuid}/bot"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
