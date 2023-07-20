package ditryWords

import (
	"fmt"
	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
	"gopkg.in/mgo.v2"
	"runedance/common/config"
	"time"
)

var filterManage *filter.DirtyManager

func Init() {
	mogoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.MongoDB.URL},
		Timeout:  time.Second * 60,
		Database: config.MongoDB.DataBase,
	}
	session, err := mgo.DialWithInfo(mogoDBDialInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	mgoStore, err := store.NewMongoStore(store.MongoConfig{
		Session: session,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	filterManage = filter.NewDirtyManager(mgoStore)
}
