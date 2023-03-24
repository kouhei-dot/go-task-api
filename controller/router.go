package controller

import (
	"net/http"
)

func Router() {
	http.HandleFunc("/", Handler)
}
