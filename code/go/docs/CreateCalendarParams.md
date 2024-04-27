# CreateCalendarParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OauthClientId** | Pointer to **NullableString** |  | [optional] 
**OauthClientSecret** | Pointer to **NullableString** |  | [optional] 
**OauthRefreshToken** | **string** |  | 
**Platform** | [**Provider**](Provider.md) |  | 
**CalendarIdsToSync** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateCalendarParams

`func NewCreateCalendarParams(oauthRefreshToken string, platform Provider, ) *CreateCalendarParams`

NewCreateCalendarParams instantiates a new CreateCalendarParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCalendarParamsWithDefaults

`func NewCreateCalendarParamsWithDefaults() *CreateCalendarParams`

NewCreateCalendarParamsWithDefaults instantiates a new CreateCalendarParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauthClientId

`func (o *CreateCalendarParams) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *CreateCalendarParams) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *CreateCalendarParams) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *CreateCalendarParams) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### SetOauthClientIdNil

`func (o *CreateCalendarParams) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *CreateCalendarParams) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetOauthClientSecret

`func (o *CreateCalendarParams) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *CreateCalendarParams) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *CreateCalendarParams) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *CreateCalendarParams) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### SetOauthClientSecretNil

`func (o *CreateCalendarParams) SetOauthClientSecretNil(b bool)`

 SetOauthClientSecretNil sets the value for OauthClientSecret to be an explicit nil

### UnsetOauthClientSecret
`func (o *CreateCalendarParams) UnsetOauthClientSecret()`

UnsetOauthClientSecret ensures that no value is present for OauthClientSecret, not even an explicit nil
### GetOauthRefreshToken

`func (o *CreateCalendarParams) GetOauthRefreshToken() string`

GetOauthRefreshToken returns the OauthRefreshToken field if non-nil, zero value otherwise.

### GetOauthRefreshTokenOk

`func (o *CreateCalendarParams) GetOauthRefreshTokenOk() (*string, bool)`

GetOauthRefreshTokenOk returns a tuple with the OauthRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthRefreshToken

`func (o *CreateCalendarParams) SetOauthRefreshToken(v string)`

SetOauthRefreshToken sets OauthRefreshToken field to given value.


### GetPlatform

`func (o *CreateCalendarParams) GetPlatform() Provider`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateCalendarParams) GetPlatformOk() (*Provider, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateCalendarParams) SetPlatform(v Provider)`

SetPlatform sets Platform field to given value.


### GetCalendarIdsToSync

`func (o *CreateCalendarParams) GetCalendarIdsToSync() []string`

GetCalendarIdsToSync returns the CalendarIdsToSync field if non-nil, zero value otherwise.

### GetCalendarIdsToSyncOk

`func (o *CreateCalendarParams) GetCalendarIdsToSyncOk() (*[]string, bool)`

GetCalendarIdsToSyncOk returns a tuple with the CalendarIdsToSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarIdsToSync

`func (o *CreateCalendarParams) SetCalendarIdsToSync(v []string)`

SetCalendarIdsToSync sets CalendarIdsToSync field to given value.

### HasCalendarIdsToSync

`func (o *CreateCalendarParams) HasCalendarIdsToSync() bool`

HasCalendarIdsToSync returns a boolean if a field has been set.

### SetCalendarIdsToSyncNil

`func (o *CreateCalendarParams) SetCalendarIdsToSyncNil(b bool)`

 SetCalendarIdsToSyncNil sets the value for CalendarIdsToSync to be an explicit nil

### UnsetCalendarIdsToSync
`func (o *CreateCalendarParams) UnsetCalendarIdsToSync()`

UnsetCalendarIdsToSync ensures that no value is present for CalendarIdsToSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


