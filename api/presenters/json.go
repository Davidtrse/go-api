package presenters

import (
	"api/utils"
	"net/http"
)

func (p *presenters) JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := utils.WriteJson(w, v)
	if err != nil {
		p.Error(w, r, err)
	}
}
