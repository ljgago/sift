package tgz

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func ReadFile(tgzReader io.Reader, filename string) ([]byte, error) {
	// Create a gzip reader
	gzReader, err := gzip.NewReader(tgzReader)
	if err != nil {
		panic(err)
	}
	defer gzReader.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return []byte{}, err
		}

		if strings.Contains(header.Name, filename) {
			return io.ReadAll(tarReader)
		}
	}

	return []byte{}, fmt.Errorf("file %s not found in tar", filename)
}
