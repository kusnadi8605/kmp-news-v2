package main

import (
	"flag"
	"fmt"
	conf "kmp-news-v2/config"
	hdlr "kmp-news-v2/handler"
	log "kmp-news-v2/logging"
	mdw "kmp-news-v2/middleware"

	"net/http"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load config file
	configFile := flag.String("conf", "config/config.yml", "main configuration file")

	log.Logf("OS: %s", runtime.GOOS)
	log.Logf("architecture: %s", runtime.GOARCH)

	flag.Parse()

	log.Logf("reads configuration from %s", *configFile)
	conf.LoadConfigFromFile(configFile)

	log.Init(conf.Param.Log.Level, conf.Param.Log.FileName)

	//redis connection
	conf.RedisDbInit(conf.Param.RedisURL)

	//mysql connection
	conn, err := conf.New(conf.Param.DBType, conf.Param.DBUrl)
	log.Logf("Load Database Conf: %s ", conf.Param.DBType)
	log.Logf("running App on port: %s ", conf.Param.ListenPort)

	if err != nil {
		log.Errorf("Unable to open database %v", err)
		os.Exit(1)
	}

	//elastic connection
	elasticConn, err := conf.ElasticConn(conf.Param.ElasticURL)

	if err != nil {
		log.Errorf("Unable to open elasticsearch %v", err)
		os.Exit(1)
	}

	http.HandleFunc("/api/news", mdw.Chain(
		hdlr.GetNewsHandler(
			conn,
			elasticConn,
			conf.Param.ElasticIndex,
			conf.Param.ElasticPerpage,
		), mdw.Method("GET")))

	var errors error
	errors = http.ListenAndServe(conf.Param.ListenPort, nil)

	if errors != nil {
		fmt.Println("error", errors)
		log.Logf("Unable to start the server: %s ", conf.Param.ListenPort)
		os.Exit(1)
	}
}
