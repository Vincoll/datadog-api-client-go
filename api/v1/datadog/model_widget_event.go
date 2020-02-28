/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// WidgetEvent struct for WidgetEvent
type WidgetEvent struct {
	Q string `json:"q"`
	// The execution method for multi-value filters. Can be either and or or
	TagsExecution *string `json:"tags_execution,omitempty"`
}

// NewWidgetEvent instantiates a new WidgetEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetEvent(q string) *WidgetEvent {
	this := WidgetEvent{}
	this.Q = q
	return &this
}

// NewWidgetEventWithDefaults instantiates a new WidgetEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetEventWithDefaults() *WidgetEvent {
	this := WidgetEvent{}
	return &this
}

// GetQ returns the Q field value
func (o *WidgetEvent) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Q
}

// SetQ sets field value
func (o *WidgetEvent) SetQ(v string) {
	o.Q = v
}

// GetTagsExecution returns the TagsExecution field value if set, zero value otherwise.
func (o *WidgetEvent) GetTagsExecution() string {
	if o == nil || o.TagsExecution == nil {
		var ret string
		return ret
	}
	return *o.TagsExecution
}

// GetTagsExecutionOk returns a tuple with the TagsExecution field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WidgetEvent) GetTagsExecutionOk() (string, bool) {
	if o == nil || o.TagsExecution == nil {
		var ret string
		return ret, false
	}
	return *o.TagsExecution, true
}

// HasTagsExecution returns a boolean if a field has been set.
func (o *WidgetEvent) HasTagsExecution() bool {
	if o != nil && o.TagsExecution != nil {
		return true
	}

	return false
}

// SetTagsExecution gets a reference to the given string and assigns it to the TagsExecution field.
func (o *WidgetEvent) SetTagsExecution(v string) {
	o.TagsExecution = &v
}

func (o WidgetEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["q"] = o.Q
	}
	if o.TagsExecution != nil {
		toSerialize["tags_execution"] = o.TagsExecution
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetEvent struct {
	value *WidgetEvent
	isSet bool
}

func (v NullableWidgetEvent) Get() *WidgetEvent {
	return v.value
}

func (v NullableWidgetEvent) Set(val *WidgetEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetEvent) IsSet() bool {
	return v.isSet
}

func (v NullableWidgetEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetEvent(val *WidgetEvent) *NullableWidgetEvent {
	return &NullableWidgetEvent{value: val, isSet: true}
}

func (v NullableWidgetEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}