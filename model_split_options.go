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

// checks if the SplitOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplitOptions{}

// SplitOptions struct for SplitOptions
type SplitOptions struct {
	SlidesRange NullableString `json:"slidesRange,omitempty"`
}

// NewSplitOptions instantiates a new SplitOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitOptions() *SplitOptions {
	this := SplitOptions{}
	return &this
}

// NewSplitOptionsWithDefaults instantiates a new SplitOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitOptionsWithDefaults() *SplitOptions {
	this := SplitOptions{}
	return &this
}

// GetSlidesRange returns the SlidesRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SplitOptions) GetSlidesRange() string {
	if o == nil || IsNil(o.SlidesRange.Get()) {
		var ret string
		return ret
	}
	return *o.SlidesRange.Get()
}

// GetSlidesRangeOk returns a tuple with the SlidesRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SplitOptions) GetSlidesRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlidesRange.Get(), o.SlidesRange.IsSet()
}

// HasSlidesRange returns a boolean if a field has been set.
func (o *SplitOptions) HasSlidesRange() bool {
	if o != nil && o.SlidesRange.IsSet() {
		return true
	}

	return false
}

// SetSlidesRange gets a reference to the given NullableString and assigns it to the SlidesRange field.
func (o *SplitOptions) SetSlidesRange(v string) {
	o.SlidesRange.Set(&v)
}
// SetSlidesRangeNil sets the value for SlidesRange to be an explicit nil
func (o *SplitOptions) SetSlidesRangeNil() {
	o.SlidesRange.Set(nil)
}

// UnsetSlidesRange ensures that no value is present for SlidesRange, not even an explicit nil
func (o *SplitOptions) UnsetSlidesRange() {
	o.SlidesRange.Unset()
}

func (o SplitOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SlidesRange.IsSet() {
		toSerialize["slidesRange"] = o.SlidesRange.Get()
	}
	return toSerialize, nil
}

type NullableSplitOptions struct {
	value *SplitOptions
	isSet bool
}

func (v NullableSplitOptions) Get() *SplitOptions {
	return v.value
}

func (v *NullableSplitOptions) Set(val *SplitOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitOptions(val *SplitOptions) *NullableSplitOptions {
	return &NullableSplitOptions{value: val, isSet: true}
}

func (v NullableSplitOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


