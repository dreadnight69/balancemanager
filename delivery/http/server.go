package http

import (
	"balanceManager/balancemanager"
	"balanceManager/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct {
	router         *mux.Router
	config         config.Config
	BalanceManager balancemanager.Service
}

func WithConfig(cfg config.Config, bm balancemanager.Service) *server {
	r := mux.NewRouter()
	res := &server{
		router:         r,
		config:         cfg,
		BalanceManager: bm,
	}
	res.initRouters()
	return res
}

func (s *server) initRouters() {
	s.router.HandleFunc("/balance", s.getBalance).Methods(http.MethodGet)
	s.router.HandleFunc("/withdrawal", s.withdrawFunds).Methods(http.MethodPost)
	s.router.HandleFunc("/deposit", s.makeDeposit).Methods(http.MethodPost)
	s.router.HandleFunc("/sendfunds", s.sendFunds).Methods(http.MethodPost)
	s.router.HandleFunc("/transactions", s.getTransactions).Methods(http.MethodGet)

	s.router.HandleFunc("/createuser", s.createUser).Methods(http.MethodGet)
}

func (s *server) Run() {
	serv := &http.Server{
		Addr:         s.config.Http.ListenHost + ":" + s.config.Http.ListenPort,
		Handler:      s.router,
		ReadTimeout:  s.config.Http.ServerRTimeout,
		WriteTimeout: s.config.Http.ServerWTimeout,
	}
	log.Println("HTTP server started on address " + serv.Addr)
	if err := serv.ListenAndServe(); err != nil {
		log.Fatal("HTTP server failed:", err)
	}
}
