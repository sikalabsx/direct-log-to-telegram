package handle_error

import (
	"log"
)

func HandleErrorFatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
