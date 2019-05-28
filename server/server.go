package server

import (
	"{{ .pkgName}}/model"
	"{{ .pkgName}}/model/mock"
	"{{ .pkgName}}/model/mysql"
	"{{ .pkgName}}/version"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Storage model.Storage
	Logger  logrus.FieldLogger
	Config  Config

	webServer *gin.Engine
}

func NewServer(options ...Option) *Server {

	s := Server{}

	cfg := Config{}
	for _, o := range options {
		o(&cfg)
	}

	if cfg.MysqlDSN != "" {
		e, err := xorm.NewEngine(`mysql`, cfg.MysqlDSN)
		checkErr(err)
		s.Storage = mysql.NewMysqlStorage(e)
	}

	if cfg.EnableMockStorage {
		s.Storage = mock.NewMockStorage()
	}

	s.webServer = newGinServer()

	return &s
}

func newGinServer() *gin.Engine {
	g := gin.Default()
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	g.GET("/version", func(ctx *gin.Context) {
		ctx.String(200, "%d", version.Version)
	})

	group := g.Group("/api")

	{
		group.GET("/", func(ctx *gin.Context) {
			// TODO:: fix me
		})
	}

	return g
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	return s.webServer.Run(s.Config.Addr)
}
