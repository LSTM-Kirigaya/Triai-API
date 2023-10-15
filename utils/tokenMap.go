package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetImageTokenMapResponse struct {
	Success bool                 `json:"success"`
	Code    int                  `json:"code"`
	Msg     string               `json:"msg"`
	Data    map[string]string    `json:"data"`
	TraceID string               `json:"trace_id"`
}

func GetImageTokenMap(userToken, fileId string) (string, error) {
	var uploadedImageUrl string
	url := "https://www.trikai.com/api/sns/trik/v2/upload/batchGetFileKeyMap?fileKey=" + fileId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return uploadedImageUrl, err
	}

	AddCommonHeaders(req, userToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return uploadedImageUrl, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return uploadedImageUrl, err
	}
	
	var response GetImageTokenMapResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return uploadedImageUrl, err
	}

	for _, realUrl := range response.Data {
		uploadedImageUrl = realUrl
		break
	}

	return "https://cdn.shanguangshipin.com/" + uploadedImageUrl, nil
}
