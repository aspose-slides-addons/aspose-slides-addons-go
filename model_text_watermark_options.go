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

// checks if the TextWatermarkOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextWatermarkOptions{}

// TextWatermarkOptions struct for TextWatermarkOptions
type TextWatermarkOptions struct {
	Text NullableString `json:"text,omitempty"`
	FontName NullableString `json:"fontName,omitempty"`
	FontSize NullableInt32 `json:"fontSize,omitempty"`
	Color NullableString `json:"color,omitempty"`
	Angle NullableInt32 `json:"angle,omitempty"`
}

// NewTextWatermarkOptions instantiates a new TextWatermarkOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextWatermarkOptions() *TextWatermarkOptions {
	this := TextWatermarkOptions{}
	return &this
}

// NewTextWatermarkOptionsWithDefaults instantiates a new TextWatermarkOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextWatermarkOptionsWithDefaults() *TextWatermarkOptions {
	this := TextWatermarkOptions{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextWatermarkOptions) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextWatermarkOptions) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *TextWatermarkOptions) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *TextWatermarkOptions) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *TextWatermarkOptions) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *TextWatermarkOptions) UnsetText() {
	o.Text.Unset()
}

// GetFontName returns the FontName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextWatermarkOptions) GetFontName() string {
	if o == nil || IsNil(o.FontName.Get()) {
		var ret string
		return ret
	}
	return *o.FontName.Get()
}

// GetFontNameOk returns a tuple with the FontName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextWatermarkOptions) GetFontNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FontName.Get(), o.FontName.IsSet()
}

// HasFontName returns a boolean if a field has been set.
func (o *TextWatermarkOptions) HasFontName() bool {
	if o != nil && o.FontName.IsSet() {
		return true
	}

	return false
}

// SetFontName gets a reference to the given NullableString and assigns it to the FontName field.
func (o *TextWatermarkOptions) SetFontName(v string) {
	o.FontName.Set(&v)
}
// SetFontNameNil sets the value for FontName to be an explicit nil
func (o *TextWatermarkOptions) SetFontNameNil() {
	o.FontName.Set(nil)
}

// UnsetFontName ensures that no value is present for FontName, not even an explicit nil
func (o *TextWatermarkOptions) UnsetFontName() {
	o.FontName.Unset()
}

// GetFontSize returns the FontSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextWatermarkOptions) GetFontSize() int32 {
	if o == nil || IsNil(o.FontSize.Get()) {
		var ret int32
		return ret
	}
	return *o.FontSize.Get()
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextWatermarkOptions) GetFontSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FontSize.Get(), o.FontSize.IsSet()
}

// HasFontSize returns a boolean if a field has been set.
func (o *TextWatermarkOptions) HasFontSize() bool {
	if o != nil && o.FontSize.IsSet() {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given NullableInt32 and assigns it to the FontSize field.
func (o *TextWatermarkOptions) SetFontSize(v int32) {
	o.FontSize.Set(&v)
}
// SetFontSizeNil sets the value for FontSize to be an explicit nil
func (o *TextWatermarkOptions) SetFontSizeNil() {
	o.FontSize.Set(nil)
}

// UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
func (o *TextWatermarkOptions) UnsetFontSize() {
	o.FontSize.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextWatermarkOptions) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextWatermarkOptions) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *TextWatermarkOptions) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *TextWatermarkOptions) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *TextWatermarkOptions) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *TextWatermarkOptions) UnsetColor() {
	o.Color.Unset()
}

// GetAngle returns the Angle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextWatermarkOptions) GetAngle() int32 {
	if o == nil || IsNil(o.Angle.Get()) {
		var ret int32
		return ret
	}
	return *o.Angle.Get()
}

// GetAngleOk returns a tuple with the Angle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextWatermarkOptions) GetAngleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Angle.Get(), o.Angle.IsSet()
}

// HasAngle returns a boolean if a field has been set.
func (o *TextWatermarkOptions) HasAngle() bool {
	if o != nil && o.Angle.IsSet() {
		return true
	}

	return false
}

// SetAngle gets a reference to the given NullableInt32 and assigns it to the Angle field.
func (o *TextWatermarkOptions) SetAngle(v int32) {
	o.Angle.Set(&v)
}
// SetAngleNil sets the value for Angle to be an explicit nil
func (o *TextWatermarkOptions) SetAngleNil() {
	o.Angle.Set(nil)
}

// UnsetAngle ensures that no value is present for Angle, not even an explicit nil
func (o *TextWatermarkOptions) UnsetAngle() {
	o.Angle.Unset()
}

func (o TextWatermarkOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextWatermarkOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.FontName.IsSet() {
		toSerialize["fontName"] = o.FontName.Get()
	}
	if o.FontSize.IsSet() {
		toSerialize["fontSize"] = o.FontSize.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Angle.IsSet() {
		toSerialize["angle"] = o.Angle.Get()
	}
	return toSerialize, nil
}

type NullableTextWatermarkOptions struct {
	value *TextWatermarkOptions
	isSet bool
}

func (v NullableTextWatermarkOptions) Get() *TextWatermarkOptions {
	return v.value
}

func (v *NullableTextWatermarkOptions) Set(val *TextWatermarkOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTextWatermarkOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTextWatermarkOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextWatermarkOptions(val *TextWatermarkOptions) *NullableTextWatermarkOptions {
	return &NullableTextWatermarkOptions{value: val, isSet: true}
}

func (v NullableTextWatermarkOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextWatermarkOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


