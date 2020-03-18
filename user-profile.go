package contentbase

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
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
			for kkk, vvv := range iav.Value[vv.ItemID] {
				tmp[kkk] = tmp[kkk] + float32(vvv)
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
	}
	return upps, nil
}

// GetScoreByCategoriesFrequencyWithRating is function for predict by frequency and rating
func GetScoreByCategoriesFrequencyWithRating(upf *UserList, iav *ItemAttributeValueData) (*UserProfilePredictionCategories, error) {
	upps := &UserProfilePredictionCategories{
		UserData: make(map[string][]UserProfilePredictionCategoriesData),
	}
	pdfq := make(map[string][]float32)
	for k, v := range upf.UserData {
		tmp := make([]float32, len(iav.Categories))
		for _, vv := range v {
			// println(fmt.Sprintf("%v", iav.Value[vv.ItemID]))
			for kkk, vvv := range iav.Value[vv.ItemID] {
				tmp[kkk] = tmp[kkk] + float32(vvv)
				// println(k+"."+vv.ItemID, tmp[kkk], strconv.Itoa(vvv), "rating="+strconv.Itoa(vv.Rating))
			}
			println(k+"."+vv.ItemID, fmt.Sprintf("%v", iav.Value[vv.ItemID]), "rating="+strconv.Itoa(vv.Rating))
		}
		pdfq[k] = tmp
	}

	for k, v := range pdfq {
		ratingRatio := 0
		// for _, e := range upf.UserData[k] {
		// 	ratingRatio += e.Rating
		// }
		for kk := range v {
			pdfq[k][kk] = (float32(pdfq[k][kk]) / float32(len(upf.UserData[k]))) * (float32(ratingRatio) / float32(5*len(upf.UserData[k])))
		}
		// println(k, fmt.Sprintf("%v", ratingRatio))
		println(k, fmt.Sprintf("%v", pdfq[k]))

		// for kk, vv := range v {
		// 	// println(fmt.Sprintf("%v", ratingRatio))
		// 	// println(fmt.Sprintf("%.2f", float32(vv/float32(len(upf.UserData[k])))), fmt.Sprintf("%.2f", float32(ratingRatio*len(upf.UserData[k]))))
		// 	upps.UserData[k] = append(upps.UserData[k], UserProfilePredictionCategoriesData{
		// 		CategoriesID: iav.Categories[kk],
		// 		Score:        float32(vv/float32(len(upf.UserData[k]))) * (0.5 * 2),
		// 	})
		// }
		// sort.Slice(upps.UserData[k], func(i, j int) bool {
		// 	return upps.UserData[k][i].Score > upps.UserData[k][j].Score
		// })
	}
	return upps, nil
}
