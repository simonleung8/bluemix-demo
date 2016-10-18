package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

const HOST_URL = "http://uploads.im/api?upload"

type response struct {
	Data data `json:"data"`
}

type data struct {
	ImgUrl string `json:"img_url"`
}

func HostImage(file multipart.File, header *multipart.FileHeader) (imageUrl string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	defer file.Close()
	fw, err := w.CreateFormFile("upload", header.Filename)
	Must(err, "Error creating multipart file")

	_, err = io.Copy(fw, file)
	Must(err, "Error writing to multipart file")

	// close for request to have the terminating boundary.
	w.Close()

	req, err := http.NewRequest("POST", HOST_URL, &b)
	Must(err, "Error creating request for image host")

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	Must(err, "Error posting request to image host")

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	Must(err, "Error reading response body from image host")

	imgData := response{}
	Must(json.Unmarshal(body, &imgData), "Error unmarshaling response body from image host")

	return imgData.Data.ImgUrl
}
