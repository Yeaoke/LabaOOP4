package http

import "errors"

var (
	ErrorLaunchServer       = errors.New("Couldn't launch server")
	ErrorUnknownTypeCompany = errors.New("Unknown company type")
)
