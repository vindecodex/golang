package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Movie struct {
		Title   string
		Year    int
		Ratings int
	}

	// instanciate slices of Movie
	movies := []Movie{
		{
			Title:   "Batman",
			Year:    2019,
			Ratings: 8,
		},
		{
			Title:   "Superman",
			Year:    2020,
			Ratings: 9,
		},
	}

	// Converting to byte of slice
	converted, err := json.Marshal(movies)

	if err != nil {
		fmt.Println("Error")
	}

	// Converting byte to string so that it is readable because bytes are numbers
	fmt.Println(string(converted))

	// Converting it with much more readable byte to string
	// first argument is the instance of interface/struct, prefix, indention you can use space or tabs here
	indented, er := json.MarshalIndent(movies, "", " ")
	if er != nil {
		fmt.Println("Error")
	}
	// We can use this instead of converting it to string
	fmt.Printf("%s\n", indented)

	var newMovies []Movie

	//we send an address because we want newMovies will be affected outside of the function
	err = json.Unmarshal(indented, &newMovies)
	fmt.Println(newMovies)
	fmt.Println(newMovies[0].Title)
}
