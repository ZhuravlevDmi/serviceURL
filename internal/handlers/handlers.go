package handlers

import (
	"encoding/json"
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"io"
	"log"
	"net/http"
)

func HandlerGetURL(MapURL storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			смотрим путь в request, если этот путь(path) есть в словаре MapURL в виде ключа,
			то редиректим пользователя, на MapURL[path], если path в словаре нет, возвращаем 400
		*/

		path := r.URL.Path[1:]
		redirectPath := storage.ReadStorage(MapURL, path)

		if redirectPath == "" {
			http.Error(w, "path is empty", http.StatusBadRequest)
			return
		} else {
			//w.WriteHeader(http.StatusTemporaryRedirect)
			w.Header().Set("Location", config.ServerURL+"/"+path)
			http.Redirect(w, r, redirectPath, http.StatusTemporaryRedirect)
		}
	}
}

func HandlerPostURL(MapURL storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			генерируем мини-урл из 6 символов и записываем новое значение в mapURL(ключ - мини-урл,
			значение body из пост запроса)
		*/
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		// обрабатываем ошибку
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := storage.RecordStorage(MapURL, string(b))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(config.ServerURL + "/" + resp))
			return
		}
		// пишем тело ответа
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(config.ServerURL + "/" + resp))
	}
}

type URL struct {
	URL string `json:"url,omitempty"`
}
type ResultURL struct {
	Result string `json:"result"`
}
type ErrorRequest struct {
	Error string `json:"error"`
}

func HandlerAPIShorten(MapURL storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v URL         // целевой объект
		var res ResultURL // целевой объект
		var e ErrorRequest
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		req, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(req, &v); err != nil || v.URL == "" {
			e.Error = "Bad Request"
			errResponse, _ := json.Marshal(e)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errResponse)
			return
		}

		resp, err := storage.RecordStorage(MapURL, v.URL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)

		res.Result = config.ServerURL + "/" + resp
		response, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}
		w.Write(response)
	}
}
