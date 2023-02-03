package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {

	log.Println(`------- Start Init Configs -------`)

	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	log.Println(`------- Inited Configs -------`)

}
