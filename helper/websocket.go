package websocket

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
    "sync"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var mutex = &sync.Mutex{}

type Message struct {
    Event string `json:"event"`
    Data  string `json:"data"`
}

func HandleConnections(c *gin.Context) {
    ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Failed to upgrade to websocket: %v", err)
        return
    }
    defer ws.Close()

    mutex.Lock()
    clients[ws] = true
    mutex.Unlock()

    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("Unexpected close error: %v", err)
            } else {
                log.Printf("WebSocket closed: %v", err)
            }
            break
        }
    }

    mutex.Lock()
    delete(clients, ws)
    mutex.Unlock()
}

func HandleMessages() {
    for {
        msg := <-broadcast
        mutex.Lock()
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("Error writing JSON: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
        mutex.Unlock()
    }
}

func NotifyClients(event, data string) {
    broadcast <- Message{Event: event, Data: data}
}
