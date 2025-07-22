package users

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	BaseURL = "https://api.tracker.yandex.net/v3/"
)

func AllUsers() []User {
	c := http.Client{}

	request, err := http.NewRequest("GET", BaseURL+"users?perPage=10000", nil)
	if err != nil {
		log.Fatalf("Ошибка при формировании запроса: %v", err)
	}

	request.Header.Set("Authorization", "OAuth "+os.Getenv("TOKEN"))
	request.Header.Set("X-Org-ID", os.Getenv("ORG_ID"))

	response, err := c.Do(request)

	if response.StatusCode != 200 {
		log.Fatalf("Ошибка при вызове API пользователей..\nСтатус: %v", response.StatusCode)
	}

	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}

	defer response.Body.Close()

	var usersList []User

	json.NewDecoder(response.Body).Decode(&usersList)

	/*
		for _, user := range usersList {
			fmt.Printf("%v\n", user.Login) //user.Display)
		} */

	return usersList
}

type User struct {
	Login       string `json:"login"`
	Display     string `json:"display"`
	TrackerUid  int    `json:"trackerUid"`
	PassportUid int    `json:"passportUid"`
}
