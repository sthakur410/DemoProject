package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	c "github.com/sthakur410/testProj/controllers"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Start of Function UserHandler")

	switch true {
	case strings.Contains(r.Method, "GET"):
		URLReturnResponseJson(w, c.GetAllUser(w, r))
	case strings.Contains(r.Method, "POST"):
		URLReturnResponseJson(w, c.AddUser(w, r))
	case strings.Contains(r.Method, "PUT"):
		URLReturnResponseJson(w, c.UpdateUser(w, r))
	default:
		log.Println("Handler not found", r.Method)
		URLReturnResponseJson(w, c.ReturnResponseFromCodeMessage(500, "bad request", nil))
	}

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Start of Function DeleteUserHandler")

	switch true {
	case strings.Contains(r.Method, "DELETE"):
		URLReturnResponseJson(w, c.DeleteUser(w, r))
	default:
		log.Println("Handler not found", r.Method)
		URLReturnResponseJson(w, c.ReturnResponseFromCodeMessage(500, "bad request", nil))
	}

}

func URLReturnResponseJson(w http.ResponseWriter, data c.Response) {
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)
	w.Write(returnJson)
}
