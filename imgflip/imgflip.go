package imgflip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ImgflipResponseData struct {
	url     string `json:"url"`
	pageUrl string `json:"page_url"`
}

type ImgflipResponse struct {
	success      bool                `json:"success"`
	data         ImgflipResponseData `json:"data"`
	errorMessage string              `json:"error_message"`
}

const imgflipBaseURL string = "https://api.imgflip.com"

func CreateMeme(templateId string, text0 string, text1 string, username string, password string) {
	// Welp, imgflip makes you send your password in the url params.
	// ABSOLUTELY_BARBARIC.jpg
	imgflipUrl := fmt.Sprintf("%s/caption_image?template_id=%s&text0=%s&text1=%s&username=%s&password=%s",
		imgflipBaseURL,
		url.QueryEscape(templateId),
		url.QueryEscape(text0),
		url.QueryEscape(text1),
		url.QueryEscape(username),
		url.QueryEscape(password))
	fmt.Println(imgflipUrl)

	response, err := http.Get(imgflipUrl)

	if err != nil {
		fmt.Println("API call Blew UP!")
		panic(err)
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Reading the response borkd!")
		panic(err)
	}

	fmt.Println(string(responseData))
	var responseObject ImgflipResponse
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.success)
	fmt.Println(responseObject.data.pageUrl)
	fmt.Println(responseObject.errorMessage)

}
