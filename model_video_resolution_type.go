/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2024 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */
package asposeslidesaddons

import (
	"encoding/json"
	"fmt"
)

// VideoResolutionType the model 'VideoResolutionType'
type VideoResolutionType string

// List of VideoResolutionType
const (
	SD VideoResolutionType = "SD"
	HD VideoResolutionType = "HD"
	FULL_HD VideoResolutionType = "FullHD"
	QHD VideoResolutionType = "QHD"
)

// All allowed values of VideoResolutionType enum
var AllowedVideoResolutionTypeEnumValues = []VideoResolutionType{
	"SD",
	"HD",
	"FullHD",
	"QHD",
}

func (v *VideoResolutionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VideoResolutionType(value)
	for _, existing := range AllowedVideoResolutionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VideoResolutionType", value)
}

// NewVideoResolutionTypeFromValue returns a pointer to a valid VideoResolutionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVideoResolutionTypeFromValue(v string) (*VideoResolutionType, error) {
	ev := VideoResolutionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VideoResolutionType: valid values are %v", v, AllowedVideoResolutionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VideoResolutionType) IsValid() bool {
	for _, existing := range AllowedVideoResolutionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VideoResolutionType value
func (v VideoResolutionType) Ptr() *VideoResolutionType {
	return &v
}

type NullableVideoResolutionType struct {
	value *VideoResolutionType
	isSet bool
}

func (v NullableVideoResolutionType) Get() *VideoResolutionType {
	return v.value
}

func (v *NullableVideoResolutionType) Set(val *VideoResolutionType) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoResolutionType) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoResolutionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoResolutionType(val *VideoResolutionType) *NullableVideoResolutionType {
	return &NullableVideoResolutionType{value: val, isSet: true}
}

func (v NullableVideoResolutionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoResolutionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

