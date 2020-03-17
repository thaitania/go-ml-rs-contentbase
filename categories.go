package contentbase

import "errors"

// ItemCategories is stuct for contain categories list of item
type ItemCategories struct {
	Categories []ItemCategoriesData
}

// ItemCategoriesData is struct for contain categories detail
type ItemCategoriesData struct {
	ID    string
	Title string
}

// InitItemCategories is function for init ItemCategories in memory
func InitItemCategories() *ItemCategories {
	return &ItemCategories{}
}

// NewCategories is function for add new categories data in category list
func (itc *ItemCategories) NewCategories(id string, title string) error {
	if id == "" {
		return errors.New("Category must have ID")
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
