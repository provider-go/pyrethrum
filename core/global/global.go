package global

import (
	"github.com/provider-go/pkg/cache"
	"github.com/provider-go/pkg/cache/typecache"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/mysql"
	"github.com/provider-go/pkg/mysql/typemysql"
	"github.com/provider-go/pkg/smcc"
	"github.com/provider-go/pkg/smcc/typesmcc"
	"github.com/provider-go/pkg/types"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	DB           *gorm.DB
	RootDir      string
	PluginConfig types.PluginNeedInstance
)

func Initialize() {
	PluginConfig = types.PluginNeedInstance{
		Mysql: MysqlInstance(),
		Cache: CacheInstance(),
		SMCC:  SMCCInstance(),
	}
}

func MysqlInstance() *gorm.DB {
	cfg := typemysql.ConfigMysql{
		Dsn:          viper.GetString("mysql.dsn"),
		MaxIdleConns: viper.GetInt("mysql.maxIdleConns"),
		MaxOpenConns: viper.GetInt("mysql.maxOpenConns"),
	}
	mysqlInstance, err := mysql.NewMysql(cfg)
	if err != nil {
		logger.Error("MysqlInstance:", "step", "NewMysql", "err", err)
	}
	logger.Info("MySQL connection successful!")
	return mysqlInstance
}

func CacheInstance() cache.Cache {
	cfg := typecache.ConfigCache{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("mysql.db"),
	}
	cacheInstance, err := cache.NewCache("redis", cfg)
	if err != nil {
		logger.Error("CacheInstance:", "step", "NewCache", "err", err)
	}
	logger.Info("Cache connection successful!")
	return cacheInstance
}

func SMCCInstance() smcc.SMCC {
	cfg := typesmcc.ConfigSMCC{
		Endpoints:   []string{viper.GetString("etcd.addr")},
		DialTimeout: time.Second * viper.GetDuration("etcd.deadline"),
	}
	smccInstance, err := smcc.NewSMCC("etcd", cfg)
	if err != nil {
		logger.Error("CacheInstance:", "step", "NewCache", "err", err)
	}
	logger.Info("Cache connection successful!")
	// 向配置中心增加配置信息
	addConfig(smccInstance)

	return smccInstance
}

func addConfig(smccInstance smcc.SMCC) {
	// 增加ipfs信息到配置中心
	err := smccInstance.SetConfig("ipfs.addr", viper.GetString("ipfs.addr"))
	if err != nil {
		logger.Error("addConfig:", "step", "SetConfig", "err", err)
	}
}
