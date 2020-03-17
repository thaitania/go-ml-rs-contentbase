package contentbase

import "errors"

// UserList is struct for contain user profile data list
type UserList struct {
	UserData map[string][]UserProfile
}

// UserProfile is struct for contain user profile usage data
type UserProfile struct {
	ItemID string
	Rating int
}

// InitUserProfile is function for init UserProfile memory
func InitUserProfile() *UserList {
	return &UserList{
		UserData: make(map[string][]UserProfile),
	}
}

// NewUserProfile is function for add new item in list
func (ul *UserList) NewUserProfile(userID string, itemID string, rating int) error {
	if userID == "" {
		return errors.New("UserProfile must have ID of user")
	}
	// ul.UserProfile[userID] = []UserProfile{
	// 	ItemID: itemID,
	// 	Rating: Rating,
	// }
	return nil
}
