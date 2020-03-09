package controllers

import (
	"net/http"

	"github.com/hafrizresa/contacts/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	utils.Response(w, utils.Message(true, "SUCCESS CREATE ACCOUNT"))
}

var AuthUser = func(w http.ResponseWriter, r *http.Request) {
	utils.Response(w, utils.Message(true, "SUCCESS LOGIN ACCOUNT"))

}
