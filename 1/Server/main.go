package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Порт, который будет прослушивать сервер
	port := 12345

	// Создаем UDP адрес сервера
	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Ошибка при разрешении адреса сервера:", err)
		os.Exit(1)
	}

	// Создаем UDP соединение
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Ошибка при создании UDP соединения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Сервер UDP слушает на порту %d...\n", port)

	// Буфер для приема входящих сообщений
	buffer := make([]byte, 1024)

	for {
		// Читаем входящее сообщение
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			continue
		}

		fmt.Printf("Получено сообщение от %s: %s\n", addr.String(), string(buffer[:n]))

		// Отправляем ответ клиенту
		response := []byte("Привет, клиент!")
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("Ошибка при отправке ответа:", err)
		}
	}
}

// go run main.go

// go mod init mymodule
// go build -o myapp.exe
// ./myapp
// Ctrl+C

//go build main.go
// go build -o myapp.exe main.go
// ./myapp
// Ctrl+C
