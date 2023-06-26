package handler

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/iggyster/lets-go-chat/internal/chat"
	"github.com/iggyster/lets-go-chat/internal/user"
)

func TestNewChat(t *testing.T) {
	var h Any = NewChatHandler(user.NewRepo(), chat.NewHub())

	_, ok := h.(*ChatHandler)
	if !ok {
		t.Errorf("faile to create new chat handler")
	}
}

func TestChat_ServeHTTP(t *testing.T) {
	repo := user.NewRepo()
	handler := NewChatHandler(repo, chat.NewHub())

	server := httptest.NewServer(handler)
	defer server.Close()

	usr := user.New("test", "pass")
	usr.SetToken("token")
	repo.Save(usr)

	url := "ws" + strings.TrimPrefix(server.URL, "http")
	ws, _, err := websocket.DefaultDialer.Dial(url+"?token=token", nil)
	defer ws.Close()

	if err != nil {
		t.Errorf("%v", err)
	}

	for i := 0; i < 10; i++ {
		msg := "ping"
		if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			t.Errorf("%v", err)
		}

		_, got, err := ws.ReadMessage()
		if err != nil {
			t.Errorf("%v", err)
		}

		if string(got) != msg {
			t.Errorf("fail to echo the message")
		}
	}
}
