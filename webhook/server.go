package webhook

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Config Config
	Sender *Sender
}

func New(config Config) *Server {
	if config.Addr == "" {
		config.Addr = ":8080"
	}

	return &Server{
		Config: config,
		Sender: &Sender{
			Wechat: config.WechatRobot(),
		},
	}
}

func (s *Server) ListenAndServe() error {
	router := s.router()

	srv := &http.Server{
		Handler:      router,
		Addr:         s.Config.Addr,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s *Server) router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/sender", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		var notify Notification
		if json.Unmarshal(data, &notify) != nil {
			return
		}

		channel := r.URL.Query().Get("channel")
		webhook := r.URL.Query().Get("webhook")

		if channel == "" || webhook == "" {
			return
		}

		if err := s.Sender.Send(channel, webhook, &notify); err != nil {
			log.Println(err)
		}
	}).Methods("POST", "PUT").Name("sender")

	return router
}
