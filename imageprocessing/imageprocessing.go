package imageprocessing

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"regexp"
)

type ImageData struct {
	imageType  string
	width      int
	height     int
	topX       int
	topY       int
	colorModel color.Model
}

func readImage(imagePath string) ImageData {
	var imgData ImageData
	//reads an image and returns an ImageData containing information about it
	isJPG, err := regexp.Match(`.jpg`, []byte(imagePath))
	if err != nil {
		log.Println(err)
	}

	isPNG, err := regexp.Match(`.png`, []byte(imagePath))
	if err != nil {
		log.Println(err)
	}

	file, err := os.Open(imagePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
	}

	bounds := img.Bounds()
	imgData.width = bounds.Max.X
	imgData.height = bounds.Max.Y
	imgData.colorModel = img.ColorModel()

	if isJPG {
		imgData.imageType = "jpg"
	}

	if isPNG {
		imgData.imageType = "png"
	}

	//calculate top left non-transparent pixel for png positioning
	//NOTE: look for efficiency here?
	if isPNG {
		coordinates := []int{0, 0}
		for coordinates[0] < bounds.Max.X {
			coordinates[1] = 0
			for coordinates[1] < bounds.Max.Y {
				if checkAlpha(coordinates, &img) {
					break
				}
				coordinates[1]++
			}

			isAlpha := checkAlpha(coordinates, &img)
			if isAlpha {
				imgData.topX = coordinates[0]
				imgData.topY = coordinates[1]
				break
			}

			coordinates[0]++
		}

	}

	log.Println("image data:", imgData)
	return imgData
}

func checkAlpha(coordinates []int, img *image.Image) bool {
	//check if the current coordinate has an alpha value > 0
	//if so, set the image data and return true
	image := *img
	_, _, _, a := image.At(coordinates[0], coordinates[1]).RGBA()
	alpha := uint8(a & 0xff)
	if alpha != 0 {
		return true
	} else {
		return false
	}
}
