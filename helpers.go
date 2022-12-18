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

func (glp *GoLaravelPackage) CreateFileIfNotExists(path string) error {
	var _, err = os.Stat(path)

	/*
	 create file
	 with given
	 path
	*/
	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err != nil {
			return err
		}

		// properly closing the created file
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}
