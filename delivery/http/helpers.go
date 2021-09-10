package http

import (
	"balanceManager/balancemanager"
	"net/http"
)

func (s *server) statusByErr(err error) int {
	switch err {
	case balancemanager.ErrPgNoRows:
		return http.StatusBadRequest
	case nil:
		return http.StatusOK
	default:
		return http.StatusInternalServerError
	}
}
