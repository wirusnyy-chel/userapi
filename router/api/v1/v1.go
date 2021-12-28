package v1

import (
	"net/http"
	"refactoring/model"
	"refactoring/model/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func New(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", searchUsers)
		r.Post("/", createUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getUser)
			r.Patch("/", updateUser)
			r.Delete("/", deleteUser)
		})
	})
}

func searchUsers(w http.ResponseWriter, r *http.Request) {
	list, err := user.GetList()
	if err == user.ErrUserNotFound {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err != nil {
		_ = render.Render(w, r, ErrInternalError(err))
		return
	}
	render.JSON(w, r, list)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	request := model.CreateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	id, err := user.Create(&request)
	if err != nil {
		_ = render.Render(w, r, ErrInternalError(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	u, err := user.GetByID(id)
	if err == user.ErrUserNotFound {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err != nil {
		_ = render.Render(w, r, ErrInternalError(err))
		return
	}
	render.JSON(w, r, u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	request := model.UpdateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")
	err := user.Update(id, request)
	if err == user.ErrUserNotFound {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err != nil {
		_ = render.Render(w, r, ErrInternalError(err))
		return
	}
	render.Status(r, http.StatusNoContent)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := user.Delete(id)
	if err == user.ErrUserNotFound {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err != nil {
		_ = render.Render(w, r, ErrInternalError(err))
		return
	}
	render.Status(r, http.StatusNoContent)
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInternalError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal error",
		ErrorText:      err.Error(),
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
