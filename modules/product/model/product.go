package productmodel

type Product struct {
	Id          uint   `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;unique;not null;"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
}
