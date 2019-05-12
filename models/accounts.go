package models

type Accounts struct {
	EnviayaAccount interface{} `json:"enviaya_account,omitempty"` // You may create several Billing accounts for your user. This can be useful for internal billing and controlling purposes or if you want to use different carrier and service configurations by account.
	CarrierAccount interface{} `json:"carrier_account,omitempty"` // If you want to use your own carrier account, you can transmit it here. Please note that you first need to configure your carrier accounts and access data in your Pak Mail Samara user on https://api224.pakmail.com.mx/.
}
