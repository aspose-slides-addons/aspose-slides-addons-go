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
	"os"
	"testing"

	slidesclient "github.com/aspose-slides-addons/aspose-slides-addons-go/asposeslidesaddons"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTextWatermark(t *testing.T) {
	apiClient := slidesclient.NewAPIClient()

	t.Run("Test Text Watermark", func(t *testing.T) {

		file, _ := os.Open("..\\TestData\\test.pptx")
		defer file.Close()

		options := *slidesclient.NewTextWatermarkOptions()
		options.SetText("Watermark text")
		options.SetColor("#FF0000")
		options.SetFontName("Arial")
		options.SetFontSize(32)
		options.SetAngle(-45)
		request := apiClient.SlidesApi.TextWatermark().Documents([]*os.File{file}).Options(options)
		resp, httpRes, err := request.Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
