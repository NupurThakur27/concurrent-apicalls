package apiCall

import (
	"Test/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func (a *API) GetDataFromGoRest(pageNumber int) []model.User {
	res, err := http.Get(a.BaseURL + "/public/v1/users?page=" + strconv.Itoa(pageNumber))
	if err != nil {
		fmt.Println("unable to make the request", err)
	}
	// fmt.Println(res)
	if res.StatusCode != 200 {
		log.Println("unable to get the data", res.StatusCode)
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var data model.Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("uanble to parse json", err)
		return nil
	}

	return data.Users
}
