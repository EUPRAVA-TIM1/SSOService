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
	ssoService       service.SSOservice
	gradjaninService service.GradjaniService
}

func NewFilesHandler(sso service.SSOservice, gs service.GradjaniService) SSOHandler {
	return ssoHandler{ssoService: sso, gradjaninService: gs}
}

func (s ssoHandler) Init(r *mux.Router) {
	r.StrictSlash(false)
	r.HandleFunc("/sso/Login", s.Login).Methods("POST")
	r.HandleFunc("/sso/Secret", s.GetSecret).Methods("GET")
	r.HandleFunc("/sso/Whoami/{jmbg}", s.Whoami).Methods("GET")
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

// TEST za sada bez tokena
func (s ssoHandler) Whoami(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jmbg := vars["jmbg"]
	user, err := s.gradjaninService.GetByJMBG(jmbg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	jsonResponse(data.GradjaninResponseDTO{
		Ime:          user.Ime,
		Prezime:      user.Prezime,
		JMBG:         user.JMBG,
		Adresa:       user.Lozinka,
		BrojTelefona: user.BrojTelefona,
		Email:        user.Email,
		Opstina: data.Opstina{
			PTT:   user.Opstina.PTT,
			Naziv: user.Opstina.PTT,
		},
	}, w, http.StatusOK)
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
