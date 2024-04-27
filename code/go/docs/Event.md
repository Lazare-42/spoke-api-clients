# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**GoogleId** | **string** |  | 
**Name** | **string** |  | 
**AccountId** | **int32** |  | 
**MeetingUrl** | **string** |  | 
**StartTime** | [**SystemTime**](SystemTime.md) |  | 
**CalendarId** | **int32** |  | 
**Attendees** | **interface{}** |  | 
**Record** | **bool** |  | 
**IsOrganizer** | **bool** |  | 
**EndTime** | [**SystemTime**](SystemTime.md) |  | 
**Status** | [**EventStatus**](EventStatus.md) |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**SessionId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**RefuseRecordingBy** | Pointer to **NullableString** |  | [optional] 
**DeleteToken** | **string** |  | 
**RecurringEventId** | Pointer to **NullableString** |  | [optional] 
**IsRecurring** | **bool** |  | 
**AgendaId** | Pointer to **NullableInt32** |  | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewEvent

`func NewEvent(id int32, googleId string, name string, accountId int32, meetingUrl string, startTime SystemTime, calendarId int32, attendees interface{}, record bool, isOrganizer bool, endTime SystemTime, status EventStatus, deleteToken string, isRecurring bool, uuid string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v int32)`

SetId sets Id field to given value.


### GetGoogleId

`func (o *Event) GetGoogleId() string`

GetGoogleId returns the GoogleId field if non-nil, zero value otherwise.

### GetGoogleIdOk

`func (o *Event) GetGoogleIdOk() (*string, bool)`

GetGoogleIdOk returns a tuple with the GoogleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleId

`func (o *Event) SetGoogleId(v string)`

SetGoogleId sets GoogleId field to given value.


### GetName

`func (o *Event) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Event) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Event) SetName(v string)`

SetName sets Name field to given value.


### GetAccountId

`func (o *Event) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Event) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Event) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetMeetingUrl

`func (o *Event) GetMeetingUrl() string`

GetMeetingUrl returns the MeetingUrl field if non-nil, zero value otherwise.

### GetMeetingUrlOk

`func (o *Event) GetMeetingUrlOk() (*string, bool)`

GetMeetingUrlOk returns a tuple with the MeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingUrl

`func (o *Event) SetMeetingUrl(v string)`

SetMeetingUrl sets MeetingUrl field to given value.


### GetStartTime

`func (o *Event) GetStartTime() SystemTime`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Event) GetStartTimeOk() (*SystemTime, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Event) SetStartTime(v SystemTime)`

SetStartTime sets StartTime field to given value.


### GetCalendarId

`func (o *Event) GetCalendarId() int32`

GetCalendarId returns the CalendarId field if non-nil, zero value otherwise.

### GetCalendarIdOk

`func (o *Event) GetCalendarIdOk() (*int32, bool)`

GetCalendarIdOk returns a tuple with the CalendarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarId

`func (o *Event) SetCalendarId(v int32)`

SetCalendarId sets CalendarId field to given value.


### GetAttendees

`func (o *Event) GetAttendees() interface{}`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *Event) GetAttendeesOk() (*interface{}, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *Event) SetAttendees(v interface{})`

SetAttendees sets Attendees field to given value.


### SetAttendeesNil

`func (o *Event) SetAttendeesNil(b bool)`

 SetAttendeesNil sets the value for Attendees to be an explicit nil

### UnsetAttendees
`func (o *Event) UnsetAttendees()`

UnsetAttendees ensures that no value is present for Attendees, not even an explicit nil
### GetRecord

`func (o *Event) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *Event) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *Event) SetRecord(v bool)`

SetRecord sets Record field to given value.


### GetIsOrganizer

`func (o *Event) GetIsOrganizer() bool`

GetIsOrganizer returns the IsOrganizer field if non-nil, zero value otherwise.

### GetIsOrganizerOk

`func (o *Event) GetIsOrganizerOk() (*bool, bool)`

GetIsOrganizerOk returns a tuple with the IsOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizer

`func (o *Event) SetIsOrganizer(v bool)`

SetIsOrganizer sets IsOrganizer field to given value.


### GetEndTime

`func (o *Event) GetEndTime() SystemTime`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Event) GetEndTimeOk() (*SystemTime, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Event) SetEndTime(v SystemTime)`

