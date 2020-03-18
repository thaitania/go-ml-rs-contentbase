package main

import (
	"fmt"

	contentbase "github.com/thaitania/go-ml-rs-contentbase"
)

func main() {
	println("=========== Example Movie Store =============")

	// Init Categories
	cl := contentbase.InitItemCategories(false)
	cl.NewCategories("action", "Action")
	cl.NewCategories("advent", "Adventure")
	cl.NewCategories("drama", "Drama")
	cl.NewCategories("horror", "Horror")
	cl.NewCategories("sci_fi", "Sci-fi")
	cl.NewCategories("war", "War")
	cl.NewCategories("western", "Western")

	// If you need to check duplicate category, just check error by err := cl.NewCategories("action", "Action")
	// err := cl.NewCategories("action", "Action")
	// if err != nil {
	// 	panic(err)
	// }

	println("Categories List:")
	println(fmt.Sprintf("%v", cl))

	// Init Item Profile
	itp := contentbase.InitItemProfiles()
	itp.NewItemProfile("m1", "movie_1", []string{"Action", "Adventure", "Western"})
	itp.NewItemProfile("m2", "movie_2", []string{"Drama"})
	itp.NewItemProfile("m3", "movie_3", []string{"Horror", "Sci-fi"})
	itp.NewItemProfile("m4", "movie_4", []string{"Action", "Adventure", "Sci-fi", "War"})
	itp.NewItemProfile("m5", "movie_5", []string{"Sci-fi", "Adventure"})
	itp.NewItemProfile("m6", "movie_6", []string{"Sci-fi", "Horror"})

	// Create table Item and Categories
	itp.ItemAttributeValue(cl)

	// Init User Profile
	// upf.UserItemRating()

	// Predict item by using UserProfile frequency
}
