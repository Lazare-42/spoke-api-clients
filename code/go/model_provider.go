/*
spoke_handler

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Provider Fields with value `\"simple\"` parse as `Kind::Simple`. Fields with value `\"fancy\"` parse as `Kind::SoFancy`.
type Provider string

// List of Provider
const (
	GOOGLE Provider = "Google"
	MICROSOFT Provider = "Microsoft"
)

// All allowed values of Provider enum
var AllowedProviderEnumValues = []Provider{
	"Google",
	"Microsoft",
}

func (v *Provider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Provider(value)
	for _, existing := range AllowedProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Provider", value)
}

// NewProviderFromValue returns a pointer to a valid Provider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderFromValue(v string) (*Provider, error) {
	ev := Provider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Provider: valid values are %v", v, AllowedProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Provider) IsValid() bool {
	for _, existing := range AllowedProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Provider value
func (v Provider) Ptr() *Provider {
	return &v
}

type NullableProvider struct {
	value *Provider
	isSet bool
}

func (v NullableProvider) Get() *Provider {
	return v.value
}

func (v *NullableProvider) Set(val *Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvider(val *Provider) *NullableProvider {
	return &NullableProvider{value: val, isSet: true}
}

func (v NullableProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

