package entity

type Task struct {
	ID          int64  `gorm:"primaryKey:auto_increment" json:"-"`
	Title       string `gorm:"type:varchar(100)" json:"-"`
	Description string `gorm:"type:text" json:"-"`
	UserID      int64  `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"-"`
}
