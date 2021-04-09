package entity

type User struct {
	ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"size:255"`
    Email string `gorm:"size:255"`
}

type Pagination struct {
    Limit int
    Offset int
    Sort string
}