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

// checks if the ReplaceTextOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceTextOptions{}

// ReplaceTextOptions struct for ReplaceTextOptions
type ReplaceTextOptions struct {
	OldValue NullableString `json:"oldValue,omitempty"`
	NewValue NullableString `json:"newValue,omitempty"`
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
}

// NewReplaceTextOptions instantiates a new ReplaceTextOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceTextOptions() *ReplaceTextOptions {
	this := ReplaceTextOptions{}
	return &this
}

// NewReplaceTextOptionsWithDefaults instantiates a new ReplaceTextOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceTextOptionsWithDefaults() *ReplaceTextOptions {
	this := ReplaceTextOptions{}
	return &this
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplaceTextOptions) GetOldValue() string {
	if o == nil || IsNil(o.OldValue.Get()) {
		var ret string
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplaceTextOptions) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *ReplaceTextOptions) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableString and assigns it to the OldValue field.
func (o *ReplaceTextOptions) SetOldValue(v string) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *ReplaceTextOptions) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *ReplaceTextOptions) UnsetOldValue() {
	o.OldValue.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplaceTextOptions) GetNewValue() string {
	if o == nil || IsNil(o.NewValue.Get()) {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplaceTextOptions) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *ReplaceTextOptions) HasNewValue() bool {
	if o != nil && o.NewValue.IsSet() {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given NullableString and assigns it to the NewValue field.
func (o *ReplaceTextOptions) SetNewValue(v string) {
	o.NewValue.Set(&v)
}
// SetNewValueNil sets the value for NewValue to be an explicit nil
func (o *ReplaceTextOptions) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
func (o *ReplaceTextOptions) UnsetNewValue() {
	o.NewValue.Unset()
}

// GetIgnoreCase returns the IgnoreCase field value if set, zero value otherwise.
func (o *ReplaceTextOptions) GetIgnoreCase() bool {
	if o == nil || IsNil(o.IgnoreCase) {
		var ret bool
		return ret
	}
	return *o.IgnoreCase
}

// GetIgnoreCaseOk returns a tuple with the IgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceTextOptions) GetIgnoreCaseOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreCase) {
		return nil, false
	}
	return o.IgnoreCase, true
}

// HasIgnoreCase returns a boolean if a field has been set.
func (o *ReplaceTextOptions) HasIgnoreCase() bool {
	if o != nil && !IsNil(o.IgnoreCase) {
		return true
	}

	return false
}

// SetIgnoreCase gets a reference to the given bool and assigns it to the IgnoreCase field.
func (o *ReplaceTextOptions) SetIgnoreCase(v bool) {
	o.IgnoreCase = &v
}

func (o ReplaceTextOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceTextOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OldValue.IsSet() {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	if o.NewValue.IsSet() {
		toSerialize["newValue"] = o.NewValue.Get()
	}
	if !IsNil(o.IgnoreCase) {
		toSerialize["ignoreCase"] = o.IgnoreCase
	}
	return toSerialize, nil
}

type NullableReplaceTextOptions struct {
	value *ReplaceTextOptions
	isSet bool
}

func (v NullableReplaceTextOptions) Get() *ReplaceTextOptions {
	return v.value
}

func (v *NullableReplaceTextOptions) Set(val *ReplaceTextOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceTextOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceTextOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceTextOptions(val *ReplaceTextOptions) *NullableReplaceTextOptions {
	return &NullableReplaceTextOptions{value: val, isSet: true}
}

func (v NullableReplaceTextOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceTextOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


