package mexcws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// MEXCWebSocket is a WebSocket client for the MEXC exchange
type MEXCWebSocket struct {
	url       string
	sendMutex *sync.Mutex

	Subs          *Subscribes
	Conn          *websocket.Conn
	ErrorListener OnError
}

// NewMEXCWebSocket returns a new MEXCWebSocket instance
func NewMEXCWebSocket(errorListener OnError) *MEXCWebSocket {
	return &MEXCWebSocket{
		url:           "wss://wbs.mexc.com/ws",
		ErrorListener: errorListener,
	}
}

// Send sends a message to the server
func (m *MEXCWebSocket) Send(message *WsReq) error {
	if m.Conn == nil {
		return fmt.Errorf("no available connection")
	}

	m.sendMutex.Lock()
	defer m.sendMutex.Unlock()

	return m.Conn.WriteJSON(message)
}

// Connect establishes a WebSocket connection to the MEXC exchange
func (m *MEXCWebSocket) Connect(ctx context.Context) error {
	var err error

	m.Conn, _, err = websocket.DefaultDialer.Dial(m.url, nil)
	if err != nil {
		return err
	}

	m.sendMutex = &sync.Mutex{}
	m.Subs = NewSubs()

	go m.keepAlive(ctx)
	go m.readLoop(ctx)

	return nil
}

// Disconnect closes the WebSocket connection
func (m *MEXCWebSocket) Disconnect() {
	err := m.Conn.Close()

	if err != nil {
		return
	}
}

// keepAlive sends a ping message to the server every 30 seconds to keep the connection alive
func (m *MEXCWebSocket) keepAlive(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := m.Send(&WsReq{Method: "PING"})
			if err != nil {
				m.ErrorListener(err)
			}
		}
	}
}

// readLoop read messages and resolve handlers
func (m *MEXCWebSocket) readLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.handleLoop()
		}
	}
}

func (m *MEXCWebSocket) handleLoop() {
	if m.Conn == nil {
		return
	}

	_, buf, err := m.Conn.ReadMessage()
	if err != nil {
		m.ErrorListener(err)

		return
	}

	message := string(buf)

	var update map[string]interface{}
	err = json.Unmarshal(buf, &update)
	if err != nil {
		m.ErrorListener(err)

		return
	}

	if update["msg"] == "PONG" {
		return
	}

	listener := m.getListener(update)
	if listener != nil {
		listener(message)

		return
	}

	log.Println(fmt.Sprintf("Unhandled: %v", update))
}

func (m *MEXCWebSocket) getListener(argJson interface{}) OnReceive {
	mapData := argJson.(map[string]interface{})

	v, _ := m.Subs.Load(fmt.Sprintf("%s", mapData["c"]))

	return v
}
