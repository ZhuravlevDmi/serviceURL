package handlers

import (
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/util"
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// принимает запрос и перенаправляет на другой хендлер в зависимости от типа запроса
	if r.Method == http.MethodPost {
		h := http.HandlerFunc(HandlerPostURL(util.MapUrl))
		h.ServeHTTP(w, r)
		return

	} else if r.Method == http.MethodGet {
		h := http.HandlerFunc(HandlerGetURL(util.MapUrl))
		h.ServeHTTP(w, r)
		return

	} else {
		h := http.HandlerFunc(HandlerURL)
		h.ServeHTTP(w, r)
		return
	}
}

func HandlerGetURL(MapUrl map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			смотрим путь в request, если этот путь(path) есть в словаре MapURL в виде ключа,
			то редиректим пользователя, на MapURL[path], если path в словаре нет, возвращаем 400
		*/

		path := r.URL.Path[1:]
		redirectPath := util.CheckMapUrl(MapUrl, path)

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
func HandlerPostURL(MapUrl map[string]string) http.HandlerFunc {
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

		resp, err := util.SetMapUrl(MapUrl, string(b))
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

//func MainHandler(w http.ResponseWriter, r *http.Request) {
//	/* принимиает POST или GET запрос
//	Если GET: смотрим путь в request, если этот путь(path) есть в словаре MapURL в виде ключа,
//	то редиректим пользователя, на MapURL[path], если path в словаре нет, возвращаем 400
//
//	Если POST: генерируем мини-урл из 6 символов и записываем новое значение в mapURL(ключ - мини-урл,
//	значение body из пост запроса)
//	*/
//	if r.Method == http.MethodPost {
//		log.Println("Метод Пост")
//
//		b, err := io.ReadAll(r.Body)
//		// обрабатываем ошибку
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		resp, err := util.SetMapUrl(util.MapUrl, string(b))
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			w.Write([]byte(resp))
//			//http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		// пишем тело ответа
//		w.WriteHeader(http.StatusCreated)
//		w.Write([]byte(resp))
//
//	} else if r.Method == http.MethodGet {
//		path := r.URL.Path[1:]
//		redirectPath := util.CheckMapUrl(util.MapUrl, path)
//
//		if redirectPath == "" {
//			http.Error(w, "path is empty", http.StatusBadRequest)
//			return
//		} else {
//			w.WriteHeader(http.StatusTemporaryRedirect)
//			http.Redirect(w, r, redirectPath, http.StatusTemporaryRedirect)
//		}
//
//	} else {
//		log.Println("Такое не обрабатываем")
//		w.WriteHeader(http.StatusBadRequest)
//		http.Error(w, "Только POST или GET запрос", http.StatusBadRequest)
//	}
//}
