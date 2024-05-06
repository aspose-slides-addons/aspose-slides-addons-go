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

// checks if the ConvertToVideoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertToVideoRequest{}

// ConvertToVideoRequest struct for ConvertToVideoRequest
type ConvertToVideoRequest struct {
	Document **os.File `json:"document,omitempty"`
	Options *VideoOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvertToVideoRequest ConvertToVideoRequest

// NewConvertToVideoRequest instantiates a new ConvertToVideoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertToVideoRequest() *ConvertToVideoRequest {
	this := ConvertToVideoRequest{}
	return &this
}

// NewConvertToVideoRequestWithDefaults instantiates a new ConvertToVideoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertToVideoRequestWithDefaults() *ConvertToVideoRequest {
	this := ConvertToVideoRequest{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *ConvertToVideoRequest) GetDocument() *os.File {
	if o == nil || IsNil(o.Document) {
		var ret *os.File
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertToVideoRequest) GetDocumentOk() (**os.File, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *ConvertToVideoRequest) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given *os.File and assigns it to the Document field.
func (o *ConvertToVideoRequest) SetDocument(v *os.File) {
	o.Document = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ConvertToVideoRequest) GetOptions() VideoOptions {
	if o == nil || IsNil(o.Options) {
		var ret VideoOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertToVideoRequest) GetOptionsOk() (*VideoOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ConvertToVideoRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given VideoOptions and assigns it to the Options field.
func (o *ConvertToVideoRequest) SetOptions(v VideoOptions) {
	o.Options = &v
}

func (o ConvertToVideoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertToVideoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConvertToVideoRequest) UnmarshalJSON(bytes []byte) (err error) {
	varConvertToVideoRequest := _ConvertToVideoRequest{}

	if err = json.Unmarshal(bytes, &varConvertToVideoRequest); err == nil {
		*o = ConvertToVideoRequest(varConvertToVideoRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvertToVideoRequest struct {
	value *ConvertToVideoRequest
	isSet bool
}

func (v NullableConvertToVideoRequest) Get() *ConvertToVideoRequest {
	return v.value
}

func (v *NullableConvertToVideoRequest) Set(val *ConvertToVideoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertToVideoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertToVideoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertToVideoRequest(val *ConvertToVideoRequest) *NullableConvertToVideoRequest {
	return &NullableConvertToVideoRequest{value: val, isSet: true}
}

func (v NullableConvertToVideoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertToVideoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


