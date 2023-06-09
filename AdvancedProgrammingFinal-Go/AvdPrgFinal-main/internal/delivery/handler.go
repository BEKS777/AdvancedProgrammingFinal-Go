package delivery

import (
	"forum/internal/service"
	"html/template"
	"net/http"
)

type Handler struct {
	Mux      *http.ServeMux
	Tmpl     *template.Template
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Mux:      http.NewServeMux(),
		Tmpl:     template.Must(template.ParseGlob("/Users/tikosch/Downloads/forum 5/ui/html/*.html")),
		Services: services,
	}
}

func (h *Handler) InitRoutes() {
	h.Mux.HandleFunc("/", h.homePage)

	h.Mux.HandleFunc("/auth/signin", h.signIn)
	h.Mux.HandleFunc("/auth/signup", h.signUp)
	h.Mux.HandleFunc("/auth/logout", h.logOut)

	h.Mux.HandleFunc("/post/", h.post)
	h.Mux.HandleFunc("/post/create", h.createPost)
	// h.Mux.HandleFunc("/post/delete/", h.deletePost)
	// h.Mux.HandleFunc("/post/change/", h.changePost)
	h.Mux.HandleFunc("/post/like/", h.likePost)
	h.Mux.HandleFunc("/post/dislike/", h.dislikePost)
	h.Mux.HandleFunc("/post/categories/", h.filterPostCategories)

	// h.Mux.HandleFunc("/comment/delete/", h.deleteComment)
	// h.Mux.HandleFunc("/comment/change/", h.changeComment)
	h.Mux.HandleFunc("/comment/like/", h.likeComment)
	h.Mux.HandleFunc("/comment/dislike/", h.dislikeComment)

	h.Mux.HandleFunc("/profile/", h.userProfilePage)

	h.Mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))
}
