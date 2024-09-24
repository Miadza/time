package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)



// Функция для отправки уведомления
func sendNotification(title, message string) {
    err := beeep.Notify(title, message, "")
    if err != nil {
        fmt.Println("Ошибка при отправке уведомления:", err)
    }
}

func main() {
    targetTime := "09:02" // Время для уведомления в формате HH:MM

    for {
        // Получаем текущее время в формате HH:MM
        currentTime := time.Now().Format("15:04")

        // Проверяем, совпадает ли текущее время с целевым
        if currentTime == targetTime {
            sendNotification("Напоминание", "Сейчас 09:02, пора действовать!")
            time.Sleep(60 * time.Second) // Ждём минуту, чтобы избежать повторов
        }

        time.Sleep(10 * time.Second) // Проверяем каждые 10 секунд
    }
}