SetEndTime sets EndTime field to given value.


### GetStatus

`func (o *Event) GetStatus() EventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Event) GetStatusOk() (*EventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Event) SetStatus(v EventStatus)`

SetStatus sets Status field to given value.


### GetError

`func (o *Event) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Event) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Event) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Event) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Event) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Event) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetSessionId

`func (o *Event) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Event) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Event) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Event) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *Event) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *Event) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetProjectId

`func (o *Event) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Event) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Event) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Event) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Event) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Event) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetRefuseRecordingBy

`func (o *Event) GetRefuseRecordingBy() string`

GetRefuseRecordingBy returns the RefuseRecordingBy field if non-nil, zero value otherwise.

### GetRefuseRecordingByOk

`func (o *Event) GetRefuseRecordingByOk() (*string, bool)`

GetRefuseRecordingByOk returns a tuple with the RefuseRecordingBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefuseRecordingBy

`func (o *Event) SetRefuseRecordingBy(v string)`

SetRefuseRecordingBy sets RefuseRecordingBy field to given value.

### HasRefuseRecordingBy

`func (o *Event) HasRefuseRecordingBy() bool`

HasRefuseRecordingBy returns a boolean if a field has been set.

### SetRefuseRecordingByNil

`func (o *Event) SetRefuseRecordingByNil(b bool)`

 SetRefuseRecordingByNil sets the value for RefuseRecordingBy to be an explicit nil

### UnsetRefuseRecordingBy
`func (o *Event) UnsetRefuseRecordingBy()`

UnsetRefuseRecordingBy ensures that no value is present for RefuseRecordingBy, not even an explicit nil
### GetDeleteToken

`func (o *Event) GetDeleteToken() string`

GetDeleteToken returns the DeleteToken field if non-nil, zero value otherwise.

### GetDeleteTokenOk

`func (o *Event) GetDeleteTokenOk() (*string, bool)`

GetDeleteTokenOk returns a tuple with the DeleteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteToken

`func (o *Event) SetDeleteToken(v string)`

SetDeleteToken sets DeleteToken field to given value.


### GetRecurringEventId

`func (o *Event) GetRecurringEventId() string`

GetRecurringEventId returns the RecurringEventId field if non-nil, zero value otherwise.

### GetRecurringEventIdOk

`func (o *Event) GetRecurringEventIdOk() (*string, bool)`

GetRecurringEventIdOk returns a tuple with the RecurringEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringEventId

`func (o *Event) SetRecurringEventId(v string)`

SetRecurringEventId sets RecurringEventId field to given value.

### HasRecurringEventId

`func (o *Event) HasRecurringEventId() bool`

HasRecurringEventId returns a boolean if a field has been set.

### SetRecurringEventIdNil

`func (o *Event) SetRecurringEventIdNil(b bool)`

 SetRecurringEventIdNil sets the value for RecurringEventId to be an explicit nil

### UnsetRecurringEventId
`func (o *Event) UnsetRecurringEventId()`

UnsetRecurringEventId ensures that no value is present for RecurringEventId, not even an explicit nil
### GetIsRecurring

`func (o *Event) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *Event) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *Event) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.


### GetAgendaId

`func (o *Event) GetAgendaId() int32`

GetAgendaId returns the AgendaId field if non-nil, zero value otherwise.

### GetAgendaIdOk

`func (o *Event) GetAgendaIdOk() (*int32, bool)`

GetAgendaIdOk returns a tuple with the AgendaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgendaId

`func (o *Event) SetAgendaId(v int32)`

SetAgendaId sets AgendaId field to given value.

### HasAgendaId

`func (o *Event) HasAgendaId() bool`

HasAgendaId returns a boolean if a field has been set.

### SetAgendaIdNil

`func (o *Event) SetAgendaIdNil(b bool)`

 SetAgendaIdNil sets the value for AgendaId to be an explicit nil

### UnsetAgendaId
`func (o *Event) UnsetAgendaId()`

UnsetAgendaId ensures that no value is present for AgendaId, not even an explicit nil
### GetUuid

`func (o *Event) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Event) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Event) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


