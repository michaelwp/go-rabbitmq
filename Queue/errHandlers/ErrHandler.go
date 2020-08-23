package errHandlers

import "log"

// on error function
func FailOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
