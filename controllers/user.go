package controllers

import (
	"fmt"
	"io"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "MASUK BOS")

	fmt.Println("masuk sini")
}

var AuthUser = func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "MASUK AUTH")
}
