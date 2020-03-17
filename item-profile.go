package contentbase

import (
	"errors"
	"fmt"
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

// InitItemProfiles is function for init ItemProfile memory
func InitItemProfiles() *ItemList {
	return &ItemList{
		ItemProfile: make(map[string]ItemProfile),
	}
}

// NewItemProfile is function for add new item in list
func (ip *ItemList) NewItemProfile(itemID string, title string, categories []string) error {
	if itemID == "" {
		return errors.New("ItemProfile must have ID")
	}
	ip.ItemProfile[itemID] = ItemProfile{
		ItemID:     itemID,
		Title:      title,
		Categories: categories,
	}
	return nil
}

// ItemAttributeValue is function for render Item-Attribute Value (IAB)
func (ip *ItemList) ItemAttributeValue(cl *ItemCategories) error {
	catKeyArr := []string{}
	catArr := []string{}
	for _, e := range cl.Categories {
		catKeyArr = append(catKeyArr, e.ID)
		catArr = append(catArr, e.Title)
	}
	println("==============================")
	println(fmt.Sprintf("%v", catArr))
	println("==============================", len(ip.ItemProfile))
	itemList := make(map[string][]int)
	for _, ee := range ip.ItemProfile {
		x := make([]int, len(catArr))
		for _, eee := range ee.Categories {
			ido := indexOf(eee, catArr)
			x[ido] = 1
		}
		itemList[ee.ItemID] = x
		println(ee.ItemID, fmt.Sprintf("%v", itemList[ee.ItemID]))
	}

	return nil
}
