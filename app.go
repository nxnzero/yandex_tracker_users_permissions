package main

import (
	"fmt"
	"log"
	"os"
	"rightmap/queues"
	"rightmap/rights"
	"rightmap/users"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	users := users.AllUsers()

	for _, user := range users {
		// fmt.Printf("Пользователь: %v \n", user.Display)
		queues := queues.AllQueues()
    
		for _, v := range queues {
			rights := rights.QueueuByUser(v.Key, user.Login)
            // Добавить проверки, если у вас есть удаленнные очереди.
      
			for _, userRight := range rights {
				line := fmt.Sprintf("Очередь: %v\t| Права доступа: %v\t| Пользователь: %v\n", v.Key, userRight, user.Display)
				_, err := file.WriteString(line)
				if err != nil {
					log.Fatal(err)
					continue
				}
			}
		}
	}
}
