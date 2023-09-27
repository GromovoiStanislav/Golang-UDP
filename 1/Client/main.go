package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// IP и порт сервера, на который мы хотим отправить сообщение
	serverAddr := "127.0.0.1:12345"

	// Создаем UDP адрес сервера
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Ошибка при разрешении адреса сервера:", err)
		os.Exit(1)
	}

	// Создаем UDP соединение
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Ошибка при создании UDP соединения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Сообщение, которое мы хотим отправить
	message := []byte("Привет, сервер!")

	// Отправляем сообщение серверу
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения:", err)
		os.Exit(1)
	}

	// Буфер для приема ответа от сервера
	buffer := make([]byte, 1024)

	// Читаем ответ от сервера
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		os.Exit(1)
	}

	// Выводим ответ на экран
	fmt.Println("Ответ от сервера:", string(buffer[:n]))
}

//go run main.go
