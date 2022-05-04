package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := model.FindUsers()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utility.RespondJSON(w, http.StatusOK, users)
	}
}

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		validate := validator.New()

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := model.CreateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// token := auth0.GetJWT(c.Request().Context())
		// fmt.Printf("jwt %+v\n", token)

		// if token == nil {
		// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": "missing or malformed jwt"})
		// }
		// // token.Claimsをjwt.MapClaimsへ変換
		// claims := token.Claims.(jwt.MapClaims)
		// // claimsの中にペイロードの情報が入っている
		// sub := claims["sub"].(string)
		// fmt.Println(sub)
		// var user model.User

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := model.FindUser(&model.User{ID: uint(id)})

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}

func UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data model.User
		validate := validator.New()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := model.UpdateUser(&model.User{ID: uint(id)}, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := model.DeleteUser(&model.User{ID: uint(id)})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}
