package mocks

import (
	"log"
	"server/initializers"
	"server/models"
	"time"
)

// ALARM_LIST contient les valeurs mockées pour les alarmes
var ALARM_LIST = []models.Alarm{
	{Name: "Projet d'intégration", RingDate: time.Now().Add(24 * time.Hour), Location: "EPHEC", Ringtone: "Ringtone 1", IsActive: true},
	{Name: "Sécurité des réseaux (Théorie)", RingDate: time.Now().Add(48 * time.Hour), Location: "EPHEC", Ringtone: "Ringtone 2", IsActive: false},
	{Name: "Sécurité des réseaux (Pratique)", RingDate: time.Now().Add(72 * time.Hour), Location: "EPHEC", Ringtone: "Ringtone 3", IsActive: true},
}

func InsertMockedAlarms() { // VALEURS MOCKEES : A RETIRER EN PROD !!!
	err := initializers.DB.Where("1 = 1").Delete(models.Alarm{}).Error
	if err != nil {
		log.Fatal("Could not delete alarms :", err)
	}
	alarms := ALARM_LIST
	for _, alarm := range alarms {
		err := initializers.DB.Create(&alarm).Error
		if err != nil {
			log.Fatal("Could not create alarm :", err)
		}
	}
}
