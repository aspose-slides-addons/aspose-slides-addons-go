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

// checks if the MergeOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeOptions{}

// MergeOptions struct for MergeOptions
type MergeOptions struct {
	MasterFileName NullableString `json:"masterFileName,omitempty"`
	ExcludeMasterFile *bool `json:"excludeMasterFile,omitempty"`
}

// NewMergeOptions instantiates a new MergeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeOptions() *MergeOptions {
	this := MergeOptions{}
	return &this
}

// NewMergeOptionsWithDefaults instantiates a new MergeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeOptionsWithDefaults() *MergeOptions {
	this := MergeOptions{}
	return &this
}

// GetMasterFileName returns the MasterFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeOptions) GetMasterFileName() string {
	if o == nil || IsNil(o.MasterFileName.Get()) {
		var ret string
		return ret
	}
	return *o.MasterFileName.Get()
}

// GetMasterFileNameOk returns a tuple with the MasterFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeOptions) GetMasterFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MasterFileName.Get(), o.MasterFileName.IsSet()
}

// HasMasterFileName returns a boolean if a field has been set.
func (o *MergeOptions) HasMasterFileName() bool {
	if o != nil && o.MasterFileName.IsSet() {
		return true
	}

	return false
}

// SetMasterFileName gets a reference to the given NullableString and assigns it to the MasterFileName field.
func (o *MergeOptions) SetMasterFileName(v string) {
	o.MasterFileName.Set(&v)
}
// SetMasterFileNameNil sets the value for MasterFileName to be an explicit nil
func (o *MergeOptions) SetMasterFileNameNil() {
	o.MasterFileName.Set(nil)
}

// UnsetMasterFileName ensures that no value is present for MasterFileName, not even an explicit nil
func (o *MergeOptions) UnsetMasterFileName() {
	o.MasterFileName.Unset()
}

// GetExcludeMasterFile returns the ExcludeMasterFile field value if set, zero value otherwise.
func (o *MergeOptions) GetExcludeMasterFile() bool {
	if o == nil || IsNil(o.ExcludeMasterFile) {
		var ret bool
		return ret
	}
	return *o.ExcludeMasterFile
}

// GetExcludeMasterFileOk returns a tuple with the ExcludeMasterFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeOptions) GetExcludeMasterFileOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeMasterFile) {
		return nil, false
	}
	return o.ExcludeMasterFile, true
}

// HasExcludeMasterFile returns a boolean if a field has been set.
func (o *MergeOptions) HasExcludeMasterFile() bool {
	if o != nil && !IsNil(o.ExcludeMasterFile) {
		return true
	}

	return false
}

// SetExcludeMasterFile gets a reference to the given bool and assigns it to the ExcludeMasterFile field.
func (o *MergeOptions) SetExcludeMasterFile(v bool) {
	o.ExcludeMasterFile = &v
}

func (o MergeOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MasterFileName.IsSet() {
		toSerialize["masterFileName"] = o.MasterFileName.Get()
	}
	if !IsNil(o.ExcludeMasterFile) {
		toSerialize["excludeMasterFile"] = o.ExcludeMasterFile
	}
	return toSerialize, nil
}

type NullableMergeOptions struct {
	value *MergeOptions
	isSet bool
}

func (v NullableMergeOptions) Get() *MergeOptions {
	return v.value
}

func (v *NullableMergeOptions) Set(val *MergeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeOptions(val *MergeOptions) *NullableMergeOptions {
	return &NullableMergeOptions{value: val, isSet: true}
}

func (v NullableMergeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


