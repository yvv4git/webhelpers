package webhelpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fileUploadParamName = "file"
	fileUploadName      = "/data/symbol_yvv.png"
	dirToSave           = "data"
)

type fileUploadHandler struct{}

func (h *fileUploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set max upload size 10 MB.
	r.ParseMultipartForm(10 << 20)

	// File from http post body.
	file, handler, err := r.FormFile(fileUploadParamName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Open temporary file.
	/* tempFile, err := ioutil.TempFile(dirToSave, "upload-*.png")
	if err != nil {
		log.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Println(err)
		return
	} */

	//log.Println("File size: ", handler.Size)
	fmt.Fprintf(w, fmt.Sprintf("%v", handler.Size))
}

func TestFileUpload(t *testing.T) {
	t.Log("Test - File upload")

	// Server.
	testHandler := &fileUploadHandler{}
	server := httptest.NewServer(testHandler)
	defer server.Close()

	// Client.
	pathToFile, _ := os.Getwd()
	pathToFile += fileUploadName

	addtitionalParams := map[string]string{
		"author": "Vladimir Elieseev",
	}

	req, err := RequestUploadFile(
		fmt.Sprintf("%s/file", server.URL),
		addtitionalParams,
		fileUploadParamName,
		pathToFile,
	)
	if err != nil {
		panic(err)
	}

	// Create http client and send request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read body from buffer.
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "17194", string(bodyData))
}
