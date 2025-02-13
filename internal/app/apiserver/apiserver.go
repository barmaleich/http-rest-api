package apiserver

import (
	"github.com/barmaleich/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil{
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil{
		return err
	}

	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter()  {
	s.router.HandleFunc("/test", s.handleHello())
	//s.router.HandleFunc("/user", s.handleGetUserByEmail())
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil{
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

//func (s *APIServer) handleGetUserByEmail() http.HandlerFunc {
//	s.logger.Info("GET user ")
//	u, _ := s.store.User().FindByEmail("nbreikin@mail.ru")
//	return func(w http.ResponseWriter, r *http.Request) {
//		_, _ = io.WriteString(w, u.Email)
//	}
//
//}

func (s *APIServer) handleHello() http.HandlerFunc {
	s.logger.Info("GET hello")
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello")
	}
	
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil{
		return err
	}

	s.store = st

	return nil
}