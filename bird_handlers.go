package main

import (
	"encoding/json"
	"fmt"
	"github.com/gojektech/proctor/proctord/logger"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	birds, err := store.GetBirds()
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		logger.Error(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := Bird{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	err = store.CreateBird(&bird)

	if err != nil {
		fmt.Println(err)
	}

	//logger.Info("Bird created %v", bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func updateBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := Bird{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	err = store.UpdateBird(&bird)

	//logger.Info("Bird updated %v", bird)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
