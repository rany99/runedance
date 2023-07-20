package dal

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"runedance/common/config"
)

var DB *gorm.DB

// Init DB
func Init() {
	config.InitConfig()
	var err error
	DB, err = gorm.Open(mysql.Open(config.Database.DSN()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatal(err)
	}
}
