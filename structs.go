package contentbase

type UserProfile struct {
	UserData map[string][]UserItemRating
}

// UserProfile["UserID"]

// 1. foreach UserProfile ได้แกน y
// 2. plot หัวตารางแกน x โดยใช้ Categories ทั้งหมด
// 3. ดูว่า UserProfile มีข้อมูล Categories ไหนบ้าง และ เขียน overall scroll
// 4. หาค่าของความชอบ ว่า User ชอบคอนเทนต์แบบไหนมากที่สุด
// 5. นำเสนอ Item ที่มีความชอบคล้ายคลึงกับ User Profile ที่สุด แล้วเรียงลำดับ

type UserItemRating struct {
	ItemID string
	Rating int
}
