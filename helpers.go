package go_laravel_package

import "os"

func (glp *GoLaravelPackage) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	/*
	 create folder
	 with given
	 path
	*/
	return nil
}
