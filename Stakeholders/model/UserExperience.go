package model

type UserExperience struct {
	ID     int `json:"id" gorm:"column:Id;primaryKey"`
	UserID int `json:"userId" gorm:"column:UserId"`
	XP     int `json:"xp" gorm:"column:XP"`
	Level  int `json:"level" gorm:"column:Level"`
}
