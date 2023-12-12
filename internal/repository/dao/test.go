package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	ErrTestNotFound = gorm.ErrRecordNotFound
)

type TestDAO interface {
	Create(ctx context.Context, userID uint64, answers []byte, signature string, signedAt int64) error
	Find(ctx context.Context, userID uint64, signature string) (Test, error)
}

type GORMTestDAO struct {
	db *gorm.DB
}

func NewTestDAO(db *gorm.DB) TestDAO {
	return &GORMTestDAO{
		db: db,
	}
}

type Test struct {
	ID     uint64 `gorm:"primaryKey,autoIncrement"`
	UserID uint64
	QAndA  []byte

	Signature string `gorm:"unique"`
	SignedAt  int64

	CreatedAt int64
	UpdatedAt int64
}

func (dao *GORMTestDAO) Create(ctx context.Context, userID uint64, answers []byte, signature string, signedAt int64) error {
	now := time.Now().Unix()
	t := &Test{
		UserID:    userID,
		QAndA:     answers,
		Signature: signature,
		SignedAt:  signedAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := dao.db.WithContext(ctx).Create(&t).Error
	return err
}

func (dao *GORMTestDAO) Find(ctx context.Context, userID uint64, signature string) (Test, error) {
	var t Test
	err := dao.db.WithContext(ctx).Where("user_id = ? AND signature = ?", userID, signature).First(&t).Error

	return t, err
}
