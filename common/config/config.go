package config

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

var loadRemoteConfigFlag = flag.Bool("remote", false, "是否从远程apollo读取配置")

var (
	Server   *ServerConfig
	COS      *COSConfig
	Database *DatabaseConfig
	JWT      *JWTConfig
	Redis    *RedisConfig
	Pulsar   *PulsarConfig
	MongoDB  *MongoDBConfig
)

type ServerConfig struct {
	RunMode            string
	HttpPort           string
	UserServiceAddr    string
	VideoServiceAddr   string
	CommentServiceAddr string
	MessageServiceAddr string
	Timeout            int
	EtcdAddress        string
	FeedCount          int64
	PprofSwitch        string
}

type COSConfig struct {
	VideoBucket string
	CoverBucket string
	SecretID    string
	SecretKey   string
}

type DatabaseConfig struct {
	DBType    string
	UserName  string
	Password  string
	Host      string
	DBName    string
	Charset   string
	ParseTime string
}

type RedisConfig struct {
	Address        string
	MaxIdle        int
	MaxActive      int
	ExpireTime     int
	MaxRandAddTime int
	BloomOpen      bool
}

type PulsarConfig struct {
	URL               string
	OperationTimeout  int
	ConnectionTimeout int
}

type MongoDBConfig struct {
	URL      string
	DataBase string
}

type JWTConfig struct {
	Secret  string
	Expires time.Duration
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		d.UserName,
		d.Password,
		d.Host,
		d.DBName,
		d.Charset,
		d.ParseTime,
	)
}

func InitConfig() {
	parseParam()

	vp := viper.New()

	if !*loadRemoteConfigFlag {
		workDirectory, err := os.Getwd()
		if err != nil {
			klog.Fatal(err)
		}
		sep := string(filepath.Separator)
		vp.AddConfigPath(workDirectory + sep + "config")
		for filepath.Base(workDirectory) != "runedance" {
			vp.AddConfigPath(workDirectory + sep + "config")
			workDirectory = filepath.Dir(workDirectory)
		}
		vp.AddConfigPath(workDirectory + sep + "config")
		vp.SetConfigName("config")
		vp.SetConfigType("yaml")
		if err := vp.ReadInConfig(); err != nil {
			klog.Fatal(err)
		}
	} else {
		//Read configuration from apollo configuration center
		c := &config.AppConfig{
			AppID:          "douyin",
			Cluster:        "dev",
			IP:             "http://127.0.0.1:8080",
			NamespaceName:  "application",
			IsBackupConfig: false,
		}
		client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
			return c, nil
		})
		klog.Info("Initializing Apollo configuration successfully")
		cache := client.GetConfigCache(c.NamespaceName)
		confValue, _ := cache.Get("conf.yaml")
		confString := fmt.Sprint(confValue)

		vp.SetConfigType("yaml")
		err := vp.ReadConfig(bytes.NewBuffer([]byte(confString)))
		if err != nil {
			fmt.Println(err)
		}
	}

	vp.UnmarshalKey("Server", &Server)
	vp.UnmarshalKey("Database", &Database)
	vp.UnmarshalKey("JWT", &JWT)
	vp.UnmarshalKey("COS", &COS)
	vp.UnmarshalKey("Redis", &Redis)
	vp.UnmarshalKey("Pulsar", &Pulsar)
	vp.UnmarshalKey("MongoDB", &MongoDB)
	JWT.Expires *= time.Hour
	Redis.ExpireTime *= 3600
	Redis.MaxRandAddTime *= 3600
}

func parseParam() {
	flag.Parse()
}
