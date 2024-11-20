package main

import "os"

func writeFile(data string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	file.Sync()

	return nil
}

func checkIfDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
        err := os.MkdirAll(path, 0755)
        if err != nil {
            return err
        }
    }
    return nil
}