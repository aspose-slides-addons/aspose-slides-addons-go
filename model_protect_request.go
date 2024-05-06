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

// checks if the ProtectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtectRequest{}

// ProtectRequest struct for ProtectRequest
type ProtectRequest struct {
	Document **os.File `json:"document,omitempty"`
	Options *ProtectionOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtectRequest ProtectRequest

// NewProtectRequest instantiates a new ProtectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectRequest() *ProtectRequest {
	this := ProtectRequest{}
	return &this
}

// NewProtectRequestWithDefaults instantiates a new ProtectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectRequestWithDefaults() *ProtectRequest {
	this := ProtectRequest{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *ProtectRequest) GetDocument() *os.File {
	if o == nil || IsNil(o.Document) {
		var ret *os.File
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectRequest) GetDocumentOk() (**os.File, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *ProtectRequest) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given *os.File and assigns it to the Document field.
func (o *ProtectRequest) SetDocument(v *os.File) {
	o.Document = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProtectRequest) GetOptions() ProtectionOptions {
	if o == nil || IsNil(o.Options) {
		var ret ProtectionOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectRequest) GetOptionsOk() (*ProtectionOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProtectRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ProtectionOptions and assigns it to the Options field.
func (o *ProtectRequest) SetOptions(v ProtectionOptions) {
	o.Options = &v
}

func (o ProtectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtectRequest) ToMap() (map[string]interface{}, error) {
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

func (o *ProtectRequest) UnmarshalJSON(bytes []byte) (err error) {
	varProtectRequest := _ProtectRequest{}

	if err = json.Unmarshal(bytes, &varProtectRequest); err == nil {
		*o = ProtectRequest(varProtectRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtectRequest struct {
	value *ProtectRequest
	isSet bool
}

func (v NullableProtectRequest) Get() *ProtectRequest {
	return v.value
}

func (v *NullableProtectRequest) Set(val *ProtectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectRequest(val *ProtectRequest) *NullableProtectRequest {
	return &NullableProtectRequest{value: val, isSet: true}
}

func (v NullableProtectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


