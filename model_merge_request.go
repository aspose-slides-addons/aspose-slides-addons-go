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
	"os"
)

// checks if the MergeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeRequest{}

// MergeRequest struct for MergeRequest
type MergeRequest struct {
	Documents []*os.File `json:"documents,omitempty"`
	Options *MergeOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MergeRequest MergeRequest

// NewMergeRequest instantiates a new MergeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeRequest() *MergeRequest {
	this := MergeRequest{}
	return &this
}

// NewMergeRequestWithDefaults instantiates a new MergeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeRequestWithDefaults() *MergeRequest {
	this := MergeRequest{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *MergeRequest) GetDocuments() []*os.File {
	if o == nil || IsNil(o.Documents) {
		var ret []*os.File
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequest) GetDocumentsOk() ([]*os.File, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *MergeRequest) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []*os.File and assigns it to the Documents field.
func (o *MergeRequest) SetDocuments(v []*os.File) {
	o.Documents = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *MergeRequest) GetOptions() MergeOptions {
	if o == nil || IsNil(o.Options) {
		var ret MergeOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequest) GetOptionsOk() (*MergeOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MergeRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given MergeOptions and assigns it to the Options field.
func (o *MergeRequest) SetOptions(v MergeOptions) {
	o.Options = &v
}

func (o MergeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MergeRequest) UnmarshalJSON(bytes []byte) (err error) {
	varMergeRequest := _MergeRequest{}

	if err = json.Unmarshal(bytes, &varMergeRequest); err == nil {
		*o = MergeRequest(varMergeRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "documents")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMergeRequest struct {
	value *MergeRequest
	isSet bool
}

func (v NullableMergeRequest) Get() *MergeRequest {
	return v.value
}

func (v *NullableMergeRequest) Set(val *MergeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeRequest(val *MergeRequest) *NullableMergeRequest {
	return &NullableMergeRequest{value: val, isSet: true}
}

func (v NullableMergeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


