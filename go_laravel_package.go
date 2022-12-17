package go_laravel_package

const version = "0.0.1"

type GoLaravelPackage struct {
	AppName     string
	Author      string
	AuthorMail  string
	Description string
	Debug       bool
	License     string
	Repository  string
	Version     string
}

func (glp *GoLaravelPackage) New(rootPath string) error {
	/*
		fetching the application root
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
