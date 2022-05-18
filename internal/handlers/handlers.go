package handlers

import (
	"encoding/json"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"io"
	"log"
	"net/http"
)

func HandlerGetURL(MapURL storage.Storage, ServerURL string) http.HandlerFunc {
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
			w.Header().Set("Location", ServerURL+"/"+path)
			http.Redirect(w, r, redirectPath, http.StatusTemporaryRedirect)
		}
	}
}

func HandlerPostURL(MapURL storage.Storage, ServerURL string, f storage.FileWorkStruct, path string) http.HandlerFunc {
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
			_, err := w.Write([]byte(ServerURL + "/" + resp))
			if err != nil {
				log.Println(err)
			}
			return
		}

		f.OpenFileWrite(path)

		defer f.Close()
		s := storage.URLStorageStruct{
			ID:  resp,
			URL: string(b),
		}
		f.WriteFile(s)

		w.WriteHeader(http.StatusCreated)

		w.WriteHeader(http.StatusCreated)
		_, e := w.Write([]byte(ServerURL + "/" + resp))
		if e != nil {
			log.Println(e)
		}
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

func HandlerAPIShorten(MapURL storage.Storage, ServerURL string, f storage.FileWorkStruct, path string) http.HandlerFunc {
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
			_, err := w.Write(errResponse)
			if err != nil {
				log.Println(err)
			}
			return
		}

		resp, err := storage.RecordStorage(MapURL, v.URL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		f.OpenFileWrite(path)

		defer f.Close()
		s := storage.URLStorageStruct{
			ID:  resp,
			URL: v.URL,
		}
		f.WriteFile(s)

		w.WriteHeader(http.StatusCreated)

		res.Result = ServerURL + "/" + resp
		response, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}
		_, er := w.Write(response)
		if er != nil {
			log.Println(er)
		}
	}
}
