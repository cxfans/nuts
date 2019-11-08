/*
	Implements Baidu OCR API
*/

package nuts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// The high-precision version can be requested 500 times a day for free
type ApiOcr struct {
	apiKey    string
	secretKey string
	apiUrl    string
	token     string
}

func NewClient(apiKey, secretKey, apiUrl string) ApiOcr {
	return ApiOcr{apiKey, secretKey, apiUrl, ""}
}

func (api *ApiOcr) GetToken() (string, error) {
	authUrl := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?"+
		"grant_type=client_credentials&client_id=%s&client_secret=%s&", api.apiKey, api.secretKey)
	resp, err := http.Get(authUrl)
	if err != nil {
		return "", fmt.Errorf("get access token failed: [%w]", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read html body failed: [%w]", err)
	}
	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("can't unmarshal data:[%v], error: [%w]", body, err)
	}
	token := data["access_token"]
	if token == nil {
		return "", fmt.Errorf("no access token:[%v]", data)
	}
	if token, ok := token.(string); ok != false {
		api.token = token
		return token, nil
	} else {
		return "", fmt.Errorf("error type: %v", data)
	}
}

func (api ApiOcr) GetWords(base64Str string) (wordsList []string, err error) {
	if api.token == "" {
		if _, err = api.GetToken(); err != nil {
			return nil, err
		}
	}
	c := &http.Client{}
	param := url.Values{}
	param.Add("image", base64Str)
	r, _ := http.NewRequest("POST", api.apiUrl+"?access_token="+api.token,
		strings.NewReader(param.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")
	resp, err := c.Do(r)
	if err != nil {
		return nil, fmt.Errorf("request failed: [%w]", err)
	}
	var data map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read html body failed: [%w]", err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal data failed: [%w]", err)
	}
	ws := data["words_result"]
	if ws == nil {
		return nil, fmt.Errorf("no characters are recognized, "+
			"or the number of free requests exceeded the limit\nresponse: %v", data)
	}
	if lst, ok := ws.([]interface{}); ok != false {
		for _, w := range lst {
			if mp, ok := w.(map[string]interface{}); ok != false {
				if str, ok := mp["words"].(string); ok != false {
					wordsList = append(wordsList, strings.TrimSpace(str))
				}
			}
		}
	} else {
		return nil, fmt.Errorf("failed, response: [%v]", data)
	}
	return wordsList, nil
}

// form the image path directly
func (api ApiOcr) GetWordsFromImage(path string) (wordsList []string, err error) {
	if base64Str, err := ImageExportBase64(path, false); err == nil {
		return api.GetWords(base64Str)
	} else {
		return nil, err
	}
}
