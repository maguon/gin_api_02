package initialize

import (
	"context"
	"gin_api_02/global"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func Mongo() {
	mongoCfg := global.SYS_CONFIG.Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+mongoCfg.Host+":"+mongoCfg.Port))
	if err != nil {
		global.SYS_LOG.Error("mongodb connect failed, err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("mongodb connect success")
		global.SYS_MONGO = client.Database(mongoCfg.DB)
		global.SYS_MONGO_CTX = ctx
	}
}
