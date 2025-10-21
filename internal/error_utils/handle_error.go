package error_utils

import (
	"log"
)

func HandleErrorFatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
