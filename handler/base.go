package handler

import (
	"net/http"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type BeforeFilter interface {
	DoFilter(http.ResponseWriter, *http.Request) bool
}

type AfterFilter interface {
	DoFilter(http.ResponseWriter, *http.Request) bool
}

type BaseHandler struct {
	BeforeFilters BeforeFilter
	AfterFilters  AfterFilter
	Handler       Handler
}

func (b *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.BeforeFilters != nil {
		if b.BeforeFilters.DoFilter(w, r) == false {
			return
		}
	}

	b.Handler.Handle(w, r)

	if b.AfterFilters != nil {
		if b.AfterFilters.DoFilter(w, r) == false {
			return
		}
	}
}

func NewBaseHandler(handler Handler, beforeFilters BeforeFilter, afterFilters AfterFilter) *BaseHandler {
	return &BaseHandler{
		BeforeFilters: beforeFilters,
		AfterFilters:  afterFilters,
		Handler:       handler,
	}
}
