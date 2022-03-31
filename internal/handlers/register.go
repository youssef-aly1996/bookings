package handlers

import (
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/forms"
	"github.com/youssef-aly1996/bookings/internal/models/user"
	"github.com/youssef-aly1996/bookings/internal/render"
	"golang.org/x/crypto/bcrypt"
)

func (repo *Repository) Login(rw http.ResponseWriter, r *http.Request)  {
	SetCsrf(r)
	td.Error = repo.App.Session.PopString(r.Context(), "error")
	render.Template(rw, "login.page.tmpl", td)
	td.Error = ""
}

func (repo *Repository) PostLogin(rw http.ResponseWriter, r *http.Request)  {
	_ = repo.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		repo.ServerErrors(rw, err)
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	f := forms.NewForm(r.PostForm)
	f.Required("email", "password")
	if !f.Valid() {
		td.Error = "invalid inputs"
		td.Form = f
		render.Template(rw, "login.page.tmpl", td)
		return
	}
	id, _, err := us.Auth(email, password)
	if err != nil {
		repo.App.Session.Put(r.Context(), "error", "invalid credentials")
		td.Error = repo.App.Session.PopString(r.Context(), "error")
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
	repo.App.Session.Put(r.Context(), "user_id", id)
	repo.App.Session.Put(r.Context(), "flash", "logged in successfully")
	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

func (repo *Repository) Signup(rw http.ResponseWriter, r *http.Request)  {
	SetCsrf(r)
	td.Form = forms.NewForm(nil)
	render.Template(rw, "signup.page.tmpl", td)
}
func (repo *Repository) PostSignup(rw http.ResponseWriter, r *http.Request)  {
	var u user.User
	err := r.ParseForm()
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}
	fn := r.Form.Get("first_name")
	ln := r.Form.Get("last_name")
	email := r.Form.Get("email")
	pass := r.Form.Get("password")
	f := forms.NewForm(r.PostForm)
	f.Required("first_name", "last_name", "email", "password")
	f.IsEmail("email")
	if !f.Valid() {
		data := make(map[string]interface{})
		data["first_name"] = fn
		data["last_name"] = ln
		data["email"] = email
		data["password"] = pass
		td.Data = data
		td.Form = f
		td.Error = "invalid inputs"
		render.Template(rw, "signup.page.tmpl", td)
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	u.FirstName = fn
	u.LastName = ln
	u.Email = email
	u.Password = string(hashedPass)
	u.AccessLevel = 1
	id, err := us.Add(u)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}
	repo.App.Session.Put(r.Context(), "user_id", id)
	td.Flash = "thanks for siging up"
	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

func (repo *Repository) Logut(rw http.ResponseWriter, r *http.Request)  {
	_= repo.App.Session.Destroy(r.Context())
	_= repo.App.Session.RenewToken(r.Context())
	http.Redirect(rw, r, "/login", http.StatusSeeOther)
}
