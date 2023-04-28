package data

import "time"

/*Secret is struct that contains the secret key used for signing JWT tokens and ExpiresAt that represents until when JWT would be used and valid
 */
type Secret struct {
	Secret    string    `json:"secret"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Credentials struct {
	JMBG     string `json:"JMBG"`
	Password string `json:"Password"`
}

type JWTReturn struct {
	Token string `json:"token"`
}

type Gradjanin struct {
	Ime          string  `json:"ime"`
	Prezime      string  `json:"prezime"`
	JMBG         string  `json:"jmbg"`
	Adresa       string  `json:"adresa"`
	BrojTelefona string  `json:"brojTelefona"`
	Email        string  `json:"email"`
	Lozinka      string  `json:"lozinka"`
	Opstina      Opstina `json:"opstina"`
}

type GradjaninResponseDTO struct {
	Ime          string  `json:"ime"`
	Prezime      string  `json:"prezime"`
	JMBG         string  `json:"jmbg"`
	Adresa       string  `json:"adresa"`
	BrojTelefona string  `json:"brojTelefona"`
	Email        string  `json:"email"`
	Opstina      Opstina `json:"opstina"`
}

type Opstina struct {
	PTT   string `json:"PTT"`
	Naziv string `json:"Naziv"`
}
