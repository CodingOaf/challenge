package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"challenge/models"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract NewRequest to a function
	req, err := http.NewRequest("GET", ASANA_API_URL+"/users", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+r.Header.Get("Authorization"))
	// TODO: Change endcoding of token
	//w.Write([]byte("Before Client.Do"))
	// TODO: Extract Client.Do to a function
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var UsersResponse models.Users
	if err := json.Unmarshal(resBody, &UsersResponse); err != nil {
		panic(err)
	}
	// format the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(UsersResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
