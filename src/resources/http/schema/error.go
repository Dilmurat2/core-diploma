package schema

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Err            error    `json:"-"`
	HTTPStatusCode int      `json:"-"`
	ErrorMessage   *Details `json:"error"`
}

type Details struct {
	StatusText  string `json:"status"`
	AppCode     int64  `json:"code,omitempty"`
	MessageText string `json:"message,omitempty"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)

	return nil
}

func BadRequest(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage: &Details{
			AppCode:     http.StatusBadRequest,
			StatusText:  http.StatusText(http.StatusBadRequest),
			MessageText: err.Error(),
		},
	}
}
