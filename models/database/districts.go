package database

type Districts struct {
	Id        string `json:"id" gorm:"type:char(7)"`
	RegencyId string `json:"regencyId"`
	Name      string `json:"name"`
}
