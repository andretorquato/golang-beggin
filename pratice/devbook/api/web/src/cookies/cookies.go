package cookies

import (
	"net/http"
	"web/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, token string) error {
	values := map[string]string{
		"id":    ID,
		"token": token,
	}

	valuesEncoded, err := s.Encode("data", values)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    valuesEncoded,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("data")
	if erro != nil {
		return nil, erro
	}

	values := make(map[string]string)
	if erro = s.Decode("data", cookie.Value, &values); erro != nil {
		return nil, erro
	}

	return values, nil
}
