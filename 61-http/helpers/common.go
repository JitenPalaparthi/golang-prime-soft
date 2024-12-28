package helpers

import (
	"encoding/json"
	"http-demo/models"
	"log"
	"os"
)

func init() { // it is a very special function, no need to explicitly call it..It gets called automatically
	ChUserData = make(chan *models.User, 5)
	go ProcessUsers("users.dat")
}

var ChUserData chan *models.User

func SaveUser(filename string, user *models.User) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	bytes = append(bytes, byte('\n'))

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func ProcessUsers(fileName string) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()
	for user := range ChUserData {
		bytes, err := json.Marshal(user)
		if err != nil {
			log.Println(err.Error())
			// send error to a channel
		}
		bytes = append(bytes, byte('\n'))

		_, err = f.Write(bytes)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
