package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

type HTTP struct {
	Port string `yaml:"port"`
}

type YamlConfig struct {
	HTTP HTTP `yaml:"app"`
}

var iCnt int = 0
var yamlConfig YamlConfig
var configFIle string

func GetConfig() {
	viperConfig := viper.New()
	// 设置配置文件目录(可以设置多个,优先级根据添加顺序来)
	dir, file := path.Split(configFIle)
	viperConfig.AddConfigPath(dir)
	// 设置配置文件名，没有后缀
	viperConfig.SetConfigName(file)
	// 设置读取文件格式为: yaml
	viperConfig.SetConfigType("yaml")

	// 读取解析
	if err := viperConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("配置文件未找到！%v\n", err)
			return
		} else {
			fmt.Printf("找到配置文件,但是解析错误,%v\n", err)
			return
		}
	}
	// 映射到结构体

	// 使用Unmarshal方法自动映射到struct上
	if err := viperConfig.Unmarshal(&yamlConfig); err != nil {
		fmt.Printf("配置映射错误,%v\n", err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	iCnt++
        // 编译打包测试版本修改下面内容
	str := "Hello world ! friend01(" + strconv.Itoa(iCnt) + ")"
	io.WriteString(w, str)
	fmt.Println(str)
}

func init() {
	rootCmd.Flags().StringVarP(&configFIle, "config", "c", "./config.yml", "set config file")
}

var rootCmd = &cobra.Command{
	Use:     `helloserver`,
	Short:   `helloserver`,
	Example: `helloserver`,
	Run: func(cmd *cobra.Command, args []string) {
		GetConfig()
		fmt.Printf("YamlConfig: %#v", yamlConfig)
		ht := http.HandlerFunc(helloHandler)
		if ht != nil {
			http.Handle("/hello", ht)
		}
		err := http.ListenAndServe(":"+yamlConfig.HTTP.Port, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err.Error())
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
