package main

import (
	contentbase "github.com/thaitania/go-ml-rs-contentbase"
)

func main() {
	println("=========== Example Movie Store =============")

	// Init Categories
	// if true: when you add new ItemProfile and categories doesn't exist, system will add categories to the categories list
	// if false: system will add categories by function NewCategories("id", "title") only!
	cl := contentbase.InitItemCategories(false)
	cl.NewCategories("Action")
	cl.NewCategories("Adventure")
	cl.NewCategories("Drama")
	cl.NewCategories("Horror")
	cl.NewCategories("War")
	cl.NewCategories("Western")
	// println("Categories List:", len(cl.Categories))
	// println(fmt.Sprintf("%v", cl))

	// If you need to check duplicate category, just check error by err := cl.NewCategories("action", "Action")
	// err := cl.NewCategories("action", "Action")
	// if err != nil {
	// 	panic(err)
	// }
	// Init Item Profile
	itp := contentbase.InitItemProfiles()
	itp.NewItemProfile(cl, "m1", "movie_1", []string{"Action", "Adventure", "Western"})
	itp.NewItemProfile(cl, "m2", "movie_2", []string{"Drama"})
	itp.NewItemProfile(cl, "m3", "movie_3", []string{"Horror", "Sci-fi"})
	itp.NewItemProfile(cl, "m4", "movie_4", []string{"Action", "Adventure", "Sci-fi", "War"})
	itp.NewItemProfile(cl, "m5", "movie_5", []string{"Sci-fi", "Adventure"})
	itp.NewItemProfile(cl, "m6", "movie_6", []string{"Sci-fi", "Horror"})

	// Create table Item and Categories
	itp.ItemAttributeValue(cl)

	// Init User Profile
	// upf.UserItemRating()

	// Predict item by using UserProfile frequency
}
