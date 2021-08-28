package apiCall

import (
	"Test/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func (a *API) GetDataFromGoRest(pageNumber int, wg *sync.WaitGroup) []model.User {
	res, err := http.Get(a.BaseURL + "/public/v1/users?page=" + strconv.Itoa(pageNumber))
	defer wg.Done()
	if err != nil {
		log.Println("unable to make the request", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	if res.StatusCode != 200 {
		log.Println(string(body))
		log.Println("unable to get the data", res.StatusCode)
		return nil
	}

	var data model.Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("uanble to parse json", err)
		return nil
	}

	return data.Users
}
