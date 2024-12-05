package imageprocessing

import (
	"testing"
)

func TestReadImage(t *testing.T) {
	//Tests if ReadImage() function returns valid data on JPG's and PNG's
	jpgImagePath := "../testdata/images/cat_1.jpg"
	solidPNGImagePath := "../testdata/images/cat_2.png"
	transparentPNGImagePath := "../testdata/images/cat_3.png"

	solidPNGData := readImage(solidPNGImagePath)
	jpgData := readImage(jpgImagePath)
	transparentPNGData := readImage(transparentPNGImagePath)

	if jpgData.width != 1024 {
		t.Fatalf(`readImage('../testdata/images/cat_1.jpg) failed to return the correct width value`)
	}

	if jpgData.height != 1024 {
		t.Fatalf(`readImage('../testdata/image/cat_1.jpg) failed to reuturn the correct height value`)
	}

	if solidPNGData.height != 1024 {
		t.Fatalf(`readImage('../testdata/images/cat_2.png) failed to return the correct width value`)
	}

	if solidPNGData.height != 1024 {
		t.Fatalf(`readImage('../testdata/image/cat_2.png) failed to reuturn the correct height value`)
	}

	if transparentPNGData.topX != 400 {
		t.Fatalf(`readImage('../testdata/image/cat_3.png) failed to reuturn the correct topX value`)
	}

	if transparentPNGData.topY != 300 {
		t.Fatalf(`readImage('../testdata/image/cat_3.png) failed to reuturn the correct topY value`)
	}
}
