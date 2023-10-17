package handler

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
)

func createRandomZipContent() ([]byte, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Generate random content and add it to the zip archive
	for i := 0; i < 5; i++ {
		fileName := fmt.Sprintf("file%d.txt", i)
		content := make([]byte, 100) // Generate random content
		rand.Read(content)
		fileWriter, err := zipWriter.Create(fileName)
		if err != nil {
			return nil, err
		}
		_, err = fileWriter.Write(content)
		if err != nil {
			return nil, err
		}
	}

	zipWriter.Close()
	return buf.Bytes(), nil
}

func TestCreateArchiveHandler(t *testing.T) {

	t.Run("Successfully create archive", func(t *testing.T) {
		randomZipContent, err := createRandomZipContent()
		if err != nil {
			t.Fatalf("Failed to create random zip content: %v", err)
		}

		var b bytes.Buffer
		writer := multipart.NewWriter(&b)

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", `form-data; name="files[]"; filename="random.zip"`)
		h.Set("Content-Type", "application/zip")

		part, err := writer.CreatePart(h)
		if err != nil {
			t.Fatalf("Failed to create multipart part: %v", err)
		}

		_, err = io.Copy(part, bytes.NewReader(randomZipContent))
		if err != nil {
			t.Fatalf("Failed to copy random zip content: %v", err)
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "http://example.com/archive", &b)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Header().Get("Content-Type") != "application/zip" {
			t.Errorf("handler returned wrong content type: got %v want %v", rr.Header().Get("Content-Type"), "application/zip")
		}
	})
}
