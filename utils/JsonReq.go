package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//json请求方式
func JsonPost(url, raw string) string {
	jsonStr := []byte(raw)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "-1"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
