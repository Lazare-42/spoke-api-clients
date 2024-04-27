# SystemTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecsSinceEpoch** | **int32** |  | 
**NanosSinceEpoch** | **int32** |  | 

## Methods

### NewSystemTime

`func NewSystemTime(secsSinceEpoch int32, nanosSinceEpoch int32, ) *SystemTime`

NewSystemTime instantiates a new SystemTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemTimeWithDefaults

`func NewSystemTimeWithDefaults() *SystemTime`

NewSystemTimeWithDefaults instantiates a new SystemTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecsSinceEpoch

`func (o *SystemTime) GetSecsSinceEpoch() int32`

GetSecsSinceEpoch returns the SecsSinceEpoch field if non-nil, zero value otherwise.

### GetSecsSinceEpochOk

`func (o *SystemTime) GetSecsSinceEpochOk() (*int32, bool)`

GetSecsSinceEpochOk returns a tuple with the SecsSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecsSinceEpoch

`func (o *SystemTime) SetSecsSinceEpoch(v int32)`

SetSecsSinceEpoch sets SecsSinceEpoch field to given value.


### GetNanosSinceEpoch

`func (o *SystemTime) GetNanosSinceEpoch() int32`

GetNanosSinceEpoch returns the NanosSinceEpoch field if non-nil, zero value otherwise.

### GetNanosSinceEpochOk

`func (o *SystemTime) GetNanosSinceEpochOk() (*int32, bool)`

GetNanosSinceEpochOk returns a tuple with the NanosSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanosSinceEpoch

`func (o *SystemTime) SetNanosSinceEpoch(v int32)`

SetNanosSinceEpoch sets NanosSinceEpoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


