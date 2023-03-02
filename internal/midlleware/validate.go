package midlleware

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/juanguaje/api-template-juanguaje/internal/entity"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func validardatos() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			request := entity.Request{}
			bRequest, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(rw, "Error: %v", err)
				return
			}
			if err = json.Unmarshal(bRequest, &request); err != nil {
				fmt.Fprintf(rw, "Error: %v", err)
				return
			}
			rw.Write([]byte("Los datos fueron validados correctamente\n"))
			f(rw, r)
		}
	}
}

func validarusuario() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Usuario logueado correctamente\n"))
			f(w, r)
		}
	}
}

func PostRequestAllPharmacy(url string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var (
				respbody []byte
			)

			bRequest, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error 1: %v", err)
				return
			}

			request := entity.Request{}
			if len(bRequest) != 0 {
				if err = json.Unmarshal(bRequest, &request); err != nil {
					fmt.Fprintf(w, "Error 2: %v", err)
					return
				}
			}

			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "Error 3: %v", err)
				return
			}
			respbody, err = io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(w, "Error 4: %v", err)
				return
			}

			respbody = bytes.TrimPrefix(respbody, []byte("\xef\xbb\xbf"))

			if request.FilterValue != "" {
				respbody, err = filterPharmacies(request, respbody)
				if err != nil {
					fmt.Fprintf(w, "Error 5: %v", err)
					return
				}
			}

			w.Write([]byte(respbody))
		}
	}
}

func AgregarMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, mwr := range middlewares {
		f = mwr(f)
	}
	return f
}

func filterPharmacies(filter entity.Request, phatmacyList []byte) ([]byte, error) {
	list := []entity.Pharmacy{}
	listOut := []entity.Pharmacy{}
	err := json.Unmarshal(phatmacyList, &list)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	for i := 0; i < len(list); i++ {
		if list[i].CommuneName == filter.FilterValue {
			listOut = append(listOut, list[i])
		}
	}

	if filter.TypeResponse == "XML" {
		return xml.Marshal(&listOut)
	}

	return json.Marshal(listOut)
}

func MensajeHandler(mensaje string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mensaje))
	})
}

func MensajeMiddleware(mensaje string, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		w.Write([]byte(mensaje))
	})
}
