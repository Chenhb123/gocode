package main

import ( // "log"
	// "fmt"
	// "os"

	// "github.com/pdfcrowd/pdfcrowd-go"
	"fmt"
	"image"
	"log"

	// "log"
	"os"

	"github.com/signintech/gopdf"
)

// test png to pdf

func main() {
	width, height, err := getImgInfo("godocker.png")
	if err != nil {
		log.Fatal(err)
	}
	// var width, height float64
	// width, height = 1920, 1080
	fmt.Println(width, height)
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: width, H: height}})
	// pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	// var err error
	// err = pdf.AddTTFFont("loma", "../ttf/Loma.ttf")
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }

	pdf.Image("godocker.png", 0, 0, nil) //print image
	// err = pdf.SetFont("loma", "", 14)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }
	// pdf.SetX(250) //move current location
	// pdf.SetY(200)
	// pdf.Cell(nil, "gopher and gopher") //print text

	pdf.WritePdf("godocker.pdf")
	// create the API client instance
	// client := pdfcrowd.NewImageToPdfClient("pandabi2020pdf", "b8aa44b8d3b22af02dfc01ad741ce098")

	// // run the conversion and write the result to a file
	// err := client.ConvertFileToFile("logo.png", "logo.pdf")

	// // check for a conversion error
	// handleError(err)
}

func getImgInfo(filepath string) (float64, float64, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		// fmt.Println("err = ", err)
		return 0, 0, err
	}
	imageRect := img.Bounds()
	// file.Close()
	k := 1
	w := -128 //init
	h := -128 //init
	if w < 0 {
		w = -imageRect.Dx() * 72 / w / k
	}
	if h < 0 {
		h = -imageRect.Dy() * 72 / h / k
	}
	if w == 0 {
		w = h * imageRect.Dx() / imageRect.Dy()
	}
	if h == 0 {
		h = w * imageRect.Dy() / imageRect.Dx()
	}
	return float64(w), float64(h), nil
}

// func handleError(err error) {
// 	if err != nil {
// 		// report the error
// 		why, ok := err.(pdfcrowd.Error)
// 		if ok {
// 			os.Stderr.WriteString(fmt.Sprintf("Pdfcrowd Error: %s\n", why))
// 		} else {
// 			os.Stderr.WriteString(fmt.Sprintf("Generic Error: %s\n", err))
// 		}

// 		// rethrow or handle the exception
// 		panic(err.Error())
// 	}
// }
