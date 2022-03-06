package setting

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"time"
)

var (
	CfgMap map[interface{}]interface{}

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string
)

func init() {
	file, err := ioutil.ReadFile("conf/app.yaml")
	if err != nil {
		log.Fatalf("Fail to load 'conf/app.yaml': %v", err)
	}
	if yaml.Unmarshal(file, &CfgMap) != nil {
		log.Fatalf("error: %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = CfgMap["RUN_MODE"].(string)
}

func LoadServer() {
	server := CfgMap["server"].(map[string]interface{})

	HTTPPort = server["HTTP_PORT"].(int)
	ReadTimeout = time.Duration(server["READ_TIMEOUT"].(int)) * time.Second
	WriteTimeout = time.Duration(server["WRITE_TIMEOUT"].(int)) * time.Second
}

func LoadApp() {
	appCfg := CfgMap["app"].(map[string]interface{})

	JwtSecret = appCfg["JWT_SECRET"].(string)
}
