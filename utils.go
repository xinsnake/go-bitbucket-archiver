package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func saveStructToFileAsJson(c interface{}, filename string) {
	var bytes []byte

	if bytes, err = json.Marshal(c); err != nil {
		fmt.Printf(err.Error())
		os.Exit(4)
	}

	if err = ioutil.WriteFile(ClientConfigFile, bytes, 0644); err != nil {
		return errors.New("Unable to write file, please check the permission settings.")
	}
}
