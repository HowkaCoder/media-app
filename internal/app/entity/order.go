package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	Username      string         `gorm:"not null" json:"username"`
	Status        string         `gorm:"not null" json:"status"`
	DeliveryDate  time.Time      `gorm:"not null" json:"delivery_date"`
	PickupPoint   string         `gorm:"not null" json:"pickup_point"`
	Products      []OrderProduct `gorm:"foreignKey:OrderID" json:"products"`
	PaymentMethod string         `gorm:"not null" json:"payment_method"`
	TotalAmount   uint           `gorm:"not null" json:"total_amount"`
}

type OrderProduct struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID     uint   `gorm:"not null" json:"order_id"`
	ProductID   uint   `gorm:"not null" json:"product_id"`
	Count       uint   `gorm:"not null" json:"count"`
	Description string `gorm:"not null" json:"description"`
	Title       string `gorm:"not null" json:"title"`
	Price       uint   `gorm:"not null" json:"price"`
	Discount    uint   `gorm:"null" json:"discount"`
	Image       Image  `gorm:"foreignKey:ProductID;references:ProductID" json:"image"`
}

type Metric struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	MetricType  string    `gorm:"not null" json:"metric_type"` // e.g., "order_count", "total_revenue"
	Value       float64   `gorm:"not null" json:"value"`
	Date        time.Time `gorm:"not null" json:"date"`
	Description string    `gorm:"not null" json:"description"`
}
