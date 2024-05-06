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

// checks if the TextWatermarkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextWatermarkRequest{}

// TextWatermarkRequest struct for TextWatermarkRequest
type TextWatermarkRequest struct {
	Documents []*os.File `json:"documents,omitempty"`
	Options *TextWatermarkOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TextWatermarkRequest TextWatermarkRequest

// NewTextWatermarkRequest instantiates a new TextWatermarkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextWatermarkRequest() *TextWatermarkRequest {
	this := TextWatermarkRequest{}
	return &this
}

// NewTextWatermarkRequestWithDefaults instantiates a new TextWatermarkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextWatermarkRequestWithDefaults() *TextWatermarkRequest {
	this := TextWatermarkRequest{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *TextWatermarkRequest) GetDocuments() []*os.File {
	if o == nil || IsNil(o.Documents) {
		var ret []*os.File
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextWatermarkRequest) GetDocumentsOk() ([]*os.File, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *TextWatermarkRequest) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []*os.File and assigns it to the Documents field.
func (o *TextWatermarkRequest) SetDocuments(v []*os.File) {
	o.Documents = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TextWatermarkRequest) GetOptions() TextWatermarkOptions {
	if o == nil || IsNil(o.Options) {
		var ret TextWatermarkOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextWatermarkRequest) GetOptionsOk() (*TextWatermarkOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TextWatermarkRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TextWatermarkOptions and assigns it to the Options field.
func (o *TextWatermarkRequest) SetOptions(v TextWatermarkOptions) {
	o.Options = &v
}

func (o TextWatermarkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextWatermarkRequest) ToMap() (map[string]interface{}, error) {
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

func (o *TextWatermarkRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTextWatermarkRequest := _TextWatermarkRequest{}

	if err = json.Unmarshal(bytes, &varTextWatermarkRequest); err == nil {
		*o = TextWatermarkRequest(varTextWatermarkRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "documents")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTextWatermarkRequest struct {
	value *TextWatermarkRequest
	isSet bool
}

func (v NullableTextWatermarkRequest) Get() *TextWatermarkRequest {
	return v.value
}

func (v *NullableTextWatermarkRequest) Set(val *TextWatermarkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTextWatermarkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTextWatermarkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextWatermarkRequest(val *TextWatermarkRequest) *NullableTextWatermarkRequest {
	return &NullableTextWatermarkRequest{value: val, isSet: true}
}

func (v NullableTextWatermarkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextWatermarkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


