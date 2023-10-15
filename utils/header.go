package utils

import "net/http"

func AddCommonHeaders(req *http.Request, userToken string) {

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.Add("Sec-Ch-Ua", "Microsoft Edge")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", " Mozilla/5.0 AppleWebKit/537.36 like Gecko) Chrome/119.0 Safari/537.36 Edg/119.0 Safari/537.36) Edg/119.0 Safari/537.36 Edg/119.0 Safari/537.36 Edg/119.0 Safari/537.36 Edg/119.0 Safari/537.36 Edg/119.0 Safari/537.36 Edg/119.0 Safari/537.36 Edg/119.0 Safari/537.36 Edg/119 Safari")
	req.Header.Add("Xy-Platform-Info", "platform=pc&deviceId=ae71a23e80a608bbe77f32f6137ed012")
	req.Header.Add("Connection", "keep-alive")
	if len(userToken) > 0 {
		req.Header.Add("Host", "www.trikai.com")
		req.Header.Add("Cookie", "trikwebapp-status=online trikwebapp-status.sig=wKrvykY-aT73ktBR2utL7zzeeyAwK_GaL2xW3Y8zmUk")
		req.Header.Add("V-User-Token", userToken)
	}
}