package main

import (
	"call-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Demo call an API")
	const url string = "http://pokeapi.co/api/v2/pokedex/kanto/"

	response, err := http.Get(url)
	if err != nil {
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var responseObj models.Response
	err = json.Unmarshal(responseData, &responseObj)
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, v := range responseObj.Pokemon {
		v.Info()
	}
}
