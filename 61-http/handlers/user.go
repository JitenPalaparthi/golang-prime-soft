package handlers

import (
	"encoding/json"
	"http-demo/helpers"
	"http-demo/models"
	"io"
	"log"
	"net/http"
	"time"
)

func CreatUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := new(models.User)
		bytes, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte("invalid body data"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bytes, user)

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte("invalid body data"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = user.Validate()
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user.Status = "active"
		user.LastModified = time.Now().Unix()

		// err = helpers.SaveUser("users.dat", user)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	w.Write([]byte("invalid body data"))
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// } else {
		// 	w.Write([]byte("user successfully stored in the file"))
		// 	w.WriteHeader(http.StatusCreated)
		// 	return
		// }
		helpers.ChUserData <- user
		w.Write([]byte("user successfully stored in the file"))
		w.WriteHeader(http.StatusCreated)
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}
