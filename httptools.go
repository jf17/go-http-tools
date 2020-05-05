package httptools

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func GetBody(url string) string {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	return string(out)
}

