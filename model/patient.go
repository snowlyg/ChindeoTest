package model

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

var PatientCount int

func CreatePatient(userId int) *Patients {
	patient := &Patients{
		Username:  Fake.Name(),
		Nickname:  Fake.Name(),
		Phone:     Fake.PhoneNumber(),
		Email:     Fake.Email(),
		Sex:       rand.Intn(1),
		Password:  Fake.PhoneNumber(),
		Status:    true,
		AvatarURL: Fake.URL(),
		OpenID:    Fake.PhoneNumber(),
		UnionID:   Fake.PhoneNumber(),
		Country:   Fake.Country(),
		Province:  Fake.Country(),
		City:      Fake.City(),
		Mac:       Fake.City(),
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := DB.Create(&patient).Error; err != nil {
		fmt.Println(fmt.Sprintf("patient create error :%v", err))
	}

	userPatient := UserPatient{PatientID: patient.ID, UserID: userId, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := DB.Table("user_patient").Create(&userPatient).Error; err != nil {
		fmt.Println(fmt.Sprintf("userPatient create error :%v", err))
	}
	PatientCount++
	return patient
}
