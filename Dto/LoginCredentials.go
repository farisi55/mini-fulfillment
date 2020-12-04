package Dto

type LoginCredentials struct {
	Msisdn    string `form:"msisdn"`
	Pin string `form:"pin"`
}
