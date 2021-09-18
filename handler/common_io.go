package handler

import "net/http"

func getResponse(res interface{}) (int, interface{}) {
	return http.StatusOK, res
}

// func getErrorResponse(res interface{}) (int, interface{}) {
// 	return http.StatusInternalServerError, res
// }
