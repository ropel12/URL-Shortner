package entity

import "fmt"

var ErrURLNotFound = fmt.Errorf("URL not found")
var ErrBadRequest = fmt.Errorf("bad request")
var ErrCustomURLIsExists = fmt.Errorf("custom url is exists")
