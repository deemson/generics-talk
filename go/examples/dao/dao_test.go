package dao_test

import (
	"encoding/json"
	"fmt"
	"github.com/deemson/generics-talk/go/examples/dao"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
	"os"
	"path"
	"testing"
)

type UserModel struct {
	id        string
	FirstName string
	LastName  string
}

func (m *UserModel) Id() string {
	return m.id
}

func (m *UserModel) SetId(id string) {
	m.id = id
}

func TestUserDao(t *testing.T) {
	dir := path.Join(os.TempDir(), "generics-talk", t.Name())
	err := os.MkdirAll(dir, 0755)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, os.RemoveAll(dir))
	}()
	dbPath := path.Join(dir, "dao.db")
	db, err := bbolt.Open(dbPath, 0600, nil)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, db.Close())
	}()

	d := dao.BoltDao[*UserModel]{
		Db:         db,
		BucketName: "users",
		Serialize: func(model *UserModel) ([]byte, error) {
			return json.Marshal(model)
		},
		Deserialize: func(bytes []byte) (*UserModel, error) {
			var u UserModel
			err := json.Unmarshal(bytes, &u)
			if err != nil {
				return nil, err
			}
			return &u, nil
		},
	}

	u := &UserModel{
		FirstName: "Joe",
		LastName:  "Doe",
	}
	assert.NoError(t, d.Create(u))
	assert.NotEqual(t, "", u.Id())

	u2, err := d.Get(u.Id())
	assert.NoError(t, err)
	assert.Equal(t, u2.FirstName, "Joe")
	assert.Equal(t, u2.LastName, "Doe")

	err = d.Delete(u.Id())
	assert.NoError(t, err)
	_, err = d.Get(u.Id())
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf(`model 'users' with ID='%s' does not exist`, u.Id()), err.Error())
}
