package utils

import "net/http"

type Utilsi interface {
	Pasejson(*http.Request, any) (any, error)
	Decoderjson(*http.Request, any) (any, error)
}

type Util struct{}

func Newutil() Utilsi {
	return &Util{}
}

func (u *Util) Pasejson(r *http.Request, ty any) (any, error) {
	return nil, nil
}

func (u *Util) Decoderjson(r *http.Request, ty any) (any, error) {
	return nil, nil
}
