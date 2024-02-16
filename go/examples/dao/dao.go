package dao

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

type Dao[T Model] interface {
	Create(model T) error
	Get(id string) (T, error)
	Delete(id string) error
}

type BoltDao[T Model] struct {
	Db          *bbolt.DB
	BucketName  string
	Serialize   func(model T) ([]byte, error)
	Deserialize func(bytes []byte) (T, error)
}

func (d BoltDao[T]) Create(model T) error {
	return d.Db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(d.BucketName))
		if err != nil {
			return errors.Wrapf(err, `failed to create bucket '%s'`, d.BucketName)
		}
		id := uuid.NewString()
		data, err := d.Serialize(model)
		if err != nil {
			return errors.Wrapf(err, `failed to serialize model '%s'`, d.BucketName)
		}
		err = bucket.Put([]byte(id), data)
		if err != nil {
			return errors.Wrapf(
				err,
				`failed to put model '%s' with ID='%s' into bucket`,
				d.BucketName,
				id,
			)
		}
		model.SetId(id)
		return nil
	})
}

func (d BoltDao[T]) Get(id string) (T, error) {
	var m T
	err := d.Db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(d.BucketName))
		data := bucket.Get([]byte(id))
		if data == nil {
			return errors.Errorf(
				`model '%s' with ID='%s' does not exist`,
				d.BucketName,
				id,
			)
		}
		var err error
		m, err = d.Deserialize(data)
		if err != nil {
			return errors.Wrapf(
				err,
				`failed to serialize model '%s' with ID='%s'`,
				d.BucketName,
				id,
			)
		}
		return nil
	})
	return m, err
}

func (d BoltDao[T]) Delete(id string) error {
	return d.Db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(d.BucketName))
		err := bucket.Delete([]byte(id))
		if err != nil {
			return errors.Wrapf(
				err,
				`failed to delete model '%s' with ID='%s'`,
				d.BucketName,
				id,
			)
		}
		return nil
	})
}
