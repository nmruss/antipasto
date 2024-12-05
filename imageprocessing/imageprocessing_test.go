package imageprocessing

import (
	"testing"
)

func TestReadImage(t *testing.T) {
	pngImagePath := "../testdata/images/cat_1.jpg"
	jpgImagePath := "../testdata/images/cat_2.png"

	pngData := readImage(pngImagePath)
	jpgData := readImage(jpgImagePath)

	if jpgData.width != 1024 {
		t.Fatalf(`readImage('../testdata/images/cat_1.jpg) failed to return the correct width value`)
	}

	if jpgData.height != 1024 {
		t.Fatalf(`readImage('../testdata/image/cat_1.jpg) failed to reuturn the correct height value`)
	}

	if pngData.width != 1024 {
		t.Fatalf(`readImage('../testdata/images/cat_2.png) failed to return the correct width value`)
	}

	if pngData.height != 1024 {
		t.Fatalf(`readImage('../testdata/image/cat_2.png) failed to reuturn the correct height value`)
	}
	// if validFolderStatus == false {
	// 	t.Fatalf(`FolderValid('../testData/valid_banner_folder') failed`)
	// }

	// if invalidFolderStatus == true {
	// 	t.Fatalf(`FolderValid('../testData/invalid_banner_folder') failed`)
	// }
}
