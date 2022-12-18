package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/titoyudha/go_blog_api/api/auth"
	"github.com/titoyudha/go_blog_api/api/model"
	"github.com/titoyudha/go_blog_api/api/responses"
	"github.com/titoyudha/go_blog_api/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.Signin(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatterError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) Signin(email, password string) (string, error) {
	var err error

	user := model.User{}

	err = server.DB.Debug().Model(model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return " ", err
	}

	err = model.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(int32(user.ID))
}
