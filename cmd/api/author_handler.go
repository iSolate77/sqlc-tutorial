package api

import (
	"errors"
	"net/http"
	"sqlc-tutorial/core"
	"sqlc-tutorial/internal/request"
	"sqlc-tutorial/internal/response"

	"github.com/jackc/pgx/v5/pgtype"
)

type AuthorRespnse struct {
	AuthorID   int         `json:"author_id"`
	AuthorName string      `json:"author_name"`
	AuthorBio  pgtype.Text `json:"author_bio"`
}

type CreateAuthorRequest struct {
	AuthorName string `json:"author_name"`
	AuthorBio  string `json:"author_bio"`
}

type AuthorController struct {
	authorService *core.Service
}

func NewAuthorController(authorService *core.Service) *AuthorController {
	return &AuthorController{authorService: authorService}
}

// func (c *AuthorController) Routes() http.Handler {
// 	router := http.NewServeMux()
// 	router.HandleFunc("POST /authors/create", c.Create)
// 	return router
// }

func (c *AuthorController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	dst := CreateAuthorRequest{}
	err := request.DecodePostForm(r, &dst)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid request")
		return
	}
	author, err := c.authorService.CreateAuthor(r.Context(), core.Author{
		Name: dst.AuthorName,
		Bio:  pgtype.Text{String: dst.AuthorBio},
	})
	if err != nil {
		errors.Join(err, errors.New("sqlc-tutorial: author already exists"))
		return
	}
	response.JSON(w, http.StatusCreated, author)
}
