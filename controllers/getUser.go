package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"github.com/sthakur410/testProj/models"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) Response {
	return ReturnResponseFromCodeMessage(200, "sucsess", models.AllUsers)
}

func AddUser(w http.ResponseWriter, r *http.Request) Response {
	user, isError, msg := validatePayload(w, r)
	if isError {
		return ReturnResponseFromCodeMessage(500, msg, nil)
	}
	models.AllUsers[user.Id] = user
	return ReturnResponseFromCodeMessage(200, "sucsess", nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) Response {
	user, isError, msg := validatePayload(w, r)
	if isError {
		return ReturnResponseFromCodeMessage(500, msg, nil)
	}
	for _, value := range models.AllUsers {
		if value.Id == user.Id {
			models.AllUsers[value.Id] = user
		}
	}
	return ReturnResponseFromCodeMessage(200, "success", nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) Response {
	params := mux.Vars(r)
	id := params["id"]
	for _, value := range models.AllUsers {
		if value.Id == cast.ToInt(id) {
			delete(models.AllUsers, value.Id)
		}
	}
	return ReturnResponseFromCodeMessage(200, "success", nil)
}

func validatePayload(w http.ResponseWriter, r *http.Request) (user models.User, isError bool, msg string) {
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshalling body", err)
		msg = "Payload is not correct"
		isError = true
		return
	}
	return
}
