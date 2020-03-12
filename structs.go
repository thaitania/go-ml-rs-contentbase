package contentbase

// UserProfile is struct for contain User Profile Data
type UserProfile struct {
	ID           string                          `json:"id"`            // ID of user
	Item         string                          `json:"item"`          // Item data
	Rating       float32                         `json:"rating"`        // Rating of item, it's must be 0.00 - 5.00
	UserProfiles map[string][]UserProfileContent `json:"user_profiles"` // Profiles of user, generate from System after user pick item
}

// UserProfileContent is function for contain key, value of user profile
type UserProfileContent struct {
	Key   string
	Value string
}

// ItemProfile is struct for contain Item Profile Data
type ItemProfile struct {
	ID       string                          `json:"id"`            // ID of item
	Name     string                          `json:"name"`          // Name of item
	Profiles map[string][]ItemProfileContent `json:"item_profiles"` // Profiles of item
}

// ItemProfileContent is struct for contain Item Profile Content
type ItemProfileContent struct {
	Key   string
	Value string
}
