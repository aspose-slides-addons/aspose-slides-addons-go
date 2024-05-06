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

// checks if the ImageWatermarkOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageWatermarkOptions{}

// ImageWatermarkOptions struct for ImageWatermarkOptions
type ImageWatermarkOptions struct {
	Angle NullableInt32 `json:"angle,omitempty"`
	Zoom NullableInt32 `json:"zoom,omitempty"`
}

// NewImageWatermarkOptions instantiates a new ImageWatermarkOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageWatermarkOptions() *ImageWatermarkOptions {
	this := ImageWatermarkOptions{}
	return &this
}

// NewImageWatermarkOptionsWithDefaults instantiates a new ImageWatermarkOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWatermarkOptionsWithDefaults() *ImageWatermarkOptions {
	this := ImageWatermarkOptions{}
	return &this
}

// GetAngle returns the Angle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageWatermarkOptions) GetAngle() int32 {
	if o == nil || IsNil(o.Angle.Get()) {
		var ret int32
		return ret
	}
	return *o.Angle.Get()
}

// GetAngleOk returns a tuple with the Angle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageWatermarkOptions) GetAngleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Angle.Get(), o.Angle.IsSet()
}

// HasAngle returns a boolean if a field has been set.
func (o *ImageWatermarkOptions) HasAngle() bool {
	if o != nil && o.Angle.IsSet() {
		return true
	}

	return false
}

// SetAngle gets a reference to the given NullableInt32 and assigns it to the Angle field.
func (o *ImageWatermarkOptions) SetAngle(v int32) {
	o.Angle.Set(&v)
}
// SetAngleNil sets the value for Angle to be an explicit nil
func (o *ImageWatermarkOptions) SetAngleNil() {
	o.Angle.Set(nil)
}

// UnsetAngle ensures that no value is present for Angle, not even an explicit nil
func (o *ImageWatermarkOptions) UnsetAngle() {
	o.Angle.Unset()
}

// GetZoom returns the Zoom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageWatermarkOptions) GetZoom() int32 {
	if o == nil || IsNil(o.Zoom.Get()) {
		var ret int32
		return ret
	}
	return *o.Zoom.Get()
}

// GetZoomOk returns a tuple with the Zoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageWatermarkOptions) GetZoomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zoom.Get(), o.Zoom.IsSet()
}

// HasZoom returns a boolean if a field has been set.
func (o *ImageWatermarkOptions) HasZoom() bool {
	if o != nil && o.Zoom.IsSet() {
		return true
	}

	return false
}

// SetZoom gets a reference to the given NullableInt32 and assigns it to the Zoom field.
func (o *ImageWatermarkOptions) SetZoom(v int32) {
	o.Zoom.Set(&v)
}
// SetZoomNil sets the value for Zoom to be an explicit nil
func (o *ImageWatermarkOptions) SetZoomNil() {
	o.Zoom.Set(nil)
}

// UnsetZoom ensures that no value is present for Zoom, not even an explicit nil
func (o *ImageWatermarkOptions) UnsetZoom() {
	o.Zoom.Unset()
}

func (o ImageWatermarkOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageWatermarkOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Angle.IsSet() {
		toSerialize["angle"] = o.Angle.Get()
	}
	if o.Zoom.IsSet() {
		toSerialize["zoom"] = o.Zoom.Get()
	}
	return toSerialize, nil
}

type NullableImageWatermarkOptions struct {
	value *ImageWatermarkOptions
	isSet bool
}

func (v NullableImageWatermarkOptions) Get() *ImageWatermarkOptions {
	return v.value
}

func (v *NullableImageWatermarkOptions) Set(val *ImageWatermarkOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableImageWatermarkOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableImageWatermarkOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageWatermarkOptions(val *ImageWatermarkOptions) *NullableImageWatermarkOptions {
	return &NullableImageWatermarkOptions{value: val, isSet: true}
}

func (v NullableImageWatermarkOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageWatermarkOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


