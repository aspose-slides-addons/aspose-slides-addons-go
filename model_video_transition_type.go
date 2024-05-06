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

// VideoTransitionType the model 'VideoTransitionType'
type VideoTransitionType string

// List of VideoTransitionType
const (
	NONE VideoTransitionType = "None"
	RANDOM VideoTransitionType = "Random"
	FROM_PRESENTATION VideoTransitionType = "FromPresentation"
	FADE VideoTransitionType = "Fade"
	DISTANCE VideoTransitionType = "Distance"
	SLIDE_LEFT VideoTransitionType = "SlideLeft"
	CIRCLE_CROP VideoTransitionType = "CircleCrop"
	DISSOLVE VideoTransitionType = "Dissolve"
)

// All allowed values of VideoTransitionType enum
var AllowedVideoTransitionTypeEnumValues = []VideoTransitionType{
	"None",
	"Random",
	"FromPresentation",
	"Fade",
	"Distance",
	"SlideLeft",
	"CircleCrop",
	"Dissolve",
}

func (v *VideoTransitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VideoTransitionType(value)
	for _, existing := range AllowedVideoTransitionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VideoTransitionType", value)
}

// NewVideoTransitionTypeFromValue returns a pointer to a valid VideoTransitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVideoTransitionTypeFromValue(v string) (*VideoTransitionType, error) {
	ev := VideoTransitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VideoTransitionType: valid values are %v", v, AllowedVideoTransitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VideoTransitionType) IsValid() bool {
	for _, existing := range AllowedVideoTransitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VideoTransitionType value
func (v VideoTransitionType) Ptr() *VideoTransitionType {
	return &v
}

type NullableVideoTransitionType struct {
	value *VideoTransitionType
	isSet bool
}

func (v NullableVideoTransitionType) Get() *VideoTransitionType {
	return v.value
}

func (v *NullableVideoTransitionType) Set(val *VideoTransitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoTransitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoTransitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoTransitionType(val *VideoTransitionType) *NullableVideoTransitionType {
	return &NullableVideoTransitionType{value: val, isSet: true}
}

func (v NullableVideoTransitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoTransitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

