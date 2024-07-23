package Controller

import (
	"WebMidterm/Interface"
	"WebMidterm/Model"
	"WebMidterm/Repository"
	_ "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BasketController struct {
	Repo Interface.Repository
}

func (h BasketController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)
	if r.Method == http.MethodGet {
		if r.URL.Path == "/basket" || r.URL.Path == "/basket/" {
			all, err := h.Repo.All()
			if err != nil {
				http.Error(w, "Cannot retrieve baskets from database", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err4 := json.NewEncoder(w).Encode(all)
			if err4 != nil {
				return
			}
		} else {
			id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/basket/"), 10, 64)
			if err != nil {
				http.Error(w, "wrong id", http.StatusBadRequest)
				return
			}

			basket, err := h.Repo.GetById(id)
			if errors.Is(err, Repository.ErrNotExists) {
				http.Error(w, "basket not found", http.StatusNotFound)
				return
			}
			if err != nil {
				http.Error(w, "Cannot return basket", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err3 := json.NewEncoder(w).Encode(basket)
			if err3 != nil {
				return
			}
		}
	}
	if r.Method == http.MethodPost {
		var raw map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&raw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, ok1 := raw["data"].(string)
		status, ok2 := raw["status"].(string)

		if !ok1 || !ok2 {
			http.Error(w, "wrong keys", http.StatusBadRequest)
			return
		}

		temp := Model.Basket{ID: 0, CreatedAt: time.Now(), UpdatedAt: time.Now(), Data: data, Status: status}
		created, err := h.Repo.Create(temp)
		if err != nil {
			http.Error(w, "Cannot create basket", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err2 := json.NewEncoder(w).Encode(created)
		if err2 != nil {
			return
		}
	}

	if r.Method == http.MethodPatch {
		id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/basket/"), 10, 64)
		if err != nil {
			http.Error(w, "wrong id", http.StatusBadRequest)
			return
		}

		previous, err := h.Repo.GetById(id)
		if errors.Is(err, Repository.ErrNotExists) {
			http.Error(w, "basket not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "Cannot update basket", http.StatusInternalServerError)
			return
		}

		var raw map[string]interface{}
		err = json.NewDecoder(r.Body).Decode(&raw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, ok1 := raw["data"].(string)
		status, ok2 := raw["status"].(string)

		if !ok1 || !ok2 {
			http.Error(w, "wrong keys", http.StatusBadRequest)
			return
		}

		updated := Model.Basket{ID: previous.ID, CreatedAt: previous.CreatedAt, UpdatedAt: time.Now(), Data: data, Status: status}
		if updated.Status == "COMPLETED" {
			http.Error(w, "basket status is completed", http.StatusBadRequest)
			return
		}

		upd, err := h.Repo.Update(id, updated)
		if err != nil {
			http.Error(w, "Cannot update basket", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err1 := json.NewEncoder(w).Encode(upd)
		if err1 != nil {
			return
		}
	}
	if r.Method == http.MethodDelete {
		id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/basket/"), 10, 64)
		if err != nil {
			http.Error(w, "wrong id", http.StatusBadRequest)
			return
		}

		err = h.Repo.Delete(id)
		if errors.Is(err, Repository.ErrDeleteFailed) {
			http.Error(w, "No basket with this id exists", http.StatusInternalServerError)
			return
		} else if err != nil {
			http.Error(w, "Cannot delete basket", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
	fmt.Println(h.Repo.All())
}

// h := Controller.BasketController{Repo: basketRepository}
/*
func GetBaskets (w http.ResponseWriter, r *http.Request) {
	all, err := h.Repo.All()
	if err != nil {
		http.Error(w, "Cannot retrieve baskets from database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err4 := json.NewEncoder(w).Encode(all)
	if err4 != nil {
		return
	}

func GetBasket (w http.ResponseWriter, r *http.Request){
	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/basket/"), 10, 64)
	if err != nil {
		http.Error(w, "wrong id", http.StatusBadRequest)
		return
	}

	basket, err := h.Repo.GetById(id)
	if errors.Is(err, Repository.ErrNotExists) {
		http.Error(w, "basket not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Cannot return basket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err3 := json.NewEncoder(w).Encode(basket)
	if err3 != nil {
		return
	}
}

func CreateBasket (w http.ResponseWriter, r *http.Request){}
var raw map[string]interface{}
err := json.NewDecoder(r.Body).Decode(&raw)
if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

data, ok1 := raw["data"].(string)
status, ok2 := raw["status"].(string)

if !ok1 || !ok2 {
http.Error(w, "wrong keys", http.StatusBadRequest)
return
}

temp := Model.Basket{ID: 0, CreatedAt: time.Now(), UpdatedAt: time.Now(), Data: data, Status: status}
created, err := h.Repo.Create(temp)
if err != nil {
http.Error(w, "Cannot create basket", http.StatusInternalServerError)
return
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
err2 := json.NewEncoder(w).Encode(created)
if err2 != nil {
return
}

func UpdateBasket (w http.ResponseWriter, r *http.Request){
	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/basket/"), 10, 64)
	if err != nil {
		http.Error(w, "wrong id", http.StatusBadRequest)
		return
	}

	previous, err := h.Repo.GetById(id)
	if errors.Is(err, Repository.ErrNotExists) {
		http.Error(w, "basket not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Cannot update basket", http.StatusInternalServerError)
		return
	}

	var raw map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, ok1 := raw["data"].(string)
	status, ok2 := raw["status"].(string)

	if !ok1 || !ok2 {
		http.Error(w, "wrong keys", http.StatusBadRequest)
		return
	}

	updated := Model.Basket{ID: previous.ID, CreatedAt: previous.CreatedAt, UpdatedAt: time.Now(), Data: data, Status: status}
	if updated.Status == "COMPLETED" {
		http.Error(w, "basket status is completed", http.StatusBadRequest)
		return
	}

	upd, err := h.Repo.Update(id, updated)
	if err != nil {
		http.Error(w, "Cannot update basket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err1 := json.NewEncoder(w).Encode(upd)
	if err1 != nil {
		return
	}
}
func DeleteBasket (w http.ResponseWriter, r *http.Request){
	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/basket/"), 10, 64)
	if err != nil {
		http.Error(w, "wrong id", http.StatusBadRequest)
		return
	}

	err = h.Repo.Delete(id)
	if errors.Is(err, Repository.ErrDeleteFailed) {
		http.Error(w, "No basket with this id exists", http.StatusInternalServerError)
		return
	} else if err != nil {
		http.Error(w, "Cannot delete basket", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
*/
