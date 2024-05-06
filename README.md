# Aspose.Slides Addons SDK
The repo contains Aspose.Slides Addons SDK for Go. 

The [Aspose.Slides Addons](https://products.aspose.dev/slides) is a free, simple, and secure RESTful service engaged to process PowerPoint presentations. The SDK is built on top of the RESTful service and allows you to make API calls just with a few lines of code. 

#### Currently supported features:
[Convert presentations](https://products.aspose.dev/slides/convert-api/) | [Merge presentations](https://products.aspose.dev/slides/merge-api/) | [Split presentation](https://products.aspose.dev/slides/split-api/) | [Convert presentation to a video](https://products.aspose.dev/slides/convert-to-video/) | [Protect presentation](https://products.aspose.dev/slides/protect/) | [Unprotect presentation](https://products.aspose.dev/slides/unprotect/) | [Add text watermarks to presentations](https://products.aspose.dev/slides/add-text-watermark/) | [Add image watermarks to presentations](https://products.aspose.dev/slides/add-image-watermark/) | [Replace text in presentations](https://products.aspose.dev/slides/replace-text/) | [Remove macros from presentation](https://products.aspose.dev/slides/remove-macros/) | [Remove annotations from presentation](https://products.aspose.dev/slides/remove-annotations/)

## How to Install
The complete source code is available in this repository folder. Simply integrate provided files into your solution.

## How to Use
You can find usage examples in the **UseCases** folder.

The [Online low-code apps](https://products.aspose.dev/slides) for demonstrating API capabilities are available as well. 

#### Sample usage

The code example below shows how to merge two presentations and save the result in PDF format. One of the presentations will provide an overall style for the output PDF.
```
file0, _ := os.Open("..\\TestData\\test.pptx")
defer file0.Close()
file1, _ := os.Open("..\\TestData\\master.pptx")
defer file1.Close()

options := *slidesclient.NewMergeOptions()
options.SetMasterFileName("master.pptx")
options.SetExcludeMasterFile(false)

request := apiClient.SlidesApi.Merge(slidesclient.PDF).Documents([]*os.File{file0, file1}).Options(options)
resp, httpRes, err := request.Execute()
```

Generatate this code using [Aspose.Slides Merge To PDF Low-Code app](https://products.aspose.dev/slides/merge-api/merge-to-pdf/)

## Resources
[Aspose.Slides Addons Low-Code Applications](https://products.aspose.dev/slides)

## Contact Us
Your feedback is very important to us. Please feel free to contact us using our [Support Forums]("https://forum.aspose.cloud/c/slides").

## FAQ
### What the Difference Between Aspose.Slides for Cloud and Aspose.Slides Addons?
The Aspose.Slides for Cloud is a complex set of tools which covers almost every aspect of manipulating PowerPoint presentation.

Meanwhile the Aspose.Slides Addons are much simpler, easy to use, and free. The addons support the most requested features and allow you to utilize them most simply.