package main

import (
	"net/http"

	"github.com/juanguaje/api-template-juanguaje/config"
	v "github.com/juanguaje/api-template-juanguaje/internal/midlleware"
	l "github.com/juanguaje/api-template-juanguaje/log"
)

func main() {
	cfg := config.LoadConfigProvider("api-pharmacy")

	http.HandleFunc("/", v.AgregarMiddleware(v.MensajeHandler("Peticion Http"), v.PostRequestAllPharmacy(cfg.GetString("url"))))
	log := l.NewLogger(cfg)
	log.Infof("Server Init")
	err := http.ListenAndServe(cfg.GetString("host")+":"+cfg.GetString("port"), nil)
	if err != nil {
		log.Errorf("error al acceder al servidor: ", err)
		return
	}
}
