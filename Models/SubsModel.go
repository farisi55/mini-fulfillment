package Models

type Subs struct {
	SubsId      string   `json:"subsid"`
	Msisdn    string `json:"msisdn"`
	Balance   int `json:"balance"`
	Serviceid   string `json:"serviceid"`
	Sysdate string `json:"sydate"`
	Pin string `json:"pin"`
}

func (b *Subs) TableName() string {
	return "subsinfo"
}

//func (u *User)BeforeCreate(tx *gorm.DB)  (err error){
//	u.Name = Service.Encrypt(u.Name)
//	return
//}
//
//func (u *User)BeforeUpdate(tx *gorm.DB)  (err error){
//	u.Name = Service.Encrypt(u.Name)
//	return
//}
//
//func (u *User) AfterFind(tx *gorm.DB) (err error) {
//	u.Name = Service.Decrypt(u.Name)
//	return
//}
