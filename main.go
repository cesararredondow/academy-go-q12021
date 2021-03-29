package main

import (
	"encoding/csv"
	"net/http"
	"os"

	"github.com/cesararredondow/course/first_deliverable/handlers"
	"github.com/cesararredondow/course/first_deliverable/routes"
	"github.com/cesararredondow/course/first_deliverable/services"
	"github.com/cesararredondow/course/first_deliverable/usecases"
	secondHandler "github.com/cesararredondow/course/second_deliverable/handlers"
	secondRouter "github.com/cesararredondow/course/second_deliverable/routes"
	secondService "github.com/cesararredondow/course/second_deliverable/services"
	secondUsecase "github.com/cesararredondow/course/second_deliverable/usecases"
	thirdHandler "github.com/cesararredondow/course/third_deriverable/handlers"
	thirdRoutes "github.com/cesararredondow/course/third_deriverable/routes"
	thirdService "github.com/cesararredondow/course/third_deriverable/services"
	thirdUsecase "github.com/cesararredondow/course/third_deriverable/usecases"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

const (
	ExitAbnormalErrorLoadingConfiguration = iota
	ExitAbnormalErrorLoadingCsvFile
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	filePath := goDotEnvVariable("filePath")
	PORT := goDotEnvVariable("PORT")
	LogLevel := goDotEnvVariable("lOGLEVEL")
	pokemonAPI := goDotEnvVariable("pokemonAPI")

	level, err := log.ParseLevel(LogLevel)
	if err != nil {
		log.Fatal("Failed creating logger: %w", err)
	}

	l := log.New()
	l.SetLevel(level)
	l.Out = os.Stdout
	l.Formatter = &log.JSONFormatter{}
	render := render.New()
	l.Info("starting the app")

	router := mux.NewRouter()
	rf, err := os.Open(filePath)
	if err != nil {
		os.Exit(ExitAbnormalErrorLoadingConfiguration)
	}

	wf, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		os.Exit(ExitAbnormalErrorLoadingCsvFile)
	}
	defer rf.Close()
	defer wf.Close()

	csvw := csv.NewWriter(wf)

	s1, _ := services.New(rf)
	u1 := usecases.New(s1)
	h1 := handlers.New(u1, l, render)
	routes.New(h1, router)

	s2, _ := secondService.New(csvw, pokemonAPI, filePath)
	u2 := secondUsecase.New(s2)
	h2 := secondHandler.New(u2, l, render)
	secondRouter.New(router, h2)

	s3, _ := thirdService.New(rf)
	u3 := thirdUsecase.New(s3)
	h3 := thirdHandler.New(u3, l, render)
	thirdRoutes.New(h3, router)

	log.Info(http.ListenAndServe(":"+PORT, router))
}
