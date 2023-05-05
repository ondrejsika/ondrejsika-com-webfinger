package main

import (
	"fmt"
	"net/http"
)

func main() {
	KEYCLOAK_URL := "https://sso.sikalabs.com/realms/ondrejsika"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resource := r.URL.Query().Get("resource")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"subject":"` + resource + `","links":[{"rel": "http://openid.net/specs/connect/1.0/issuer","href": "` + KEYCLOAK_URL + `"}]}`))
	})

	fmt.Println("Server started on 0.0.0.0:8000, see http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
