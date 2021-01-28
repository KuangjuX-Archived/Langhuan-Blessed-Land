package WebSocketService


import (
	"bytes"
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/json"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/auth"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/DataBases/redis"
)


const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// 用户名称
	username []byte
	// 房间号
	roomID []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump(conn redis.RedisConn) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		message = []byte(string(c.username) + ": " + string(message))
		
		//write to redis
		key := c.roomID
		_, err = conn.Do("LPUSH", key, message)
		if err != nil{
			fmt.Printf("error: %s\n", err)
		}

		// build special string to broadcast
		message = []byte(string(c.roomID) + "&?!*" + string(message))		
		c.hub.broadcast <- []byte(message)
	}

}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			// 使用“&”分割获取房间号
			// 聊天内容不得包含&字符
			// msg[0]为房间号 msg[1]为打印内容
			// msg := strings.Split(string(message), "&")
			// if msg[0] == string(c.hub.roomID[c]) {
			// 	w.Write([]byte(msg[1]))
			// }
			w.Write(message)
			// Add queued chat messages to the current websocket message.
			// n := len(c.send)
			// for i := 0; i < n; i++ {
			// 	if msg[0] == string(c.hub.roomID[c]) {
			// 		w.Write(newline)
			// 		w.Write(<-c.send)
			// 	}
			// }
			if err := w.Close(); err != nil {
				log.Printf("error: %v\n", err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		
		}
	}
	
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, c *gin.Context) {


	res, err := auth.GetUserByToken(c)
	if err != nil{
		json.JsonError(c, err)
		return
	}
	userInfo, _ := res.(Models.User)

	userName := userInfo.Username
	roomID := c.Query("room_id")


	// Get Redis Connection
	pool := redis.RedisPool
	redisConn := pool.Get()
	// defer redisConn.Close()
	

	

	// Change Http to WebSocket
	var upgrader = websocket.Upgrader{
		// Solve Cors
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("user: %s\n", userName)
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), username: []byte(userName), roomID: []byte(roomID)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump(redisConn)
}