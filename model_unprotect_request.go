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

// checks if the UnprotectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnprotectRequest{}

// UnprotectRequest struct for UnprotectRequest
type UnprotectRequest struct {
	Document **os.File `json:"document,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UnprotectRequest UnprotectRequest

// NewUnprotectRequest instantiates a new UnprotectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnprotectRequest() *UnprotectRequest {
	this := UnprotectRequest{}
	return &this
}

// NewUnprotectRequestWithDefaults instantiates a new UnprotectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnprotectRequestWithDefaults() *UnprotectRequest {
	this := UnprotectRequest{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *UnprotectRequest) GetDocument() *os.File {
	if o == nil || IsNil(o.Document) {
		var ret *os.File
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprotectRequest) GetDocumentOk() (**os.File, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *UnprotectRequest) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given *os.File and assigns it to the Document field.
func (o *UnprotectRequest) SetDocument(v *os.File) {
	o.Document = &v
}

func (o UnprotectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnprotectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UnprotectRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUnprotectRequest := _UnprotectRequest{}

	if err = json.Unmarshal(bytes, &varUnprotectRequest); err == nil {
		*o = UnprotectRequest(varUnprotectRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnprotectRequest struct {
	value *UnprotectRequest
	isSet bool
}

func (v NullableUnprotectRequest) Get() *UnprotectRequest {
	return v.value
}

func (v *NullableUnprotectRequest) Set(val *UnprotectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnprotectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnprotectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnprotectRequest(val *UnprotectRequest) *NullableUnprotectRequest {
	return &NullableUnprotectRequest{value: val, isSet: true}
}

func (v NullableUnprotectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnprotectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


