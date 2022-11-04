package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, req *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var appError *AppError

		err := h(w, req)
		if err != nil {
			if errors.As(err, &appError) {
				if errors.Is(err, ErrorNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrorNotFound.Marshal())
					return
				}
				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(ErrorBadRequest.Marshal())
				return
			}
			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err).Marshal())
			return
		}
	}
}
