package service

import (
	"fmt"
	"reflect"
)

type InvalidState struct {
	msg string
}

func (i InvalidState) Error() string {
	return i.msg
}

type NotFoundError struct {
	Type interface{}
	Id   int
}

func (i NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID[%d] not found!", reflect.TypeOf(i.Type).Name(), i.Id)
}
