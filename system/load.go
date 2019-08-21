package system

import (
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

// 读取数据库配置
type dbConf struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
}

var (
	cfg          *ini.File
	SemesterConf *semesterConf
	CrawlServer  string
	ArticleNum   int
)

func initIni() {
	cfg = initCfg()
	SemesterConf = readSemesterConf()
	CrawlServer = cfg.Section("server").Key("CrawlServer").String()
	ArticleNum, _ = cfg.Section("server").Key("ReturnArticleNum").Int()
}

func initCfg() *ini.File {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		Save.ReadIniPanic("Fail to read file ", err)
	}
	return cfg
}

// 读取服务的配置
func ReadDbIni() *dbConf {
	return &dbConf{
		DriverName: cfg.Section("database").Key("DriverName").String(),
		Host:       cfg.Section("database").Key("Host").String(),
		Port:       cfg.Section("database").Key("Port").String(),
		Database:   cfg.Section("database").Key("Database").String(),
		Username:   cfg.Section("database").Key("Username").String(),
		Password:   cfg.Section("database").Key("Password").String(),
	}
}

type serverConf struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// 读取学期的配置
func ReadServerIni() *serverConf {
	httpPort, err := cfg.Section("server").Key("HttpPort").Int()
	if err != nil {
		Save.ReadIniPanic("Fail to read HttpPort configuration ", err)
	}
	readTimeOut, err := cfg.Section("server").Key("ReadTimeout").Int()
	if err != nil {
		Save.ReadIniPanic("Fail to read ReadTimeout configuration ", err)
	}
	writeTimeout, err := cfg.Section("server").Key("ReadTimeout").Int64()
	if err != nil {
		Save.ReadIniPanic("Fail to read WriteTimeout configuration ", err)
	}
	return &serverConf{
		RunMode:      cfg.Section("server").Key("RunMode").String(),
		HttpPort:     fmt.Sprintf(":%d", httpPort),
		ReadTimeout:  time.Duration(readTimeOut) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}
}

type semesterConf struct {
	SchoolYear string
	TeachWeek  int
	Semester   string
}

func readSemesterConf() *semesterConf {
	teachWeek, err := cfg.Section("semester").Key("TeachWeek").Int()
	if err != nil {
		Save.ReadIniPanic("Fail to read TeachWeek configuration ", err)
	}
	return &semesterConf{
		SchoolYear: cfg.Section("semester").Key("SchoolYear").String(),
		TeachWeek:  teachWeek,
		Semester:   cfg.Section("semester").Key("Semester").String(),
	}
}
