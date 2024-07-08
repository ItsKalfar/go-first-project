package handlers

import (
	"encoding/json"
	"firstproject/cmd/utils"
	"net/http"
)

type User struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       uint   `json:"age"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    data := []User{
        {
            FirstName: "Sagar",
            LastName:  "Gund",
            Age:       23,
        },
        {
            FirstName: "Sanket",
            LastName:  "Purohit",
            Age:       28,
        },
        {
            FirstName: "Saket",
            LastName:  "karlekar",
            Age:       23,
        },
        {
            FirstName: "Deepti",
            LastName:  "Chaturvedi",
            Age:       23,
        },
    }
    utils.SendResponse(w, http.StatusOK, true, "Users fetched successfully", data)
    
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
    var user User;
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        utils.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
        return
    }
    defer r.Body.Close() 
    utils.SendResponse(w, http.StatusCreated, true, "User received successfully", user)
}
