package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/iggyster/lets-go-chat/internal/handler"
	"io"
	"log"
	"net/http"
	"testing"

	ws "github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/iggyster/lets-go-chat/internal/router"
	"github.com/stretchr/testify/assert"
)

func TestStartChat(t *testing.T) {
	app := fiber.New()
	router.Init(app)

	go func() {
		log.Fatal(app.Listen(":8080"))
	}()
	jsonData := []byte(`{"userName": "test", "password": "123qweasd"}`)
	bodyReader := bytes.NewReader(jsonData)

	resp, err := http.Post("http://localhost:8080/user", "application/json", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	bodyReader = bytes.NewReader(jsonData)

	resp, err = http.Post("http://localhost:8080/user/login", "application/json", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	url := &handler.LoginResource{}
	err = json.Unmarshal(respBody, &url)

	conn, _, err := ws.DefaultDialer.Dial(url.Url, nil)
	assert.Nil(t, err)

	for i := 0; i < 10; i++ {
		err := conn.WriteMessage(websocket.TextMessage, []byte("hello"))
		assert.Nil(t, err)
		_, msg, err := conn.ReadMessage()
		assert.Nil(t, err)
		assert.Equal(t, []byte("hello"), msg)
	}

	err = app.Shutdown()
	assert.Nil(t, err)
}
