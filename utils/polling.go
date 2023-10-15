package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type PollingResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    PollingData `json:"data"`
	TraceID string      `json:"trace_id"`
}

type AigcImages struct {
	OriginalImageURL string `json:"original_image_url"`
	AigcImageURL     string `json:"aigc_image_url"`
	MergeImageURL    string `json:"merge_image_url"`
	FormImageURL     string `json:"form_image_url"`
	ImageIndex       int    `json:"image_index"`
	CollectStatus    bool   `json:"collect_status"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	IsDownload       bool   `json:"is_download"`
	IsLike           bool   `json:"is_like"`
	IsPublish        bool   `json:"is_publish"`
	IsFhd            bool   `json:"is_fhd"`
	IsEvolution      bool   `json:"is_evolution"`
	IsFineTune       bool   `json:"is_fine_tune"`
}
type PollingData struct {
	ID                int          `json:"id"`
	AlbumID           string       `json:"album_id"`
	AlbumType         int          `json:"album_type"`
	AlbumScale        string       `json:"album_scale"`
	ExecuteStatus     int          `json:"execute_status"`
	AuditStatus       int          `json:"audit_status"`
	Prompt            string       `json:"prompt"`
	NegPrompt         string       `json:"neg_prompt"`
	ReferenceImageURL string       `json:"reference_image_url"`
	Strength          float64      `json:"strength"`
	UserInfo          UserInfo     `json:"user_info"`
	AigcImages        []AigcImages `json:"aigc_images"`
	CreateTimeStamp   int64        `json:"create_time_stamp"`
	UpdateTimeStamp   int64        `json:"update_time_stamp"`
	DeleteStatus      int          `json:"delete_status"`
	IsPositive        bool         `json:"is_positive"`
}

func DoOnePolling(userToken, albumId string) ([]AigcImages, error) {
	url := "https://www.trikai.com/api/sns/trik/v2/aigc/album/getUserAlbumDetail?albumId=" + albumId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	AddCommonHeaders(req, userToken)
	req.Header.Add("Referer", " https://www.trikai.com/apps/trikwebapp/create")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var response PollingResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(response.Data.AigcImages)
	return response.Data.AigcImages, nil
}

func downloadImage(link, outputPath string) error {
	response, err := http.Get(link)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func DownloadImages(imageLinks []string, outputDir string) error {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Println("cannot create save dir:", err)
		return err
	}

	for _, link := range imageLinks {
		if len(link) < 8 {
			continue
		}
		imageName := filepath.Base(link)
		outputPath := filepath.Join(outputDir, imageName)

		fmt.Println("Download", link)
		err := downloadImage(link, outputPath)
		if err != nil {
			fmt.Println("Download Fails:", err)
		} else {
			fmt.Println("Save to", outputPath)
		}
	}
	return nil
}

func Polling(userToken, albumId, outputDir string) {
	var count int = 1
	var maxTries int = 10
	for {
		leftTries := maxTries - count
		if leftTries == 0 {
			break
		}
		fmt.Println("do #", count, " polling,", leftTries, "times left")
		count ++

		images, err := DoOnePolling(userToken, albumId)
		if err != nil {
			fmt.Println(err)
			return
		}
		if images == nil {
			fmt.Println("images are nil")
			return
		}
		if len(images) > 0 {
			imageLinks := make([]string, 5)
			for _, aigcImage := range images {
				imageLinks = append(imageLinks, aigcImage.AigcImageURL)
			}
			DownloadImages(imageLinks, outputDir)
			break
		}
		time.Sleep(5 * time.Second)
	}
}
