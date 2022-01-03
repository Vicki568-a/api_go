package main

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"golang.org/x/net/websocket"
	"sync"
	"time"
)

type ChatRoom struct {
	clients    map[string]Client
	clientsMtx sync.Mutex
	queue      chan Message
}

// Инициализация чата
func (chatRoom *ChatRoom) Init() {
	chatRoom.queue = make(chan Message, 5)
	chatRoom.clients = make(map[string]Client)

	go func() {
		for {
			chatRoom.BroadCast()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// Регистрируем нового клиента
// возвращает указатель на клиента или Nil, если имя уже занят
func (chatRoom *ChatRoom) Join(name string, conn *websocket.Conn) *Client {
	defer chatRoom.clientsMtx.Unlock()

	chatRoom.clientsMtx.Lock()
	if _, exists := chatRoom.clients[name]; exists && len(name) >= 3 {
		return nil
	}
	client := Client{
		name:     name,
		conn:     conn,
		chatRoom: chatRoom,
	}
	chatRoom.clients[name] = client

	lastActive, _ := GetUserLastActive(name)

	client.Send([]Message{Message{uuid.NewV4(), "client-handshake", "", lastActive, ""}})

	messages := GetMessagesForUser(client.name, time.Now())

	AddUserOrUpdateLastActive(name)

	client.Send(messages)

	chatRoom.AddMessage(Message{uuid.NewV4(), "system-message", "", time.Now(), name + " has joined the chat."})
	return &client
}

// Выход из чата
func (chatRoom *ChatRoom) Leave(name string) {
	chatRoom.clientsMtx.Lock()
	delete(chatRoom.clients, name)
	AddUserOrUpdateLastActive(name)
	chatRoom.clientsMtx.Unlock()
	chatRoom.AddMessage(Message{uuid.NewV4(), "system-message", "", time.Now(), name + " has left the chat."})
}

// Добавление сообщения в очередь
func (chatRoom *ChatRoom) AddMessage(message Message) {
	chatRoom.queue <- message
	if message.Name != "" {
		StoreMessage(message)
	}
}

// Трансляция всех сообщений в очереди одним блоком
func (chatRoom *ChatRoom) BroadCast() {

	messages := make([]Message, 0)

infLoop:
	for {
		select {
		case message := <-chatRoom.queue:
			messages = append(messages, message)
		default:
			break infLoop
		}
	}
	if len(messages) > 0 {
		for _, client := range chatRoom.clients {
			client.Send(messages)
		}
	}
}
