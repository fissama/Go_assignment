package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config of application
type Config struct {
	AppVersion string
	MongoDB    MongoDB
}

type MongoDB struct {
	URI        string
	Cluster    string
	User       string
	Password   string
	DB         string
	Collection string
}

func GetConfig(params ...string) MongoDB {
	configuration := MongoDB{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}

	content, err := ioutil.ReadFile("./configs/config_" + env + ".json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `configuration`
	err = json.Unmarshal(content, &configuration)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return configuration
}
