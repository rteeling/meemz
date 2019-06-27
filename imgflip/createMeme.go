package imgflip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/logrusorgru/aurora"
)

type ImgflipResponseData struct {
	Url     string `json:"url"`
	PageUrl string `json:"page_url"`
}

type ImgflipResponse struct {
	Success      bool                `json:"success"`
	Data         ImgflipResponseData `json:"data"`
	ErrorMessage string              `json:"error_message"`
}

const imgflipBaseURL string = "https://api.imgflip.com"

// Calls imgflip api and returns link to meme
func CreateMeme(templateId string, text0 string, text1 string, username string, password string) string {
	// Welp, imgflip makes you send your password in the url params.
	// ABSOLUTELY_BARBARIC.jpg
	imgflipUrl := fmt.Sprintf("%s/caption_image?template_id=%s&text0=%s&text1=%s&username=%s&password=%s",
		imgflipBaseURL,
		url.QueryEscape(templateId),
		url.QueryEscape(text0),
		url.QueryEscape(text1),
		url.QueryEscape(username),
		url.QueryEscape(password))

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

	var responseObject ImgflipResponse
	json.Unmarshal(responseData, &responseObject)

	if !responseObject.Success {
		fmt.Println(aurora.Sprintf("[%s]      %s", aurora.Red("ERROR"), responseObject.ErrorMessage))

		os.Exit(1)
	}

	return responseObject.Data.Url
}
