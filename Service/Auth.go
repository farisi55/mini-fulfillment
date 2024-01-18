package Service

type AuthInformation struct {
	Msisdn string
	Permission []string
}

func StaticAuthService() []AuthInformation{
	var authlist = []AuthInformation{
		AuthInformation{
			Msisdn:      "0818123",
			Permission: []string{"GET", "POST","PUT","DELETE"},
		},
		AuthInformation{
			Msisdn:      "0818321",
			Permission: []string{"GET", "POST"},
		},
	}
	return authlist


}