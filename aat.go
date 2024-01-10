package aat

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/avvmoto/buf-readerat"
	"github.com/snabb/httpreaderat"
)

// AAT_XML is a URI pointing to AAT XML data stored on the Getty servers.
const AAT_XML string = "http://aatdownloads.getty.edu/VocabData/aat_xml_0622.zip"

// TermsReader returns an `io.ReadCloser` instance containing AAT XML data. If 'uri' is not empty that read will be
// read from 'uri' which is assumed to be a local file on disk. If 'uri' is empty then the data will returned using
// the `FetchTerms` function.
func TermsReader(uri string) (io.ReadCloser, error) {

	if uri == "" {
		return FetchTerms()
	}

	return os.Open(uri)
}

// FetchTerms returns an `io.ReadCloser` instance containing AAT XML data derived from the URI defined in `AAT_XML`.
func FetchTerms() (io.ReadCloser, error) {

	req, err := http.NewRequest("GET", AAT_XML, nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create HTTP request, %w", err)
	}

	http_r, err := httpreaderat.New(nil, req, nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create HTTP reader, %w", err)
	}

	buf_r := bufra.NewBufReaderAt(http_r, 1024*1024)

	zip_r, err := zip.NewReader(buf_r, http_r.Size())

	if err != nil {
		return nil, fmt.Errorf("Failed to create archive reader, %w", err)
	}

	r, err := zip_r.Open("AAT.xml")

	if err != nil {
		return nil, fmt.Errorf("Failed to open AAT reader, %w", err)
	}

	return r, nil
}
