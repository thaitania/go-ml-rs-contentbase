package contentbase

import (
	"errors"
	"fmt"
	"sort"
)

// UserList is struct for contain user profile data list
type UserList struct {
	UserData map[string][]UserProfileItem
}

// UserProfileItem is struct for contain user profile usage data
type UserProfileItem struct {
	ItemID string
	Rating int
}

// UserProfilePredictionCategories is struct for contain user prediction score
type UserProfilePredictionCategories struct {
	UserData map[string][]UserProfilePredictionCategoriesData
}

// UserProfilePredictionCategoriesData is struct for contain user prediction score Data
type UserProfilePredictionCategoriesData struct {
	CategoriesID string
	Score        float32
}

// InitUserProfile is function for init UserProfileItem memory
func InitUserProfile() *UserList {
	return &UserList{
		UserData: make(map[string][]UserProfileItem),
	}
}

// NewUserProfile is function for add new item in list
func (ul *UserList) NewUserProfile(userID string, itemID string, rating int) error {
	if userID == "" {
		return errors.New("UserProfile must have 'ID' of user")
	}
	if itemID == "" {
		return errors.New("UserProfile must have 'ItemID' of the item")
	}

	if ul.UserData[userID] != nil {
		ul.UserData[userID] = append(ul.UserData[userID], UserProfileItem{ItemID: itemID, Rating: rating})
	} else {
		ul.UserData[userID] = []UserProfileItem{UserProfileItem{ItemID: itemID, Rating: rating}}
	}
	// ul.UserProfile[userID] = []UserProfile{
	// 	ItemID: itemID,
	// 	Rating: Rating,
	// }
	return nil
}

// GetScoreByCategoriesFrequency is function for predict by frequency only (predict without rating)
func GetScoreByCategoriesFrequency(upf *UserList, iav *ItemAttributeValueData) (*UserProfilePredictionCategories, error) {
	upps := &UserProfilePredictionCategories{
		UserData: make(map[string][]UserProfilePredictionCategoriesData),
	}
	pdfq := make(map[string][]float32)
	for k, v := range upf.UserData {
		tmp := make([]float32, len(iav.Categories))
		for _, vv := range v {
			for k, vvv := range iav.Value[vv.ItemID] {
				tmp[k] = tmp[k] + float32(vvv)
			}
			// println(k+"."+vv.ItemID, fmt.Sprintf("%v", iav.Value[vv.ItemID]), "rating="+strconv.Itoa(vv.Rating))
		}
		pdfq[k] = tmp
	}

	for k, v := range pdfq {
		// println(k, fmt.Sprintf("%v", v))
		for kk, vv := range v {
			v[kk] = vv / float32(len(upf.UserData[k]))

			upps.UserData[k] = append(upps.UserData[k], UserProfilePredictionCategoriesData{
				CategoriesID: iav.Categories[kk],
				Score:        v[kk],
			})
		}
		sort.Slice(upps.UserData[k], func(i, j int) bool {
			return upps.UserData[k][i].Score > upps.UserData[k][j].Score
		})
		println("Result=", k, fmt.Sprintf("%v", upps.UserData[k]))
	}
	return upps, nil
}
