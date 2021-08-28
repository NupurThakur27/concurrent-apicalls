package main

import (
	"Test/apiCall"
	"Test/model"
	"log"
	"sync"
)

const URL = "https://gorest.co.in/"

func main() {
	log.Println("Start...")
	var users []model.User

	pageNumbers := []int{1, 2, 3}
	wg := sync.WaitGroup{}
	for _, id := range pageNumbers {
		wg.Add(1)
		go func(id int) {
			a := apiCall.API{BaseURL: URL}
			u := a.GetDataFromGoRest(id, &wg)
			users = append(users, u...)
		}(id)
	}
	wg.Wait()
	log.Println(users)

}
