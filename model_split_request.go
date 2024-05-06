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

// checks if the SplitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplitRequest{}

// SplitRequest struct for SplitRequest
type SplitRequest struct {
	Document **os.File `json:"document,omitempty"`
	Options *SplitOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SplitRequest SplitRequest

// NewSplitRequest instantiates a new SplitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitRequest() *SplitRequest {
	this := SplitRequest{}
	return &this
}

// NewSplitRequestWithDefaults instantiates a new SplitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitRequestWithDefaults() *SplitRequest {
	this := SplitRequest{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *SplitRequest) GetDocument() *os.File {
	if o == nil || IsNil(o.Document) {
		var ret *os.File
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitRequest) GetDocumentOk() (**os.File, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *SplitRequest) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given *os.File and assigns it to the Document field.
func (o *SplitRequest) SetDocument(v *os.File) {
	o.Document = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SplitRequest) GetOptions() SplitOptions {
	if o == nil || IsNil(o.Options) {
		var ret SplitOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitRequest) GetOptionsOk() (*SplitOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SplitRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SplitOptions and assigns it to the Options field.
func (o *SplitRequest) SetOptions(v SplitOptions) {
	o.Options = &v
}

func (o SplitRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitRequest) ToMap() (map[string]interface{}, error) {
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

func (o *SplitRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSplitRequest := _SplitRequest{}

	if err = json.Unmarshal(bytes, &varSplitRequest); err == nil {
		*o = SplitRequest(varSplitRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSplitRequest struct {
	value *SplitRequest
	isSet bool
}

func (v NullableSplitRequest) Get() *SplitRequest {
	return v.value
}

func (v *NullableSplitRequest) Set(val *SplitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitRequest(val *SplitRequest) *NullableSplitRequest {
	return &NullableSplitRequest{value: val, isSet: true}
}

func (v NullableSplitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


