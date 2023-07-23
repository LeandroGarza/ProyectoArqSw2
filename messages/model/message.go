package model

type Message struct {
	Id        int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Userid    int    `gorm:"type:int;not null"`
	Itemid    string `gorm:"type:varchar(64);not null"`
	Content   string `gorm:"type:varchar(512);not null"`
	Createdat string `gorm:"type:datetime"`
}

type Messages []Message
