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
)

// checks if the VideoOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoOptions{}

// VideoOptions struct for VideoOptions
type VideoOptions struct {
	Duration NullableInt32 `json:"duration,omitempty"`
	Transition NullableInt32 `json:"transition,omitempty"`
	TransitionType *VideoTransitionType `json:"transitionType,omitempty"`
	ResolutionType *VideoResolutionType `json:"resolutionType,omitempty"`
}

// NewVideoOptions instantiates a new VideoOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoOptions() *VideoOptions {
	this := VideoOptions{}
	return &this
}

// NewVideoOptionsWithDefaults instantiates a new VideoOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoOptionsWithDefaults() *VideoOptions {
	this := VideoOptions{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoOptions) GetDuration() int32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoOptions) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *VideoOptions) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *VideoOptions) SetDuration(v int32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *VideoOptions) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *VideoOptions) UnsetDuration() {
	o.Duration.Unset()
}

// GetTransition returns the Transition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoOptions) GetTransition() int32 {
	if o == nil || IsNil(o.Transition.Get()) {
		var ret int32
		return ret
	}
	return *o.Transition.Get()
}

// GetTransitionOk returns a tuple with the Transition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoOptions) GetTransitionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transition.Get(), o.Transition.IsSet()
}

// HasTransition returns a boolean if a field has been set.
func (o *VideoOptions) HasTransition() bool {
	if o != nil && o.Transition.IsSet() {
		return true
	}

	return false
}

// SetTransition gets a reference to the given NullableInt32 and assigns it to the Transition field.
func (o *VideoOptions) SetTransition(v int32) {
	o.Transition.Set(&v)
}
// SetTransitionNil sets the value for Transition to be an explicit nil
func (o *VideoOptions) SetTransitionNil() {
	o.Transition.Set(nil)
}

// UnsetTransition ensures that no value is present for Transition, not even an explicit nil
func (o *VideoOptions) UnsetTransition() {
	o.Transition.Unset()
}

// GetTransitionType returns the TransitionType field value if set, zero value otherwise.
func (o *VideoOptions) GetTransitionType() VideoTransitionType {
	if o == nil || IsNil(o.TransitionType) {
		var ret VideoTransitionType
		return ret
	}
	return *o.TransitionType
}

// GetTransitionTypeOk returns a tuple with the TransitionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoOptions) GetTransitionTypeOk() (*VideoTransitionType, bool) {
	if o == nil || IsNil(o.TransitionType) {
		return nil, false
	}
	return o.TransitionType, true
}

// HasTransitionType returns a boolean if a field has been set.
func (o *VideoOptions) HasTransitionType() bool {
	if o != nil && !IsNil(o.TransitionType) {
		return true
	}

	return false
}

// SetTransitionType gets a reference to the given VideoTransitionType and assigns it to the TransitionType field.
func (o *VideoOptions) SetTransitionType(v VideoTransitionType) {
	o.TransitionType = &v
}

// GetResolutionType returns the ResolutionType field value if set, zero value otherwise.
func (o *VideoOptions) GetResolutionType() VideoResolutionType {
	if o == nil || IsNil(o.ResolutionType) {
		var ret VideoResolutionType
		return ret
	}
	return *o.ResolutionType
}

// GetResolutionTypeOk returns a tuple with the ResolutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoOptions) GetResolutionTypeOk() (*VideoResolutionType, bool) {
	if o == nil || IsNil(o.ResolutionType) {
		return nil, false
	}
	return o.ResolutionType, true
}

// HasResolutionType returns a boolean if a field has been set.
func (o *VideoOptions) HasResolutionType() bool {
	if o != nil && !IsNil(o.ResolutionType) {
		return true
	}

	return false
}

// SetResolutionType gets a reference to the given VideoResolutionType and assigns it to the ResolutionType field.
func (o *VideoOptions) SetResolutionType(v VideoResolutionType) {
	o.ResolutionType = &v
}

func (o VideoOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.Transition.IsSet() {
		toSerialize["transition"] = o.Transition.Get()
	}
	if !IsNil(o.TransitionType) {
		toSerialize["transitionType"] = o.TransitionType
	}
	if !IsNil(o.ResolutionType) {
		toSerialize["resolutionType"] = o.ResolutionType
	}
	return toSerialize, nil
}

type NullableVideoOptions struct {
	value *VideoOptions
	isSet bool
}

func (v NullableVideoOptions) Get() *VideoOptions {
	return v.value
}

func (v *NullableVideoOptions) Set(val *VideoOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoOptions(val *VideoOptions) *NullableVideoOptions {
	return &NullableVideoOptions{value: val, isSet: true}
}

func (v NullableVideoOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


