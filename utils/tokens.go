package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"errors"
)

type GetUploadTokenResponse struct {
	Success bool               `json:"success"`
	Code    int                `json:"code"`
	Msg     string             `json:"msg"`
	Data    GetUploadTokenData `json:"data"`
	TraceID string             `json:"trace_id"`
}

type UploadTempPermits struct {
	Token         string   `json:"token"`
	UploadAddr    string   `json:"uploadAddr"`
	ExpireTime    int64    `json:"expireTime"`
	Qos           float64  `json:"qos"`
	CloudType     int      `json:"cloudType"`
	Bucket        string   `json:"bucket"`
	Region        string   `json:"region"`
	FileIds       []string `json:"fileIds"`
	MasterCloudID int      `json:"masterCloudId"`
	UploadID      int      `json:"uploadId"`
	SecretID      string   `json:"secretId,omitempty"`
	SecretKey     string   `json:"secretKey,omitempty"`
	CdnDomain     string   `json:"cdnDomain,omitempty"`
}

type GetUploadTokenData struct {
	UploadTempPermits []UploadTempPermits `json:"uploadTempPermits"`
}

func GetUploadToken(userToken string) (string, string, error) {
	var uploadToken, fileId string

	url := "https://www.trikai.com/api/sns/trik/v2/upload/token?bizName=vertical&scene=image&fileCount=1&biz_name=vertical&version=1&file_count=1&subsystem=web_resource&source=web"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return uploadToken, fileId, err
	}

	AddCommonHeaders(req, userToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return uploadToken, fileId, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return uploadToken, fileId, err
	}
	var response GetUploadTokenResponse
	err = json.Unmarshal(body, &response)
	if len(response.Data.UploadTempPermits) == 0 {
		err = errors.New("uploadTempPermits array is empty")
		return uploadToken, fileId, err
	}
	uploadTempPermit := response.Data.UploadTempPermits[0]
	uploadToken = uploadTempPermit.Token
	fileId = uploadTempPermit.FileIds[0]
	return uploadToken, fileId, nil
}
