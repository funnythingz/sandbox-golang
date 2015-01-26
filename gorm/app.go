package main

import (
	"./db"
	"./models"
	"github.com/manveru/faker"
	"log"
)

func main() {

	dbmap.DBConfiguer()
	fake, _ := faker.New("en")

	person := model.Person{Name: fake.Name()}

	log.Println(dbmap.DB.NewRecord(person))
	log.Println(dbmap.DB.Create(&person))
	log.Println(dbmap.DB.NewRecord(person))
	log.Println(dbmap.DB.Save(&person))

	log.Println("first")
	log.Println(dbmap.DB.First(&person))
	log.Println(person)

	var persons []model.Person
	dbmap.DB.Find(&persons)

	log.Println(persons)

	for k, v := range persons {
		log.Println(k, v.Name)
	}

}
