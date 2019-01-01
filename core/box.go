package core

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/devectron/sunlight/log"
)

const (
	UPLOAD    = "https://content.dropboxapi.com/2/files/upload"
	TMPUPLOAD = "https://api.dropboxapi.com/2/files/get_temporary_upload_link"
	SHARE     = ""
	REMOVE    = ""
	PATH      = "/Apps/devectron.sunlight/"
)

func Upload(name string, file io.Reader, token string) string {
	log.Inf("Uploading %s...", name)
	client := &http.Client{}
	req, err := http.NewRequest("POST", UPLOAD, file)
	if err != nil {
		log.Err("Error while requesting ...\n %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Dropbox-API-Arg", "{\"path\": \""+PATH+name+"\",\"mode\": \"add\",\"autorename\": true,\"mute\": false,\"strict_conflict\": false}")
	req.Header.Add("Content-Type", "application/octet-stream")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return GetTmpLink(name, token)
}
func GetTmpLink(name string, token string) string {
	b := bytes.NewReader([]byte("{\"commit_info\": {\"path\": \"" + PATH + name + "\",\"mode\": \"add\",\"autorename\": true,\"mute\": false,\"strict_conflict\": false},\"duration\": 300}"))
	req, err := http.NewRequest("POST", TMPUPLOAD, b)
	if err != nil {
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Err("Error while responding ...\n%v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err("Error while reading from body...\n%v", err)
	}
	var data map[string]interface{}
	er := json.Unmarshal([]byte(body), &data)
	if er != nil {
		log.Err("Error while Unmarshing json...\n%v", er)
	}
	return data["link"].(string)
}
func Share() string {
	return ""
}
func Delete() {

}
