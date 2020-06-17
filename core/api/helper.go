package api

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func genUID() string {
	u := uuid.Must(uuid.NewV4(), errors.New("UUID Error"))
	return u.String()
}

func HTTPResponse(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Println("HTTP ERROR", err.Error())
		return "ERROR"
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print("HTTP ERROR", err.Error())
		return "ERROR"
	}
	return string(body)
}

func ResponseXML(response string, data interface{}) error {
	err := xml.Unmarshal([]byte(response), data)
	if err != nil {
		log.Println("XML ERROR", err.Error())
	}
	return err
}
