package main

import (
	"net/http"

	"github.com/AhmedARmohamed/crash_Course/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
