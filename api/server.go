package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"google.golang.org/appengine"

	"github.com/gorilla/mux"
)

// InitServer creates a new server, adds to net/http default handler
// and returns the server
func InitServer() http.Handler {
	s := new(Server)
	s.RegisterService("/api/user", new(UserService))

	http.Handle("/", s)
	return s
}

// Server is the main http handler for our api
type Server struct {
	router   *mux.Router
	services map[string]Service
}

func contextHandler(cr ContextResponder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		rs := cr(ctx)

		if rs == nil {
			return
		}

		e := json.NewEncoder(w)
		e.Encode(rs)
	}
}

// RegisterService add service to server with given url
func (s *Server) RegisterService(url string, srv Service) {
	if s.services == nil {
		s.services = map[string]Service{}
	}
	s.services[url] = srv
}

// CreateRouter creates a mux router for url matching
func (s *Server) CreateRouter() {
	router := mux.NewRouter()
	s.router = router
	for url, service := range s.services {
		s.BindService(url, service)
	}
}

// BindService creates a routes from a service and adds it to the
// server's router
func (s *Server) BindService(url string, service Service) {
	ms := []string{}
	for m := range service.Methods() {
		ms = append(ms, m)
	}

	s.router.NewRoute().
		Name(service.Name()).
		Methods(ms...).Path(url).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rm := strings.ToUpper(r.Method)
			m := service.Methods()[rm]
			contextHandler(m)(w, r)
		})

	for u, h := range service.AdditionalRoutes() {
		s.router.HandleFunc(url+u, contextHandler(h))
	}
}

// ServeHTTP fulfills http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.router == nil {
		s.CreateRouter()
	}

	// Adding HTTP OPTIONS support, should add more security here
	// http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if r.Method == "OPTIONS" {
		return
	}

	s.router.ServeHTTP(w, r)
}
