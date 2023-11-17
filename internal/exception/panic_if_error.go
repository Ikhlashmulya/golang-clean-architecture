package exception

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}