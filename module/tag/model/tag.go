package tagmodel

type Tag struct {
	Id    uint   `json:"id" gorm:"column:id;"`
	Title string `json:"title" gorm:"column:title;unique;not null;"`
}
