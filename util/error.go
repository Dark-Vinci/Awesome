package util

import "errors"

var FileNotFountError = errors.New("file not found")
var InternalServerError = errors.New("something went wrong")
var NoFileUploaded = errors.New("no file has been uploaded")
var FileNotUploaded = errors.New("file with name 'file' is missing in request")
