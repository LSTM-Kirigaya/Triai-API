package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateTaskResponse struct {
	Success bool           `json:"success"`
	Code    int            `json:"code"`
	Msg     string         `json:"msg"`
	Data    CreateTaskData `json:"data"`
	TraceID string         `json:"trace_id"`
}
type UserInfo struct {
	UserID string `json:"user_id"`
}
type CreateTaskData struct {
	ID                int           `json:"id"`
	AlbumID           string        `json:"album_id"`
	AlbumType         int           `json:"album_type"`
	AlbumScale        string        `json:"album_scale"`
	ExecuteStatus     int           `json:"execute_status"`
	AuditStatus       int           `json:"audit_status"`
	Prompt            string        `json:"prompt"`
	NegPrompt         string        `json:"neg_prompt"`
	ReferenceImageURL string        `json:"reference_image_url"`
	Strength          float64       `json:"strength"`
	UserInfo          UserInfo      `json:"user_info"`
	AigcImages        []interface{} `json:"aigc_images"`
	CreateTimeStamp   int64         `json:"create_time_stamp"`
	UpdateTimeStamp   int64         `json:"update_time_stamp"`
	DeleteStatus      int           `json:"delete_status"`
	IsPositive        bool          `json:"is_positive"`
}

func CreateTask(userToken, prompt, negPrompt, imageKey string) (string, error) {
	var albumId string
	url := "https://www.trikai.com/api/sns/trik/v2/aigc/album/create"
	method := "POST"

	payload := fmt.Sprintf(`{"albumType":1,"source":4,"prompt":"%s","negPrompt":"%s","albumScale":"3:4","imageKey":"%s","strength":0.3}`,
							prompt, negPrompt, imageKey)	
	
	fmt.Println(payload)
	payloadBytes := []byte(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))

	AddCommonHeaders(req, userToken)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		return albumId, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return albumId, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return albumId, err
	}

	
	var response CreateTaskResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return albumId, err
	}

	if response.Success == false {
		fmt.Println("fail to get album id:", response.Msg)
		return albumId, err
	}

	return response.Data.AlbumID, nil
}
