package domain

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string `gorm:"column:name"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
