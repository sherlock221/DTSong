package util

import (
	"strings"
)



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