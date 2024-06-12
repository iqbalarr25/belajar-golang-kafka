package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var ctx context.Context
var db *gorm.DB
var cache *redis.Client

func InitDatabase() {
	var err error
	conf := GetDatabase()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Host, conf.Username, conf.Password, conf.Database, conf.Port, conf.Timezone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NowFunc: func() time.Time {
			jakarta, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(jakarta)
		},
	})
	if err != nil {
		panic(err)
	}
}

func CreateDBConnection() *gorm.DB {
	return db
}

func InitCache(context context.Context) {
	ctx = context
	conf := GetCache()
	cache = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.Database,
	})
}

func SaveCache(key string, value interface{}) error {
	err := cache.Set(ctx, key, value, 60*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func ReadCache(key string) ([]byte, error) {
	val, err := cache.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}
