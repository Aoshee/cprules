package mantf

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func rulesDownload(uri string) ([]byte, error) {
	fmt.Printf("Downloading...: %s\n", uri)
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read File: Size of download: %d\n", len(d))
	return d, err
}

func WriteFile(dst string, d []byte) error {
	fmt.Printf("Write File: Size of download: %d\n", len(d))
	err := ioutil.WriteFile(dst, d, 0444)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DownloadTF(uri string, dst string) {
	fmt.Printf("Download From: %s\n", uri)
	if d, err := rulesDownload(uri); err == nil {
		fmt.Printf("downloaded %s.\n", uri)
		if WriteFile(dst, d) == nil {
			fmt.Printf("Saved %s as %s\n", uri, dst)
		}
	}
}

func ExtractCompress(gzipStream io.Reader) {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("Extract Compressed File Failed")
	}
	uncomReader := tar.NewReader(uncompressedStream)
	for true {
		header, err := uncomReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Extract Compressed Next() Failed: %s", err.Error())
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0775); err != nil {
				log.Fatal("Extract Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				log.Fatal("Extract Create() failed: %s", err.Error())
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, uncomReader); err != nil {
				log.Fatalf("Extract Copy() failed: %s", err.Error())
			}
		default:
			log.Fatal("Extract TarGz: unknown type: %s in %s", header.Typeflag, header.Name)
		}
	}
}
