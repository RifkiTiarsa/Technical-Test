package config

const (
	ApiGroup = "/api/v1"

	PostCustomerRegister = "/customer/register"
	PostCustomerLogin    = "/customer/login"

	PostProduct   = "/product"
	GetProductId  = "/product/:id"
	GetAllProduct = "/product"
	DelProduct    = "/product/:id"
	PutProduct    = "/product/:id"

	PostMerchant = "/merchant"

	PostTopup         = "/topup"
	PostTopupCallback = "/topup/callback"

	PostCustomerLogout = "/customer/logout"
)
