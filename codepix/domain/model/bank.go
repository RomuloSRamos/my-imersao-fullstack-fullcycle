package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json: "code"  gorm:"type:verchar(20)" valid:"notnull"`
	Name     string     `json: "code"  gorm:"type:verchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid: "-"`
}

func (Bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(Bank)
	if err != nil {
		return err
	}

	return nil
}

func Newbank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}
	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
