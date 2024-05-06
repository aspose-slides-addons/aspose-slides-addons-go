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
	"fmt"
)

// ExportFormat the model 'ExportFormat'
type ExportFormat string

// List of ExportFormat
const (
	ODP ExportFormat = "Odp"
	OTP ExportFormat = "Otp"
	PPTX ExportFormat = "Pptx"
	PPTM ExportFormat = "Pptm"
	POTX ExportFormat = "Potx"
	PPT ExportFormat = "Ppt"
	PPS ExportFormat = "Pps"
	PPSM ExportFormat = "Ppsm"
	POT ExportFormat = "Pot"
	POTM ExportFormat = "Potm"
	PDF ExportFormat = "Pdf"
	XPS ExportFormat = "Xps"
	PPSX ExportFormat = "Ppsx"
	TIFF ExportFormat = "Tiff"
	HTML ExportFormat = "Html"
	SWF ExportFormat = "Swf"
	TXT ExportFormat = "Txt"
	DOC ExportFormat = "Doc"
	DOCX ExportFormat = "Docx"
	BMP ExportFormat = "Bmp"
	JPEG ExportFormat = "Jpeg"
	PNG ExportFormat = "Png"
	EMF ExportFormat = "Emf"
	WMF ExportFormat = "Wmf"
	GIF ExportFormat = "Gif"
	EXIF ExportFormat = "Exif"
	ICO ExportFormat = "Ico"
	SVG ExportFormat = "Svg"
)

// All allowed values of ExportFormat enum
var AllowedExportFormatEnumValues = []ExportFormat{
	"Odp",
	"Otp",
	"Pptx",
	"Pptm",
	"Potx",
	"Ppt",
	"Pps",
	"Ppsm",
	"Pot",
	"Potm",
	"Pdf",
	"Xps",
	"Ppsx",
	"Tiff",
	"Html",
	"Swf",
	"Txt",
	"Doc",
	"Docx",
	"Bmp",
	"Jpeg",
	"Png",
	"Emf",
	"Wmf",
	"Gif",
	"Exif",
	"Ico",
	"Svg",
}

func (v *ExportFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExportFormat(value)
	for _, existing := range AllowedExportFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExportFormat", value)
}

// NewExportFormatFromValue returns a pointer to a valid ExportFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExportFormatFromValue(v string) (*ExportFormat, error) {
	ev := ExportFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExportFormat: valid values are %v", v, AllowedExportFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExportFormat) IsValid() bool {
	for _, existing := range AllowedExportFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExportFormat value
func (v ExportFormat) Ptr() *ExportFormat {
	return &v
}

type NullableExportFormat struct {
	value *ExportFormat
	isSet bool
}

func (v NullableExportFormat) Get() *ExportFormat {
	return v.value
}

func (v *NullableExportFormat) Set(val *ExportFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableExportFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableExportFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportFormat(val *ExportFormat) *NullableExportFormat {
	return &NullableExportFormat{value: val, isSet: true}
}

func (v NullableExportFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

