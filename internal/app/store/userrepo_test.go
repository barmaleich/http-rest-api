package store_test

import (
	"github.com/barmaleich/http-rest-api/internal/app/model"
	"github.com/barmaleich/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	databaseURL = "host=localhost dbname=restapi_test user=postgres password=postgres sslmode=disable"
)

func TestUser_Create(t *testing.T)  {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail_NotExistingUser(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@mail.ru"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)


}

func TestUserRepo_FindByEmail_ExistingUser(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	u := model.TestUser(t)
	u.Email = email

	_, _ = s.User().Create(u)
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
