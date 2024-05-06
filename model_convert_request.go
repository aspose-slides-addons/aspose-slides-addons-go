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

// checks if the ConvertRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertRequest{}

// ConvertRequest struct for ConvertRequest
type ConvertRequest struct {
	Documents []*os.File `json:"documents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvertRequest ConvertRequest

// NewConvertRequest instantiates a new ConvertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertRequest() *ConvertRequest {
	this := ConvertRequest{}
	return &this
}

// NewConvertRequestWithDefaults instantiates a new ConvertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertRequestWithDefaults() *ConvertRequest {
	this := ConvertRequest{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *ConvertRequest) GetDocuments() []*os.File {
	if o == nil || IsNil(o.Documents) {
		var ret []*os.File
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertRequest) GetDocumentsOk() ([]*os.File, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *ConvertRequest) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []*os.File and assigns it to the Documents field.
func (o *ConvertRequest) SetDocuments(v []*os.File) {
	o.Documents = v
}

func (o ConvertRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConvertRequest) UnmarshalJSON(bytes []byte) (err error) {
	varConvertRequest := _ConvertRequest{}

	if err = json.Unmarshal(bytes, &varConvertRequest); err == nil {
		*o = ConvertRequest(varConvertRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "documents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvertRequest struct {
	value *ConvertRequest
	isSet bool
}

func (v NullableConvertRequest) Get() *ConvertRequest {
	return v.value
}

func (v *NullableConvertRequest) Set(val *ConvertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertRequest(val *ConvertRequest) *NullableConvertRequest {
	return &NullableConvertRequest{value: val, isSet: true}
}

func (v NullableConvertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


