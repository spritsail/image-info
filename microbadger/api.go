package api

import (
	"encoding/json"
	"github.com/microscaling/microbadger/api"
	"io/ioutil"
	"log"
	"net/http"
)

var mbURL string

func Init(url *string) {
	if url != nil {
		mbURL = *url
	} else {
		mbURL = "https://api.microbadger.com/v1"
	}
}

func GetImage(imageName string) (i api.Image, status int, err error) {
	url := mbURL + "/images/" + imageName
	resp, err := http.Get(url)
	status = resp.StatusCode
	if err != nil {
		log.Printf("Failed to build API GET request err %v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("GET request failed %s %d: %s", url, resp.StatusCode, resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read from response to %s", url)
		return
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		log.Printf("Failed to unmarshal response to %s", url)
	}
	return
}

func GetLabels(imageName string) (labels map[string]string, err error) {
	i, status, err := GetImage(imageName)
	if err != nil || status != http.StatusOK {
		return
	}

	labels = i.Labels
	return
}

func GetTag(tagname string, info *api.Image) *api.ImageVersion {
	if info == nil {
		return nil
	}

	for id, tag := range info.Versions {
		for _, tagobj := range tag.Tags {
			if tagobj.Tag == tagname {
				return &info.Versions[id]
			}
		}
	}

	return nil
}
