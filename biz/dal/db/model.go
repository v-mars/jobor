package db

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func Paginate(page uint64, pageSize uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (int(page) - 1) * int(pageSize)
		return db.Offset(offset).Limit(int(pageSize))
	}
}

func WhereWithScope(key string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key+" = ?", value)
	}
}

func OrWithScope(key string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Or(key+" = ?", value)
	}
}

func GetOneWithScope(ctx context.Context, Db *gorm.DB, table string, out interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	err := Db.WithContext(ctx).Table(table).Scopes(scopes...).First(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func GetOneById(ctx context.Context, Db *gorm.DB, table string, id int64, out interface{}) error {
	err := Db.WithContext(ctx).Table(table).Where("id = ?", id).First(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func EditOneById(ctx context.Context, Db *gorm.DB, table string, id int64, out interface{}) error {
	err := Db.WithContext(ctx).Table(table).Where("id = ?", id).Updates(out).Error
	if err != nil {
		return err
	}
	return nil
}

func Create(ctx context.Context, Db *gorm.DB, table string, out interface{}) error {
	err := Db.WithContext(ctx).Table(table).Create(out).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteOneById(ctx context.Context, Db *gorm.DB, table string, id int64, model interface{}) error {
	return Db.WithContext(ctx).Table(table).Where("id = ?", id).Delete(&model).Error
}

func GetCountWithScope(ctx context.Context, Db *gorm.DB, table string, out interface{}, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var i int64
	err := Db.WithContext(ctx).Table(table).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return i, err
	}
	return i, nil
}

func Exist(ctx context.Context, Db *gorm.DB, table string, scopes ...func(*gorm.DB) *gorm.DB) bool {
	var count int64
	err := Db.WithContext(ctx).Table(table).Scopes(scopes...).Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}

//func SelectScopes() func(Db *gorm.DB) *gorm.DB {
//	return func(Db *gorm.DB) *gorm.DB {
//		return Db.Where("distinct *")
//	}
//}

func ExistWhere(ctx context.Context, Db *gorm.DB, table string, where interface{}, args ...interface{}) bool {
	var whereScope = func(Db *gorm.DB) *gorm.DB {
		return Db.Where(where, args...)
	}
	return Exist(ctx, Db, table, whereScope)
}
