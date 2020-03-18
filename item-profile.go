package contentbase

import (
	"errors"
	"fmt"
	"strconv"
)

// ItemList is struct for contain item profile list
type ItemList struct {
	ItemProfile map[string]ItemProfile
}

// ItemProfile is struct for contain item profile data
type ItemProfile struct {
	ItemID     string
	Title      string
	Categories []string
}

// ItemAttributeValueData is struct for contain ItemAttributeValue list
type ItemAttributeValueData struct {
	Categories []string
	Value      map[string][]int
}

// InitItemProfiles is function for init ItemProfile memory
func InitItemProfiles() *ItemList {
	return &ItemList{
		ItemProfile: make(map[string]ItemProfile),
	}
}

// NewItemProfile is function for add new item in list
func (ip *ItemList) NewItemProfile(cl *ItemCategories, itemID string, title string, categories []string) error {
	if itemID == "" {
		return errors.New("ItemProfile must have ID")
	}
	if cl.AutoAddNewCategories == true {
		for _, e := range categories {
			if indexOfCategories(e, cl.Categories) == -1 {
				cl.NewCategories(e)
			}
		}
	}
	ip.ItemProfile[itemID] = ItemProfile{
		ItemID:     itemID,
		Title:      title,
		Categories: categories,
	}
	return nil
}

// ItemAttributeValue is function for render Item-Attribute Value (IAB)
func (ip *ItemList) ItemAttributeValue(cl *ItemCategories) (*ItemAttributeValueData, error) {
	catArr := []string{}
	if len(cl.Categories) == 0 {
		return &ItemAttributeValueData{}, errors.New("Categories list is empty: length=" + strconv.Itoa(len(cl.Categories)))
	}
	for _, e := range cl.Categories {
		catArr = append(catArr, e.Title)
	}
	itemList := make(map[string][]int)
	for _, ee := range ip.ItemProfile {
		x := make([]int, len(catArr))
		for _, eee := range ee.Categories {
			ido := indexOf(eee, catArr)
			if ido >= 0 {
				x[ido] = 1
			}
		}
		itemList[ee.ItemID] = x
	}

	return &ItemAttributeValueData{Categories: catArr, Value: itemList}, nil
}

// PrintGUIItemAttributeValue is function for print ItemAttributeValue in Command line interface
func PrintGUIItemAttributeValue(iacd *ItemAttributeValueData) error {
	if len(iacd.Categories) == 0 {
		return errors.New("Categories list is empty: length=" + strconv.Itoa(len(iacd.Categories)))
	}
	if len(iacd.Value) == 0 {
		return errors.New("ItemProfile list is empty: length=" + strconv.Itoa(len(iacd.Value)))
	}
	println(fmt.Sprintf("%v", iacd.Categories))
	for k, v := range iacd.Value {
		println(k, fmt.Sprintf("%v", v))
	}

	return nil
}
