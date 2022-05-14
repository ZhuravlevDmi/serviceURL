package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testMapURL storage.Storage = &storage.StorageMapURL{MapURL: map[string]string{
	"fgt56f": "https://yandex.ru",
	"4dgtd5": "https://google.com",
}}

func TestHandlerGetURL(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		path string
		want want
	}{
		{
			name: "test statusCode 200",
			path: "4dgtd5",
			want: want{
				statusCode: 307,
			},
		},
		{
			name: "test statusCode 400",
			path: "3ddd34",
			want: want{
				statusCode: 400,
			},
		},
		{
			name: "test statusCode 400 v2",
			path: "",
			want: want{
				statusCode: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.path != "" {
				tt.path = "/" + tt.path
			} else {
				tt.path = "/"
			}
			request := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(HandlerGetURL(testMapURL))

			h.ServeHTTP(w, request)
			result := w.Result()

			defer result.Body.Close()
			if result.StatusCode != tt.want.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.want.statusCode, w.Code)
			}

		})
	}
}

func TestHandlerPostURL(t *testing.T) {
	type want struct {
		statusCode  int
		lenResponse int
	}
	tests := []struct {
		name string
		body string
		want want
	}{
		{
			name: "test statusCode201",
			body: "https://vk.com",
			want: want{
				statusCode:  201,
				lenResponse: len(config.ServerURL) + 7,
			},
		},
		{
			name: "test statusCode400",
			body: "https://yandex.ru",
			want: want{
				statusCode:  400,
				lenResponse: len(config.ServerURL) + 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(tt.body)))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(HandlerPostURL(testMapURL))

			h.ServeHTTP(w, request)
			result := w.Result()

			if result.StatusCode != tt.want.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.want.statusCode, w.Code)
			}
			// получаем и проверяем тело запроса
			defer result.Body.Close()
			resBody, err := io.ReadAll(result.Body)
			if err != nil {
				t.Fatal(err)
			}
			if len([]rune(string(resBody))) != tt.want.lenResponse {
				t.Errorf("Expected body %d, got %d", tt.want.lenResponse, len([]rune(string(resBody))))
			}

		})
	}
}

func TestHandlerAPIShorten(t *testing.T) {
	type URLRequest struct {
		URL string `json:"url"`
	}

	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		body URLRequest
		want want
	}{
		{
			name: "test statusCode201",
			body: URLRequest{URL: "https://vk.ru"},
			want: want{
				statusCode: 201,
			},
		},
		{
			name: "test statusCode400",
			body: URLRequest{URL: "https://vk.com"},
			want: want{
				statusCode: 400,
			},
		},
		//{
		//	name: "test statusCode400",
		//	body: ErrorUrlRequest{Urls: "https://yandex.ru"},
		//	want: want{
		//		statusCode: 201,
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, _ := json.Marshal(tt.body)

			request := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(HandlerAPIShorten(testMapURL))

			h.ServeHTTP(w, request)
			result := w.Result()
			defer result.Body.Close()
			if result.StatusCode != tt.want.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.want.statusCode, w.Code)
			}
			// получаем и проверяем тело запроса
			//defer result.Body.Close()
			//resBody, err := io.ReadAll(result.Body)
			//if err != nil {
			//	t.Fatal(err)
			//}
			//if len([]rune(string(resBody))) != tt.want.lenResponse {
			//	t.Errorf("Expected body %d, got %d", tt.want.lenResponse, len([]rune(string(resBody))))
			//}

		})
	}
}
