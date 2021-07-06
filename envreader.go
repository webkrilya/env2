package envreader

import (
	"github.com/kardianos/osext"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var Env map[string]string

func ReadEnv() map[string]string {
	var err error
	Env = make(map[string]string)

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal("Cant find env folder path", err)
	}
	file, err := os.Open(folderPath + "/.env")
	if err != nil {
		file, err = os.Open(".env")
	}
	if err != nil {
		log.Fatal("Cant find env file", err)
	}

	b, _ := ioutil.ReadAll(file)

	rows := strings.Split(string(b), "\n")
	for _, v := range rows {
		if v == "" || strings.Index(v, "#") == 0 {
			continue
		}

		pos := strings.Index(v,"=")
		if pos>-1{
			Env[strings.TrimSpace(v[:pos])] = strings.TrimSpace(v[pos:])
		}
	}

	return Env
}
