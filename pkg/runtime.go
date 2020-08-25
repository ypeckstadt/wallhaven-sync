package pkg

import "log"

func LogFatalWhenError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
