/**
*	NAME: Aaron Meek
*	DATE: 2022 - 08 - 15
*
*	This contains all page handlers and repository structures
*	for app configuration
 */

package handlers

import (
	"net/http"

	"github.com/urhumantoast/NewPersonalWebsite/pkg/config"
	"github.com/urhumantoast/NewPersonalWebsite/pkg/models"
	"github.com/urhumantoast/NewPersonalWebsite/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// Repo - The repository used by the handlers
var Repo *Repository

// NewRepo - Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers - Sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home - Renders home about page and displays form
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About - Renders the about page and displays form
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}

// Contact - Renders the contact page and displays form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}

// ContactComplete - Renders the contact-complete page and displays form
func (m *Repository) ContactComplete(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact-complete.page.html", &models.TemplateData{})
}

// ProjMain - Renders the main project page and displays form
func (m *Repository) ProjMain(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "proj-main.page.html", &models.TemplateData{})
}

// ProjApp - Renders the applications & software page and displays form
func (m *Repository) ProjApp(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "proj-app.page.html", &models.TemplateData{})
}

// ProjEmb - Renders the embedded systems page and displays form
func (m *Repository) ProjEmb(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "proj-emb.page.html", &models.TemplateData{})
}

// ProjElc - Renders the electronics page and displays form
func (m *Repository) ProjElc(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "proj-elc.page.html", &models.TemplateData{})
}

// Placeholder - Renders the placeholder page and displays form
func (m *Repository) Placeholder(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "placeholder.page.html", &models.TemplateData{})
}
