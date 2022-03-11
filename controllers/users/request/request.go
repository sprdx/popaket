package requests

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Username string `gorm:"unique" json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	MSISDN   string `gorm:"unique" json:"msisdn" form:"msisdn"`
}

type LoginRequest struct {
	MSISDN   string `json:"msisdn" form:"msisdn"`
	Password string `json:"password" form:"password"`
}
