package Models

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"theapp/Config"
)

func GetAllUser(subs *[]Subs) (err error)  {
	if err = Config.DB.Find(subs).Error; err != nil{
		return err
	}
	return nil
}

func PurchasePacket(subs *Subs) (err error)  {
	if err = Config.DB.Create(subs).Error; err != nil{
		return err
	}
	sentry.CaptureMessage("User created")
	return nil
}

func GetSubsInfo(subs *Subs, id string)(err error)  {
	if err = Config.DB.Where("subs_id=?", id).First(subs).Error; err != nil{
		return err
	}
	return nil
}

func UpdateUser(subs *Subs, id string) (err error) {
	Config.DB.Save(subs)
	return nil
}

func UnregPacket(subs *Subs,  id string, serviceid string)(err error)  {
	fmt.Println(">>>>>",serviceid,id)
	Config.DB.Where("subs_id=? and serviceid=?", id, serviceid).Delete(subs)
	return nil
}

