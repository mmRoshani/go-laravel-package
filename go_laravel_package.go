package go_laravel_package

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "0.0.1"

/*
glp, go laravel package is accessible
for any other project
that use it
*/
type GoLaravelPackage struct {
	AppName     string
	Author      string
	AuthorMail  string
	Debug       bool
	Description string
	ErrorLog    *log.Logger
	InfoLog     *log.Logger
	License     string
	Repository  string
	RootPath    string
	Routes      *chi.Mux
	Version     string
	config      config
}

// glp configuration

type config struct {
	port     string
	renderer string
}

func (glp *GoLaravelPackage) New(rootPath string) error {

	/*
		phase 1:

			1.1 fetching the application root
			path and initializing
			necessary folders
	*/

	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"data",
			"handlers",
			"logs",
			"middleware",
			"migrations",
			"public",
			"tmp",
			"views",
		},
	}

	err := glp.Init(pathConfig)

	if err != nil {
		return err
	}

	/*
		phase 2:

			2.1	creating the .env file if dose not exist
			2.2	reading the .env file as it holding the project
				configurations, then populate them
			 	as system environment
			2.3 set the essential environments and configuration
			 	in GoLaravelPackage struct
				as application started
	*/

	// 2.1
	err = glp.checkDotEnv(rootPath)

	if err != nil {
		return err
	}

	// 2.2
	err = godotenv.Load(rootPath + "/.env")

	if err != nil {
		return err
	}
	// 2.3
	glp.RootPath = rootPath
	glp.Version = version

	glp.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	glp.config.port = os.Getenv("PORT")
	glp.config.renderer = os.Getenv("RENDERER")

	infoLog, errorLog := glp.startLoggers()
	glp.ErrorLog = errorLog
	glp.InfoLog = infoLog

	glp.Routes = glp.routes().(*chi.Mux)

	return nil
}

func (glp *GoLaravelPackage) Init(p initPaths) error {
	root := p.rootPath
	// checking for necessary folder paths existence
	for _, path := range p.folderNames {
		// create folder if the path dose not exist
		err := glp.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (glp *GoLaravelPackage) checkDotEnv(path string) error {
	err := glp.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}

func (glp *GoLaravelPackage) startLoggers() (*log.Logger, *log.Logger) {

	var errorLog *log.Logger
	var infoLog *log.Logger

	errorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)

	return infoLog, errorLog
}

// http server opener
func (glp *GoLaravelPackage) ListenAndServe() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", glp.config.port),
		ErrorLog:     glp.ErrorLog,
		Handler:      glp.routes(),
		IdleTimeout:  25 * time.Second,
		ReadTimeout:  25 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	glp.InfoLog.Printf("Server start listening on port %s", glp.config.port)
	err := server.ListenAndServe()
	glp.ErrorLog.Fatal(err)
}
