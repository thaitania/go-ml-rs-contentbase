package contentbase

import "errors"

// ItemCategories is stuct for contain categories list of item
type ItemCategories struct {
	AutoAddNewCategories bool
	Categories           []ItemCategoriesData
}

// ItemCategoriesData is struct for contain categories detail
type ItemCategoriesData struct {
	ID    string
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
func (itc *ItemCategories) NewCategories(id string, title string) error {
	if id == "" {
		return errors.New("Category must have ID")
	}

	for _, v := range itc.Categories {
		if v.ID == id {
			return errors.New("Categories " + title + " already exist!")
		}
	}
	itc.Categories = append(itc.Categories, ItemCategoriesData{
		ID:    id,
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
