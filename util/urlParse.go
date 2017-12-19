package util

import (
	"strings"
	"os/exec"
	"os"
	"fmt"
)


func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	s = strings.Replace(s, "\\", "/", -1)
	s = strings.Replace(s, "\\\\", "/", -1)
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	return path
}


func ConcatParamsToUrl(m map[string]string, url string)(string){
	var str,index =  "",0;
	hasParams := strings.Index(url,"?") != -1
	for k, v := range m {
		if hasParams {
			str += "&"+k+"="+v
		}else{
			str += "?"+k+"="+v
			hasParams = true
		}
		index++
	}
	url += str
	return  url
}