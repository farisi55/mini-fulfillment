package Service

type LoginService interface {
	LoginUser(msisdn string, pin string) bool
}

type LoginInformation struct {
	msisdn string
	pin string
}

func StaticLoginService() LoginService  {
	return &LoginInformation{
		msisdn:    "0818123",
		pin: "123",
	}
}

func (l *LoginInformation)LoginUser(msisdn string, pin string)bool  {
	return l.msisdn == msisdn && l.pin == pin
}
