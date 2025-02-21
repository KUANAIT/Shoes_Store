package models

type ShippingInfo struct {
	FreeShipping      bool   `json:"free_shipping" bson:"free_shipping"`
	ExpeditedShipping string `json:"expedited_shipping" bson:"expedited_shipping"`
}

type Shoe struct {
	ID            string       `json:"id" bson:"id"`
	Name          string       `json:"name" bson:"name"`
	Brand         string       `json:"brand" bson:"brand"`
	Size          []int        `json:"size" bson:"size"`
	Price         float64      `json:"price" bson:"price"`
	Material      string       `json:"material" bson:"material"`
	Color         []string     `json:"color" bson:"color"`
	ReleaseDate   string       `json:"release_date" bson:"release_date"`
	Discount      float64      `json:"discount" bson:"discount"`
	StockQuantity int          `json:"stock_quantity" bson:"stock_quantity"`
	Rating        float64      `json:"rating" bson:"rating"`
	SKU           string       `json:"sku" bson:"sku"`
	Weight        string       `json:"weight" bson:"weight"`
	Features      []string     `json:"features" bson:"features"`
	Category      string       `json:"category" bson:"category"`
	ShippingInfo  ShippingInfo `json:"shipping_info" bson:"shipping_info"`
}

type User struct {
	ID               string   `json:"id" bson:"id"`
	Username         string   `json:"username" bson:"username"`
	Email            string   `json:"email" bson:"email"`
	FirstName        string   `json:"first_name" bson:"first_name"`
	LastName         string   `json:"last_name" bson:"last_name"`
	Password         string   `json:"password" bson:"password"`
	ShippingAddress  string   `json:"shipping_address" bson:"shipping_address"`
	Phone            string   `json:"phone" bson:"phone"`
	RegistrationDate string   `json:"registration_date" bson:"registration_date"`
	OrderIDs         []string `json:"order_ids" bson:"order_ids"`
}

type Order struct {
	ID              string       `json:"id" bson:"id"`
	UserID          string       `json:"user_id" bson:"user_id"`
	OrderDate       string       `json:"order_date" bson:"order_date"`
	ShippingAddress string       `json:"shipping_address" bson:"shipping_address"`
	Items           []Item       `json:"items" bson:"items"`
	TotalPrice      float64      `json:"total_price" bson:"total_price"`
	Status          string       `json:"status" bson:"status"`
	ShippingInfo    ShippingInfo `json:"shipping_info" bson:"shipping_info"`
}

type Item struct {
	ShoeID   string  `json:"shoe_id" bson:"shoe_id"`
	Quantity int     `json:"quantity" bson:"quantity"`
	Price    float64 `json:"price" bson:"price"`
}
