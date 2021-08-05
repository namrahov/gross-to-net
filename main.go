package gross_to_net

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

}

func initEnvVars() {
	_ = godotenv.Load("profiles/" + *profile + ".env")
}
