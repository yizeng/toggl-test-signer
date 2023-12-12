package web

import (
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
	"github.com/yizeng/toggl-test-signer/internal/repository"
	"github.com/yizeng/toggl-test-signer/internal/repository/dao"
	"github.com/yizeng/toggl-test-signer/internal/service"
	v1 "github.com/yizeng/toggl-test-signer/internal/web/handler/v1"
)

type Server struct {
	Address string
	Router  *chi.Mux
}

func NewServer() *Server {
	db := initDB()

	userDAO := dao.NewTestDAO(db)
	userRepo := repository.NewTestRepository(userDAO)
	userSvc := service.NewUserService(userRepo)
	userHandler := v1.NewUserHandler(userSvc)

	adminSvc := service.NewAdminService(userRepo)
	adminHandler := v1.NewAdminHandler(adminSvc)

	s := &Server{
		Address: getServerAddress(),
		Router:  chi.NewRouter(),
	}

	s.MountMiddlewares()
	s.MountHandlers(adminHandler, userHandler)

	return s
}

func initDB() *gorm.DB {
	dsn := viper.GetString("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		// 我只会在初始化过程中 panic
		// panic 相当于整个 goroutine 结束
		// 一旦初始化过程出错，应用就不要启动了
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func getServerAddress() string {
	host := viper.Get("HOST")
	port := viper.Get("PORT")
	addr := fmt.Sprintf("%v:%v", host, port)

	return addr
}

func (s *Server) MountMiddlewares() {
	jwtSecret := viper.GetString("JWT_SECRET")
	tokenAuth := jwtauth.New("HS512", []byte(jwtSecret), nil)

	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Heartbeat("/"))
	s.Router.Use(jwtauth.Verifier(tokenAuth))
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))
}

func (s *Server) MountHandlers(adminHandler *v1.AdminHandler, userHandler *v1.UserHandler) {
	apiV1Router := chi.NewRouter()
	apiV1Router.Route("/", func(r chi.Router) {
		r.Post("/users/sign-answers", userHandler.HandleSignTest)
		r.Post("/admin/verify-signature", adminHandler.HandleVerifySignature)
	})

	s.Router.Mount("/api/v1", apiV1Router)

	s.printAllRoutes()
}

func (s *Server) printAllRoutes() {
	zap.L().Debug("printing all routes...")

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)

		zap.L().Debug(fmt.Sprintf("%v\t%v", method, route))

		return nil
	}

	if err := chi.Walk(s.Router, walkFunc); err != nil {
		zap.L().Error("printing all routes failed", zap.Error(err))
	}
}
