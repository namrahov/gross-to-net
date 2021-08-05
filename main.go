package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/namrahov/gross-to-net/config"
	mw "github.com/namrahov/gross-to-net/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	port    *string
	profile *string
)

func init() {
	port = flag.String("port", ":8443", "Port number")
	profile = flag.String("profile", "local", "Application run profile")
}

func main() {
	flag.Parse()

	log.Info("Application is starting with profile: ", *profile)

	initEnvVars()

	config.LoadConfig()

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(mw.Default)

	handler.convertGrossToNet(router)

	log.Info("Starting server at port ", *port)
	log.Fatal(http.ListenAndServe(*port, router))

}

func initEnvVars() {
	_ = godotenv.Load("profiles/" + *profile + ".env")
}
