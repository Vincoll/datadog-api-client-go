/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// HttpLogError struct for HttpLogError
type HttpLogError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the Code field value
func (o *HttpLogError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// SetCode sets field value
func (o *HttpLogError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *HttpLogError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// SetMessage sets field value
func (o *HttpLogError) SetMessage(v string) {
	o.Message = v
}

type NullableHttpLogError struct {
	Value        HttpLogError
	ExplicitNull bool
}

func (v NullableHttpLogError) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableHttpLogError) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}