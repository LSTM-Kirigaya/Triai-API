package main

import (
	// "bytes"
	// "encoding/json"
	"flag"
	"fmt"
	"trikai/utils"
	// "net/http"
)


func main() {
	userToken := flag.String("userToken", "", "userToken")
	imagePath := flag.String("image", "", "path of image")
	prompt := flag.String("prompt", "", "prompt")
	negPrompt := flag.String("negPrompt", "", "negative prompt")
	outputDir := flag.String("outputDir", "", "output dictionary")
	
	flag.Parse()

	uploadToken, fileId, err := utils.GetUploadToken(*userToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = utils.UploadImage(uploadToken, fileId, *imagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	uploadedImageUrl, err := utils.GetImageTokenMap(*userToken, fileId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("upload image to: ", uploadedImageUrl)
	albumId, err := utils.CreateTask(*userToken, *prompt, *negPrompt, fileId)
	if len(albumId) == 0 {
		fmt.Println("albumId is empty !")
		return
	}
	fmt.Println("albumId", albumId)
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.Polling(*userToken, albumId, *outputDir)
}