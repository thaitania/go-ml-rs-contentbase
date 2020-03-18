package main

import (
	"fmt"
	"log"

	contentbase "github.com/thaitania/go-ml-rs-contentbase"
)

func main() {
	println("=========== Example Movie Store =============")

	// Init Categories
	// If passing function with: true: when you add new ItemProfile and categories doesn't exist, system will add categories to the categories list
	// If passing function with: false: system will only add categories by function NewCategories("id", "title")
	cl := contentbase.InitItemCategories(true)
	cl.NewCategories("Action")
	cl.NewCategories("Adventure")
	cl.NewCategories("Drama")
	cl.NewCategories("Horror")
	cl.NewCategories("War")
	cl.NewCategories("Western")

	// If you need to check duplicate category, just check error by err := cl.NewCategories("action", "Action")
	// err := cl.NewCategories("action", "Action")
	// if err != nil {
	// 	panic(err)
	// }

	// Init Item Profile
	itp := contentbase.InitItemProfiles()
	itp.NewItemProfile(cl, "m1", "movie_1", []string{"Action"})
	// If ItemProfile is duplicate, system will replace with new ItemProfile Data instead
	itp.NewItemProfile(cl, "m1", "movie_1", []string{"Action", "Adventure", "Western"})
	itp.NewItemProfile(cl, "m2", "movie_2", []string{"Drama"})
	itp.NewItemProfile(cl, "m3", "movie_3", []string{"Horror", "Sci-fi"})
	itp.NewItemProfile(cl, "m4", "movie_4", []string{"Action", "Adventure", "Sci-fi", "War"})
	itp.NewItemProfile(cl, "m5", "movie_5", []string{"Sci-fi", "Adventure"})
	itp.NewItemProfile(cl, "m6", "movie_6", []string{"Action", "Sci-fi", "Horror"})

	// Create table Item and Categories
	iav, err := itp.ItemAttributeValue(cl)
	if err != nil {
		log.Fatal(err)
	}

	err = contentbase.PrintGUIItemAttributeValue(iav)
	if err != nil {
		log.Fatal(err)
	}

	// Init User Profile
	upf := contentbase.InitUserProfile()
	upf.NewUserProfile("user1", "m1", 5)
	upf.NewUserProfile("user1", "m6", 5)

	// upf.NewUserProfile("user2", "m2", 5)
	// upf.NewUserProfile("user2", "m4", 1)
	// upf.NewUserProfile("user2", "m5", 1)
	// println(fmt.Sprintf("%v", upf))

	// Generate score of sequence categories of user
	// Frequency is calculate by frequency of Categories only. (calculate without Rating weight)
	// score result must be 0.00 - 1.00 (result meaning: 0.00 is user never watch movie in category, 1.00 is user watch movie in category very often)
	println("============== GetScoreByCategoriesFrequency ==============")
	uppf, _ := contentbase.GetScoreByCategoriesFrequency(upf, iav)
	for k, v := range uppf.UserData {
		println("Result=", k, fmt.Sprintf("%v", v))
	}

	// println("============== GetScoreByCategoriesFrequencyAndRatingWeight ==============")
	// uppfr, _ := contentbase.GetScoreByCategoriesFrequencyWithRating(upf, iav)
	// for k, v := range uppfr.UserData {
	// 	println("Result=", k, fmt.Sprintf("%v", v))
	// }
}
