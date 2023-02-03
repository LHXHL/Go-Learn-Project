package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Weapon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Yuan struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Weapon *Weapon `json:"weapon"`
}

var yuans []Yuan

func fakeData() {
	yuans = append(yuans, Yuan{
		ID:   1,
		Name: "烟绯",
		Weapon: &Weapon{
			ID:   1,
			Name: "流浪法章",
		},
	})
	yuans = append(yuans, Yuan{
		ID:   2,
		Name: "琴",
		Weapon: &Weapon{
			ID:   2,
			Name: "圣剑",
		},
	})
	yuans = append(yuans, Yuan{
		ID:   3,
		Name: "瑶瑶",
		Weapon: &Weapon{
			ID:   3,
			Name: "白缨枪",
		},
	})
}

func main() {
	fakeData()
	route := mux.NewRouter()

	route.HandleFunc("/api/roles", getAllRoles).Methods("GET")
	route.HandleFunc("/api/roles/{id}", getRolesById).Methods("GET")
	route.HandleFunc("/api/roles", createRole).Methods("POST")
	route.HandleFunc("/api/roles/{id}", updateRole).Methods("PUT")
	route.HandleFunc("/api/roles/{id}", deleteRole).Methods("DELETE")

	log.Println("Listening on port 8080", "http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", route))
}

func updateRole(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range yuans {
		id, _ := strconv.Atoi(params["id"])
		if id == item.ID {
			yuans = append(yuans[:index], yuans[index+1:]...)

			var yuan Yuan
			_ = json.NewDecoder(request.Body).Decode(&yuan)
			yuans = append(yuans, yuan)
			json.NewEncoder(writer).Encode(yuans)
			return
		}
	}
}

func deleteRole(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range yuans {
		id, _ := strconv.Atoi(params["id"])
		if id == item.ID {
			yuans = append(yuans[:index], yuans[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(yuans)
}

func createRole(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var yuan Yuan
	_ = json.NewDecoder(request.Body).Decode(&yuan)
	if yuan.Name == "" || yuan.ID == 0 {
		json.NewEncoder(writer).Encode(yuan)
		return
	}
	for index, element := range yuans {
		if element.ID == yuan.ID {
			yuans[index] = yuan
			json.NewEncoder(writer).Encode(yuans)
			return
		}
	}
	yuans = append(yuans, yuan)
	json.NewEncoder(writer).Encode(yuans)
}

func getRolesById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range yuans {
		id, _ := strconv.Atoi(params["id"])
		if id == item.ID {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func getAllRoles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(yuans)

}
