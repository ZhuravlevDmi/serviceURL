package handlers

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"io"
	"net/http"
)

func MainHandler(MapUrl storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// принимает запрос и перенаправляет на другой хендлер в зависимости от типа запроса
		if r.Method == http.MethodPost {
			h := http.HandlerFunc(HandlerPostURL(MapUrl))
			h.ServeHTTP(w, r)
			return

		} else if r.Method == http.MethodGet {
			h := http.HandlerFunc(HandlerGetURL(MapUrl))
			h.ServeHTTP(w, r)
			return

		} else {
			h := http.HandlerFunc(HandlerURL)
			h.ServeHTTP(w, r)
			return
		}
	}
}

func HandlerGetURL(MapUrl storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			смотрим путь в request, если этот путь(path) есть в словаре MapURL в виде ключа,
			то редиректим пользователя, на MapURL[path], если path в словаре нет, возвращаем 400
		*/

		path := r.URL.Path[1:]
		//redirectPath := MapUrl.Read(path)
		redirectPath := storage.ReadStorage(MapUrl, path)

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

func HandlerPostURL(MapUrl storage.Storage) http.HandlerFunc {
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

		//resp, err := MapUrl.Record(string(b))
		resp, err := storage.RecordStorage(MapUrl, string(b))
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

func HandlerURL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Обрабатываются только GET или POST запрос"))

}
