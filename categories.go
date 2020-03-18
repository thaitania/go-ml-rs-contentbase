package contentbase

import "errors"

// ItemCategories is stuct for contain categories list of item
type ItemCategories struct {
	AutoAddNewCategories bool
	Categories           []ItemCategoriesData
}

// ItemCategoriesData is struct for contain categories detail
type ItemCategoriesData struct {
	Title string
}

// InitItemCategories is function for init ItemCategories in memory
//
// if autoAddNewCategories = true: system will auto add new Categories if the category doesn't exist
func InitItemCategories(autoAddNewCategories bool) *ItemCategories {
	return &ItemCategories{
		AutoAddNewCategories: autoAddNewCategories,
	}
}

// NewCategories is function for add new categories data in category list
func (itc *ItemCategories) NewCategories(title string) error {
	if title == "" {
		return errors.New("Category must have ID")
	}

	if indexOfCategories(title, itc.Categories) != -1 {
		return errors.New("Categories " + title + " already exist!")
	}

	itc.Categories = append(itc.Categories, ItemCategoriesData{
		Title: title,
	})
	return nil
}

// indexOf is function for find index by array
func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func indexOfCategories(word string, data []ItemCategoriesData) int {
	for k, v := range data {
		if v.Title == word {
			return k
		}
	}
	return -1
}
