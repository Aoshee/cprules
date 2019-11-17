package mantf

import (
	"compress/flate"
	"fmt"

	"github.com/mholt/archiver"
)

func zipF(fileName string) {
	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
	if fileExists(fileName) {
		fmt.Println("Exist Files")
	} else {
		err := z.Archive([]string{fileName}, "/tmp/test.zip")
		if err != nil {
			fmt.Println("Success")
		}
	}
}
