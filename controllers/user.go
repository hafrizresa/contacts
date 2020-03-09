package controllers

import (
	"fmt"
	"io"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Response) {
	io.WriteString(w, "MASUK BOS")

	fmt.Println("masuk sini")
}
