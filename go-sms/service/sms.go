package service

import (
	"LHXHL/go-sms/model"
	"github.com/matoous/go-nanoid/v2"
	"github.com/spf13/viper"
	"regexp"
)

type SmsCore interface {
	IfExist(phone string) bool
	GenCode(phone string) (string, error)
	ValidPhone(phone string) bool
	ValidCode(phone string, code string) bool
	ClearAll()
	Total() int
}

type SmsCoreImpl struct {
}

func (s SmsCoreImpl) IfExist(phone string) bool {
	get, err := model.DbNow.Get(phone)
	if err != nil {
		return false
	}
	if get == nil {
		return false
	}
	return true
}

func (s SmsCoreImpl) GenCode(phone string) (string, error) {
	newCode := genWay()
	err := model.DbNow.Set(phone, []byte(newCode))
	if err != nil {
		return "", err
	}

	err1 := model.DbNow.Expire(phone, viper.GetInt("EXPIRE"))
	if err != nil {
		return "", err1
	}
	TxSendSms(phone, newCode)
	return newCode, nil
}

func (s SmsCoreImpl) ValidPhone(phone string) bool {
	rePhone := `1[3-9]\d{9}$`
	return regexp.MustCompile(rePhone).MatchString(phone)
}

func (s SmsCoreImpl) ValidCode(phone string, code string) bool {
	if !s.ValidPhone(phone) {
		return false
	}

	if !s.IfExist(phone) {
		return false
	}

	get, err := model.DbNow.Get(phone)
	if err != nil {
		return false
	}
	if string(get) == code {
		err := model.DbNow.Del(phone)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func (s SmsCoreImpl) ClearAll() {
	err := model.DbNow.FLush()
	if err != nil {
		return
	}
}

func (s SmsCoreImpl) Total() int {
	i, err := model.DbNow.Len()
	if err != nil {
		return 0
	}
	return i
}

func genWay() string {
	generate, err := gonanoid.Generate("0123456789", 6)
	if err != nil {
		return ""
	}
	return generate
}

var (
	Sms SmsCore = SmsCoreImpl{}
)
