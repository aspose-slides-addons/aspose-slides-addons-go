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

// checks if the ProtectionOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtectionOptions{}

// ProtectionOptions struct for ProtectionOptions
type ProtectionOptions struct {
	ViewPassword NullableString `json:"viewPassword,omitempty"`
	EditPassword NullableString `json:"editPassword,omitempty"`
	MarkAsFinal *bool `json:"markAsFinal,omitempty"`
}

// NewProtectionOptions instantiates a new ProtectionOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectionOptions() *ProtectionOptions {
	this := ProtectionOptions{}
	return &this
}

// NewProtectionOptionsWithDefaults instantiates a new ProtectionOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectionOptionsWithDefaults() *ProtectionOptions {
	this := ProtectionOptions{}
	return &this
}

// GetViewPassword returns the ViewPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionOptions) GetViewPassword() string {
	if o == nil || IsNil(o.ViewPassword.Get()) {
		var ret string
		return ret
	}
	return *o.ViewPassword.Get()
}

// GetViewPasswordOk returns a tuple with the ViewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionOptions) GetViewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewPassword.Get(), o.ViewPassword.IsSet()
}

// HasViewPassword returns a boolean if a field has been set.
func (o *ProtectionOptions) HasViewPassword() bool {
	if o != nil && o.ViewPassword.IsSet() {
		return true
	}

	return false
}

// SetViewPassword gets a reference to the given NullableString and assigns it to the ViewPassword field.
func (o *ProtectionOptions) SetViewPassword(v string) {
	o.ViewPassword.Set(&v)
}
// SetViewPasswordNil sets the value for ViewPassword to be an explicit nil
func (o *ProtectionOptions) SetViewPasswordNil() {
	o.ViewPassword.Set(nil)
}

// UnsetViewPassword ensures that no value is present for ViewPassword, not even an explicit nil
func (o *ProtectionOptions) UnsetViewPassword() {
	o.ViewPassword.Unset()
}

// GetEditPassword returns the EditPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionOptions) GetEditPassword() string {
	if o == nil || IsNil(o.EditPassword.Get()) {
		var ret string
		return ret
	}
	return *o.EditPassword.Get()
}

// GetEditPasswordOk returns a tuple with the EditPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionOptions) GetEditPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditPassword.Get(), o.EditPassword.IsSet()
}

// HasEditPassword returns a boolean if a field has been set.
func (o *ProtectionOptions) HasEditPassword() bool {
	if o != nil && o.EditPassword.IsSet() {
		return true
	}

	return false
}

// SetEditPassword gets a reference to the given NullableString and assigns it to the EditPassword field.
func (o *ProtectionOptions) SetEditPassword(v string) {
	o.EditPassword.Set(&v)
}
// SetEditPasswordNil sets the value for EditPassword to be an explicit nil
func (o *ProtectionOptions) SetEditPasswordNil() {
	o.EditPassword.Set(nil)
}

// UnsetEditPassword ensures that no value is present for EditPassword, not even an explicit nil
func (o *ProtectionOptions) UnsetEditPassword() {
	o.EditPassword.Unset()
}

// GetMarkAsFinal returns the MarkAsFinal field value if set, zero value otherwise.
func (o *ProtectionOptions) GetMarkAsFinal() bool {
	if o == nil || IsNil(o.MarkAsFinal) {
		var ret bool
		return ret
	}
	return *o.MarkAsFinal
}

// GetMarkAsFinalOk returns a tuple with the MarkAsFinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectionOptions) GetMarkAsFinalOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkAsFinal) {
		return nil, false
	}
	return o.MarkAsFinal, true
}

// HasMarkAsFinal returns a boolean if a field has been set.
func (o *ProtectionOptions) HasMarkAsFinal() bool {
	if o != nil && !IsNil(o.MarkAsFinal) {
		return true
	}

	return false
}

// SetMarkAsFinal gets a reference to the given bool and assigns it to the MarkAsFinal field.
func (o *ProtectionOptions) SetMarkAsFinal(v bool) {
	o.MarkAsFinal = &v
}

func (o ProtectionOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtectionOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewPassword.IsSet() {
		toSerialize["viewPassword"] = o.ViewPassword.Get()
	}
	if o.EditPassword.IsSet() {
		toSerialize["editPassword"] = o.EditPassword.Get()
	}
	if !IsNil(o.MarkAsFinal) {
		toSerialize["markAsFinal"] = o.MarkAsFinal
	}
	return toSerialize, nil
}

type NullableProtectionOptions struct {
	value *ProtectionOptions
	isSet bool
}

func (v NullableProtectionOptions) Get() *ProtectionOptions {
	return v.value
}

func (v *NullableProtectionOptions) Set(val *ProtectionOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionOptions(val *ProtectionOptions) *NullableProtectionOptions {
	return &NullableProtectionOptions{value: val, isSet: true}
}

func (v NullableProtectionOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


