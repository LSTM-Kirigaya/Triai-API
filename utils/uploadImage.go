package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func UploadImage(uploadToken, fileId, imagePath string) error {
	url := "https://ros-upload.xiaohongshu.com/" + fileId
	method := "PUT"

	imageFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer imageFile.Close()
	
	client := &http.Client{}
	req, err := http.NewRequest(method, url, imageFile)

	if err != nil {
		fmt.Println(err)
		return err
	}

	AddCommonHeaders(req, "")
	req.Header.Add("X-Cos-Security-Token", uploadToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		return nil
	} else {
		return errors.New("status code is not 200")
	}
}
