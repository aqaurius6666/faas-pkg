package lib

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ResponseOK(w http.ResponseWriter, data interface{}) {
	out := map[string]interface{}{
		"success": true,
		"data":    data,
		"code":    http.StatusOK,
	}
	ResponseJson(w, http.StatusOK, out)
}

func Response400(w http.ResponseWriter, err error) {
	out := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
		"code":    http.StatusBadRequest,
	}
	ResponseJson(w, http.StatusOK, out)
}

func Response401(w http.ResponseWriter, err error) {
	out := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
		"code":    http.StatusUnauthorized,
	}
	ResponseJson(w, http.StatusOK, out)
}

func ResponseError(w http.ResponseWriter, err error) {
	out := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
		"code":    http.StatusInternalServerError,
	}
	ResponseJson(w, http.StatusOK, out)
}
