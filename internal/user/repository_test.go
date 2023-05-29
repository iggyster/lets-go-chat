package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestRepo Repo = &InMemoryRepo{}

func TestInMemoryRepo_FindByUsername(t *testing.T) {
	usr := New("test", "pass")
	TestRepo.Save(usr)

	_, err := Repository.FindByUsername("non")

	assert.Error(t, err)
}

func TestInMemoryRepo_FindByToken(t *testing.T) {
	usr := New("test", "pass")
	usr.SetToken("token")
	TestRepo.Save(usr)

	assert.NotNil(t, TestRepo.FindByToken("token"))
	assert.Nil(t, TestRepo.FindByToken("non"))
}

func TestInMemoryRepo_IsExists(t *testing.T) {
	usr := New("test", "pass")
	TestRepo.Save(usr)

	assert.True(t, TestRepo.IsExists("test"))
	assert.False(t, TestRepo.IsExists("non"))
}

func TestInMemoryRepo_Save(t *testing.T) {
	usr := New("test", "pass")

	TestRepo.Save(usr)

	expected, err := TestRepo.FindByUsername("test")
	assert.NoError(t, err)
	assert.Same(t, expected, usr)
}
