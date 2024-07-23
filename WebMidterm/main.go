package main

import (
	"WebMidterm/Controller"
	"WebMidterm/Interface"
	"WebMidterm/Repository"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

/*
func Db() Controller.BasketController {
	db, err := sql.Open("sqlite3", "Basket.db")
	if err != nil {
		log.Fatal(err)
	}

	var basketRepository Interface.Repository
	basketRepository = Repository.NewSQLiteRepository(db)

	if err := basketRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	h := Controller.BasketController{Repo: basketRepository}
	return h
}
*/

type BasketController struct {
	Repo Interface.Repository
}

func main() {

	//r := mux.NewRouter()
	//var h = Migration.Db()
	db, err := sql.Open("sqlite3", "Basket.db")
	if err != nil {
		log.Fatal(err)
	}

	var basketRepository Interface.Repository
	basketRepository = Repository.NewSQLiteRepository(db)

	if err := basketRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	h := Controller.BasketController{Repo: basketRepository}
	//r.HandleFunc("/basket/", Controller.GetBaskets).Methods("GET")
	//r.HandleFunc("/basket/", Controller.CreateBasket).Methods("POST")
	//r.HandleFunc("/basket/{id}", Controller.GetBasket).Methods("GET")
	//r.HandleFunc("/basket/{id}", Controller.UpdateBasket).Methods("PATCH")
	//r.HandleFunc("/basket/{id}", Controller.DeleteBasket).Methods("DELETE")

	//http.Handle("/", r)
	http.Handle("/basket/", h)

	fmt.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", h))

	/*

		err = http.ListenAndServe(":8080", nil)
		fmt.Println(err)
	*/

}

/*
func (h BasketController) GetBaskets (w http.ResponseWriter, r *http.Request) {
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
}

func (h BasketController) GetBasket (w http.ResponseWriter, r *http.Request){
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

func (h BasketController) CreateBasket (w http.ResponseWriter, r *http.Request){
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


func (h BasketController) UpdateBasket (w http.ResponseWriter, r *http.Request){
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

func  (h BasketController) DeleteBasket (w http.ResponseWriter, r *http.Request) {
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
}*/
