/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// SyntheticsErrorCode the model 'SyntheticsErrorCode'
type SyntheticsErrorCode string

// List of SyntheticsErrorCode
const (
	SYNTHETICSERRORCODE_NO_ERROR            SyntheticsErrorCode = "NO_ERROR"
	SYNTHETICSERRORCODE_UNKNOWN             SyntheticsErrorCode = "UNKNOWN"
	SYNTHETICSERRORCODE_DNS                 SyntheticsErrorCode = "DNS"
	SYNTHETICSERRORCODE_SSL                 SyntheticsErrorCode = "SSL"
	SYNTHETICSERRORCODE_TIMEOUT             SyntheticsErrorCode = "TIMEOUT"
	SYNTHETICSERRORCODE_DENIED              SyntheticsErrorCode = "DENIED"
	SYNTHETICSERRORCODE_INCORRECT_ASSERTION SyntheticsErrorCode = "INCORRECT_ASSERTION"
)

// Ptr returns reference to SyntheticsErrorCode value
func (v SyntheticsErrorCode) Ptr() *SyntheticsErrorCode {
	return &v
}

type NullableSyntheticsErrorCode struct {
	Value        SyntheticsErrorCode
	ExplicitNull bool
}

func (v NullableSyntheticsErrorCode) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsErrorCode) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}