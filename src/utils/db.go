package utils

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type ServerInfo struct {
	Host      string
	HostLocal string
	Port      string
	User      string
	Password  string
	Dbname    string
}

func InitXorm() (*xorm.Engine, error) {

	s := ServerInfo{
		Host:      viper.GetString(`db.host`),
		HostLocal: viper.GetString(`db.hostLocal`),
		Port:      viper.GetString(`db.port`),
		User:      viper.GetString(`db.user`),
		Password:  viper.GetString(`db.password`),
		Dbname:    viper.GetString(`db.dbname`),
	}
	mysqlFormat := `%s:%s@tcp(%s:%s)/%s`
	mysqlInfo := ``
	hostName, _ := os.Hostname()
	if strings.Contains(hostName, `local`) {
		mysqlInfo = fmt.Sprintf(mysqlFormat,
			s.User, s.Password, s.HostLocal, s.Port, s.Dbname)
	} else {
		mysqlInfo = fmt.Sprintf(mysqlFormat,
			s.User, s.Password, s.Host, s.Port, s.Dbname)
	}

	e, err := xorm.NewEngine("mysql", mysqlInfo)
	if err != nil {
		return nil, err
	}

	e.Logger().SetLevel(log.LOG_INFO)
	e.ShowSQL(true)

	return e, err
}
