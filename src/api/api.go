package api

import (
	"net/http"
	"encoding/json"
	"bytes"
	"log"
)

type Api struct {
    User string
    Password string
    AppKey string
    BaseUrl string
	IsVerbose bool
}

type rowWithResultString struct {
    Result   string      `json:"result"`
}

type versionResponse struct {
    Rows   []rowWithResultString      `json:"rows"`
}

func extractStringResult(resp *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	res := versionResponse {}
	json.Unmarshal(buf.Bytes(), &res)
	return res.Rows[0].Result
}

func (this *Api) postForString(url string) string {
	client := &http.Client {}
	req, err := http.NewRequest("POST", this.BaseUrl + url, nil)
	if err == nil {
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		if this.AppKey != "" {
			if this.IsVerbose {
				log.Print("DEBUG: Authenticating with AppKey")
			}
			req.Header.Add("AppKey", this.AppKey)
		} else if this.User != "" && this.Password != "" {
			if this.IsVerbose {
				log.Print("DEBUG: Authenticating with HTTP Basic Auth")
			}
			req.SetBasicAuth(this.User, this.Password)
		} else {
			log.Fatal("ERROR: Don't know how to authenticate")
		}
		resp, err := client.Do(req)
		if err == nil {
			return extractStringResult(resp)
		} else {
			log.Fatal(err)
			return ""
		}
	} else {
		log.Fatal(err)
		return ""
	}
}

func (this *Api) GetVersion() string {
	return this.postForString("/Subsystems/PlatformSubsystem/Services/GetThingworxVersion")
}

