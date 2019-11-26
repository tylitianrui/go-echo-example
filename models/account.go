package models

type BasicModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt uint64
	UpdatedAt uint64
	Status    uint8
}

type Account struct {
	ID       uint   `gorm:"primary_key"`
	User     string `gorm:"index"`
	Password string `gorm:"size:255"`
	Mobile   string `gorm:"type:varchar(11);unique_index"`
	Status   uint   `gorm:"default:'0'"`
	Created  int
}

func (a Account) TableName() string {
	return "account"
}
