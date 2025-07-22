package queues

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	BaseURL = "https://api.tracker.yandex.net/v3/"
)

func AllQueues() []Queue {
	c := http.Client{}

	// Формируем запрос
	request, err := http.NewRequest("GET", BaseURL+"queues?perPage=10000", nil)
	if err != nil {
		log.Fatalf("Ошибка при формировании запроса: %v", err)
	}

	// Записываем заголовки
	request.Header.Set("Authorization", "OAuth "+os.Getenv("TOKEN"))
	request.Header.Set("X-Org-ID", os.Getenv("ORG_ID"))

	// Отправляем запрос
	response, err := c.Do(request)

	// Обрабатываем ошибки
	if response.StatusCode != 200 {
		log.Fatalf("Ошибка при вызове API в правах очередей.\nСтатус: %v", response.StatusCode)
	}

	if err != nil {
		log.Fatalf("Ошибка при отправке запроса запроса: %v", err)
	}

	// Закрываем тело ответа
	defer response.Body.Close()

	var queueList []Queue

	// Декодируем ответ
	json.NewDecoder(response.Body).Decode(&queueList)

	/*
		for _, user := range usersList {
			fmt.Printf("%v\n", user.Login) //user.Display)
		} */

	return queueList
}

type Queue struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
