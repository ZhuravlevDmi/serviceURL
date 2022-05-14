package handlers

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"io"
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


