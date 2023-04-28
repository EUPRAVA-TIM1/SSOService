package handlers

import (
	"EuprvaSsoService/data"
	"EuprvaSsoService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type SSOHandler interface {
	Init(r *mux.Router)
}

const (
	badRequestMsg       = "Bad Request"
	contentType         = "Content-Type"
	internalSrvErrMsg   = "Internal server error"
	unsupportedMediaMsg = "Unsupported media type"
)

type ssoHandler struct {
	ssoService service.SSOservice
}

func NewFilesHandler(s service.SSOservice) SSOHandler {
	return ssoHandler{ssoService: s}
}

func (s ssoHandler) Init(r *mux.Router) {
	r.StrictSlash(false)
	r.HandleFunc("/sso/Login", s.Login).Methods("POST")
	r.HandleFunc("/sso/Secret", s.GetSecret).Methods("GET")
	http.Handle("/", r)
}

func (s ssoHandler) Login(w http.ResponseWriter, r *http.Request) {
	var usr data.Credentials
	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		http.Error(w, "Problem with decoding JSON", http.StatusBadRequest)
		return
	}

	if len(usr.JMBG) != 13 {
		http.Error(w, "JMBG must be 13 chars long", http.StatusBadRequest)
		return
	}

	token, err := s.ssoService.Login(usr)

	if err != nil {
		switch err {
		case service.JWTError:

			jsonResponse(err.Error(), w, http.StatusInternalServerError)
			//case service.WrongCredentials:
			//	er := handler.logLimitService.Increment(usr.Username, ctx)
			//	if er != nil {
			//		span.SetStatus(codes.Error, err.Error())
			//		return
			//	}
			//	jsonResponse(err.Error(), w, http.StatusBadRequest)
			//case service.DoesntExistsError:
			//	jsonResponse(err.Error(), w, http.StatusNotFound)
		}
		return
	}
	jsonResponse(token, w, http.StatusOK)
}

func (s ssoHandler) GetSecret(w http.ResponseWriter, r *http.Request) {
	issuer := r.Header.Get("X-Service-Name")

	secret, err := s.ssoService.GetSecret(issuer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(secret, w, http.StatusOK)
}

func jsonResponse(object interface{}, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(object)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status != 0 {
		w.WriteHeader(status)
	}
	_, err = w.Write(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
