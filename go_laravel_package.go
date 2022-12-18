package go_laravel_package

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "0.0.1"

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
	Version     string
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
			"handlers",
			"migrations",
			"views",
			"data",
			"public",
			"tmp",
			"logs",
			"middleware",
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
			2.3 set the essential environments and version
				in GoLaravelPackage struct as
				application started
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
	glp.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	/*
		phase 3:

			3.1 create info and error loggers
	*/

	infoLog, errorLog := glp.startLoggers()
	glp.InfoLog = infoLog
	glp.ErrorLog = errorLog

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

	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
