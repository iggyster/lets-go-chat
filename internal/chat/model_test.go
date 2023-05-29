package chat

import (
	"testing"

	"github.com/iggyster/lets-go-chat/internal/user"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()

	assert.Same(t, c, New())
	assert.IsType(t, &Chat{}, c)
	assert.Implements(t, (*UserChat)(nil), c)
}

func TestChat_GetUsers(t *testing.T) {
	usr := user.New("test", "pass")

	c := New()

	c.AddUser(usr)

	assert.Len(t, c.GetUsers(), 1)
	
	c.DisconnectUser(usr)
}

func TestChat_DisconnectUser(t *testing.T) {
	usr := user.New("test", "pass")

	c := New()
	c.AddUser(usr)

	c.DisconnectUser(usr)

	assert.Len(t, c.GetUsers(), 0)
}

func TestChat_IsUserActivated(t *testing.T) {
	usr := user.New("test", "pass")

	c := New()
	c.AddUser(usr)

	assert.True(t, c.IsUserActivated(usr.Id))
	
	c.DisconnectUser(usr)
}
