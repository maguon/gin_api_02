package global

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"

	config "gin_api_02/config"
)

var (
	SYS_PORT      int
	SYS_DB        *gorm.DB
	SYS_DBList    map[string]*gorm.DB
	SYS_CONFIG    config.Server
	SYS_REDIS     *redis.Client
	SYS_MONGO     *mongo.Database
	SYS_MONGO_CTX context.Context
	SYS_VP        *viper.Viper
	SYS_LOG       *zap.Logger
)
