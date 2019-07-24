/*
https://rapidapi.com/matchilling/api/chuck-norris
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://matchilling-chuck-norris-jokes-v1.p.rapidapi.com/jokes/random", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Set("X-RapidAPI-Host", "matchilling-chuck-norris-jokes-v1.p.rapidapi.com")
	req.Header.Set("X-RapidAPI-Key", "84870875eamsh9436a31c7fc27c9p1eb664jsn1dca8e6f043d")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(body))
	})
	http.ListenAndServe(":8080", nil)
}
