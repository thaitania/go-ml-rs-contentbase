package main

import (
	"fmt"

	contentbase "github.com/thaitania/go-ml-rs-contentbase"
)

func main() {
	println("=========== Example Movie Store =============")

	// Init Categories
	cl := contentbase.InitItemCategories()
	cl.NewCategories("cat_1", "Action")
	cl.NewCategories("cat_2", "Adventure")
	cl.NewCategories("cat_3", "Drama")
	cl.NewCategories("cat_4", "Horror")
	cl.NewCategories("cat_5", "Sci-fi")
	cl.NewCategories("cat_6", "War")
	cl.NewCategories("cat_7", "Western")
	println("Categories List:")
	println(fmt.Sprintf("%v", cl))

	// Init Item Profile
	itp := contentbase.InitItemProfiles()
	itp.NewItemProfile("m1", "movie_1", []string{"Action", "Adventure", "Western"})
	itp.NewItemProfile("m2", "movie_2", []string{"Drama"})
	itp.NewItemProfile("m3", "movie_3", []string{"Horror", "Sci-fi"})
	itp.NewItemProfile("m4", "movie_4", []string{"Action", "Adventure", "Sci-fi", "War"})

	// Create table Item and Categories
	itp.ItemAttributeValue(cl)

	// Init User Profile
	// upf.UserItemRating()
}
