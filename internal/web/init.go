package web

import (
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
	"github.com/yizeng/toggl-test-signer/internal/service"
	v1 "github.com/yizeng/toggl-test-signer/internal/web/handler/v1"
)

type Server struct {
	Address string
	Router  *chi.Mux
}

func NewServer() *Server {
	s := &Server{
		Address: getServerAddress(),
		Router:  chi.NewRouter(),
	}

	s.MountMiddlewares()
	s.MountHandlers()

	return s
}

func getServerAddress() string {
	host := viper.Get("HOST")
	port := viper.Get("PORT")
	addr := fmt.Sprintf("%v:%v", host, port)

	return addr
}

func (s *Server) MountMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Heartbeat("/"))
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))
}

func (s *Server) MountHandlers() {
	apiV1Router := chi.NewRouter()
	apiV1Router.Route("/", func(r chi.Router) {
		adminSvc := service.NewAdminService()
		adminHandler := v1.NewAdminHandler(adminSvc)
		r.Post("/admin/verify-signature", adminHandler.HandleVerifySignature)

		userSvc := service.NewUserService()
		userHandler := v1.NewUserHandler(userSvc)
		r.Post("/users/sign-answers", userHandler.HandleSignAnswers)
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
