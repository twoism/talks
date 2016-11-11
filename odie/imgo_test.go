import (
	"github.com/lyft/imgo/mock"
	"github.com/lyft/odie/client/mongo"
)

func TestRepo_GetById(t *testing.T) {
	u := test_repo.UserModel{
		Id:   1,
		Name: "test",
	}

	sess := &mock.Session{
		Results: &u,
	}

	repo := New(mongo.New(sess))

	resp, err := repo.GetByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, &u, resp)
}