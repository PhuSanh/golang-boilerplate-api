package setting

import (
	"github.com/labstack/gommon/log"
	"golang-boilerplate/config"
	"golang-boilerplate/database"
	"golang-boilerplate/database/mysql"
)

var cfg = config.GetConfig()

func InitMysql() (err error) {
	database.MysqlConn, err = mysql.NewConn(cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.DatabaseName, cfg.Mysql.MaxIdleConn, cfg.Mysql.MaxOpenConn)
	if err != nil {
		log.Print("Connect database fail")
	} else {
		log.Print("Connected database success")
	}
	return
}