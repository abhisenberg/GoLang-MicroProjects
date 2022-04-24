package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
Convert the incoming data from requests into objects ("unmarshalling of json").
Params: HTTP request r, an empty interface x


**/
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { //Reading the body of the request
		if err := json.Unmarshal([]byte(body), x); err != nil { //Unmarshalls the contents of body and stores the result in x
			return
		}
	}
}
