package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

// 存储所有连接的客户端
var clients = make(map[int64]*websocket.Conn)
var mu sync.Mutex

// 处理WebSocket连接
func WebSocketControllerWebSocket(c *gin.Context) {
	userIDStr := c.Query("user_id")
	toIDStr := c.Query("to_id")

	// 将user_id转换为int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return
	}
	// 将to_id转换为int64
	toID, err := strconv.ParseInt(toIDStr, 10, 64)
	if err != nil {
		toID = -1
	}

	// 升级HTTP连接到WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	// 处理连接关闭错误
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Println("Error closing connection:", closeErr)
		}
	}()

	// 添加连接到客户端列表
	mu.Lock()
	clients[userID] = conn
	mu.Unlock()

	// 连接断开时移除客户端
	defer func() {
		mu.Lock()
		delete(clients, userID) // 移除连接
		mu.Unlock()
	}()

	// 处理WebSocket消息
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received message: %s\n", msg)

		// 广播消息给所有连接的客户端
		mu.Lock()
		if toID != -1 {
			for userid, client := range clients {
				if userid == toID {
					err = client.WriteMessage(websocket.TextMessage, msg)
					if err != nil {
						log.Println("Error writing message:", err)
						// 处理关闭客户端连接错误
						if closeErr := client.Close(); closeErr != nil {
							log.Println("Error closing client connection:", closeErr)
						}
						delete(clients, userid) // 移除无法写入的连接
					}
				}
			}
		} else {
			for userid, client := range clients {
				err = client.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Println("Error writing message:", err)
					// 处理关闭客户端连接错误
					if closeErr := client.Close(); closeErr != nil {
						log.Println("Error closing client connection:", closeErr)
					}
					delete(clients, userid) // 移除无法写入的连接
				}
			}
		}
		mu.Unlock()
	}
}
