package util

import (
	"encoding/json"
)




func ToJSON(i interface{})(str string,err error){
	b, err := json.Marshal(i)
	return  string(b),err
}

func ParseJSON(str string, i interface{})(err error){
	err =json.Unmarshal([]byte(string(str)),i)
	return err;
}




