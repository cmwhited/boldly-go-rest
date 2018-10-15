package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/gorilla/mux"
)

// Http Handler for the Ping Endpoint (/api/v1/ping)
// When Ping is hit, Call the DoPing() method and return the results.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	ping := DoPing()
	js, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

// Http Handler for the GetBank endpoint (/api/v1/user/{owningUserId}/bank/{bankId})
// Grab the owningUserId and bankId out of the route params and use them to get a Bank record
func GetBankHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	owningUserId := params["owningUserId"]      // get the owning user id from request route params
	bankId := params["bankId"]                  // get the bank id from the request route params
	_bankId := uuid.FromStringOrNil(bankId)     // convert bank id string to uuid
	bank, err := GetBank(owningUserId, _bankId) // get bank record
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(bank) // convert bank struct to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

// Http Handler for the SaveBank endpoint (/api/v1/bank)
// Gets the bank out of the request body and saves it
func SaveBankHandler(w http.ResponseWriter, r *http.Request) {
	var bank Bank
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(b, &bank)
	// call save bank service call
	created, err := bank.SaveBank()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// marshal to JSON && return the created bank
	js, err := json.Marshal(created) // convert bank struct to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}
