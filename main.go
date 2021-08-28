package main

import (
	"Test/apiCall"
	"Test/model"
	"fmt"
	"net/http"
	"sync"
)

const URL = "https://gorest.co.in/"

func main() {
	fmt.Println("Start...")
	client := &http.Client{}
	var users []model.User
	a := apiCall.API{client, URL}

	pageNumbers := []int{1, 2, 3}
	wg := sync.WaitGroup{}
	for id := range pageNumbers {
		wg.Add(1)
		go func(id int) {
			u := a.GetDataFromGoRest(id)
			users = append(users, u...)

			wg.Done()
		}(id)
	}
	wg.Wait()
	fmt.Println(users)

}
