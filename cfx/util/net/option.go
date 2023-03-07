package net

import "net/http"

type Option interface {
	ApplyOption(r *http.Request)
}
type HeaderOption struct {
	Key string
	Val string
}

func (ho *HeaderOption) ApplyOption(r *http.Request) {
	r.Header.Set(ho.Key, ho.Val)
}
