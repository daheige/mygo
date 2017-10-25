package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{} //bytes buffer
	bodyWriter := multipart.NewWriter(bodyBuf)

	//create form file
	fileWriter, err := bodyWriter.CreateFormFile("upload_file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//open file fd
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	defer fp.Close()

	//copy
	_, err = io.Copy(fileWriter, fp)

	if err != nil {
		return err
	}

	//file type
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	res, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res.Status)
	fmt.Println(string(res_body))
	return nil
}
