package main

import "net/http"

func Secure(w http.ResponseWriter, r *http.Request){
	secret := "secret123"
	key := r.Header.Get("X-API-Key")
	if key != secret{
		http.Error(w, "you are not authorized to access this page", http.StatusUnauthorized)
		return
	}else{
		http.Error(w, "Welcome", http.StatusAccepted)
	}

	
}