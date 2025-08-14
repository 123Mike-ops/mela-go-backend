package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct{
	Server struct{
		Port int `yaml:"port"`
	}`Yaml:"server"`

    Database struct{
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		DBName string `yaml:"dbname"`
	}`yaml:database`
}

func LoadConfig(path string )(*Config,error){
	data,err:=os.ReadFile(path);
	if err!=nil {
		return nil,fmt.Errorf("failed to load the conf file: %w",err)
	}
	var cfg Config
	if err:=yaml.Unmarshal(data,&cfg);err!=nil{
		return nil,fmt.Errorf("failed to parse the config: %w",err)
	}
	return &cfg,nil

}