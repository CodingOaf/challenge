package controller

import (
	"challenge/models"
	"encoding/json"
	"io"
	"net/http"
)

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract NewRequest to a function
	req, err := NewRequest("GET", w, "/projects", nil)
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

	var ProjectsResponse models.Workspaces
	if err := json.Unmarshal(resBody, &ProjectsResponse); err != nil {
		panic(err)
	}
	// format the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ProjectsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
