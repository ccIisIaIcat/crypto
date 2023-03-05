package global

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"gopkg.in/ini.v1"
)

func NeverStop() {
	for {
		time.Sleep(time.Hour)
	}
}

func GetConfig(conf_path string) Config {
	cfg, err := ini.Load(conf_path)
	if err != nil {
		panic(err)
	}
	Config_obj := Config{}
	Config_obj.MysqlInfo = make(map[string]ConfigMysql, 0)
	Config_obj.UserInfo = make(map[string]ConfigUser, 0)
	// 读取Mysql，目前Mysql类型Local
	temp_list := []string{"Local", "Local2", "Rm", "Rm1"}
	for i := 0; i < len(temp_list); i++ {
		temp_mysql := ConfigMysql{}
		temp_mysql.Host = cfg.Section("Mysql").Key("Host" + temp_list[i]).String()
		temp_mysql.Port = cfg.Section("Mysql").Key("Port" + temp_list[i]).String()
		temp_mysql.User = cfg.Section("Mysql").Key("User" + temp_list[i]).String()
		temp_mysql.Password = cfg.Section("Mysql").Key("Password" + temp_list[i]).String()
		Config_obj.MysqlInfo[temp_list[i]] = temp_mysql
	}
	temp_list = []string{"1", "Simulate"}
	for i := 0; i < len(temp_list); i++ {
		temp_user := ConfigUser{}
		temp_user.Apikey = cfg.Section("User").Key("Apikey" + temp_list[i]).String()
		temp_user.Secretkey = cfg.Section("User").Key("Secretkey" + temp_list[i]).String()
		temp_user.Passphrase = cfg.Section("User").Key("Passphrase" + temp_list[i]).String()
		Config_obj.UserInfo[temp_list[i]] = temp_user
	}
	return Config_obj
}

func ComputeHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
