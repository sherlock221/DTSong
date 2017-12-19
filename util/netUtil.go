package util

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"io"
)

func GetUrl(url string, params map[string]string)(res string,err error){
	if(params != nil){
		url = ConcatParamsToUrl(params,url)
	}
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body),err
}


func CurlToFile(url string,filePath string)(err error){
	file, errFile := os.Create(filePath)
	if errFile != nil {
		fmt.Println(errFile)
	}
	res, err := http.Get(url)
	if err != nil {
		return  err
	}
	defer res.Body.Close()
	io.Copy(file, res.Body)
	return  nil

}