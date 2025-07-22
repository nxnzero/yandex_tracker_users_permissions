package rights

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Структура для хранения данных о правах доступа пользователя
type UserPermissions struct {
	User struct {
		ID      string `json:"id"`
		Display string `json:"display"`
	} `json:"user"`
	Permissions map[string]interface{} `json:"permissions"` // Права представлены как объект
}

func QueueuByUser(queue, user string) []string {
	c := http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Формируем URL запроса
	URL := fmt.Sprintf("https://api.tracker.yandex.net/v3/queues/%v/permissions/users/%v", queue, user)

	request, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		log.Fatalf("Ошибка при формировании запроса: %v", err)
	}

	// Устанавливаем заголовки
    // TOKEN и ORG_ID - переменные окружения. Запиши свои, либо замени на токен и орг_айди
	request.Header.Set("Authorization", "OAuth "+os.Getenv("TOKEN")) 
	request.Header.Set("X-Org-ID", os.Getenv("ORG_ID")) // либо заменить на X-Org-Cloud-ID

	// Отправляем запрос
	response, err := c.Do(request)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)

	}

	defer response.Body.Close()

	// Проверяем статус ответа
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка при вызове API по правам в очередях. Статус: %v", response.StatusCode)
		log.Fatal(response.Body)
	}

	// Декодируем ответ
	var permissions UserPermissions
	err = json.NewDecoder(response.Body).Decode(&permissions)
	if err != nil {
		log.Fatalf("Ошибка при декодировании ответа: %v", err)
	}

	// Если нужно получить просто список названий прав
	var permissionList []string
	for perm := range permissions.Permissions {
		permissionList = append(permissionList, perm)
	}

	// Выводим результат
	/*
		fmt.Printf("Пользователь: %s (%s)\n", permissions.User.Display, permissions.User.ID)
		fmt.Println("Права доступа:") */

	// Перебираем все права в объекте permissions
	/*
		for perm, value := range permissions.Permissions {
			fmt.Printf("- %s: %v\n", perm, value)
		} */

	return permissionList
}
