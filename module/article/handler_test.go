package article

import (
	"fmt"
	"testing"

	"github.com/test_cache_CQRS/module/primitive"

	"github.com/go-resty/resty/v2"
)

func TestAuth(t *testing.T) {
	t.Run("TestArticleRecordSaveSuccess", func(t *testing.T) {
		payload := primitive.ArticleReq{
			Author: "new author",
			Title:  "title",
			Body:   "body",
		}
		client := resty.New()
		resp, err := client.R().
			SetBody(payload).
			Post("http://localhost:1234/api/v1/articles")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 200 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
		}
	})

	t.Run("TestArticleRecordSaveFailed", func(t *testing.T) {
		payload := primitive.ArticleReq{
			Author: "new author",
			Title:  "title",
			Body:   "",
		}
		client := resty.New()
		resp, err := client.R().
			SetBody(payload).
			Post("http://localhost:1234/api/v1/articles")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 400 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 400, resp.StatusCode())
		}
	})

	t.Run("TestGetArticlesSuccess", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			Get("http://localhost:1234/api/v1/articles")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 200 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
		}
	})

	t.Run("TestGetDetailArticlesSuccess", func(t *testing.T) {
		id := 7
		url := fmt.Sprintf("%s/%d", "http://localhost:1234/api/v1/articles", id)
		client := resty.New()
		resp, err := client.R().
			Get(url)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 200 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
		}
	})

	t.Run("TestGetDetailArticlesFailedNotFound", func(t *testing.T) {
		id := 9999999
		url := fmt.Sprintf("%s/%d", "http://localhost:1234/api/v1/articles", id)
		client := resty.New()
		resp, err := client.R().
			Get(url)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 404 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 404, resp.StatusCode())
		}
	})

	t.Run("TestGetDetailArticlesFailedZeroId", func(t *testing.T) {
		id := 0
		url := fmt.Sprintf("%s/%d", "http://localhost:1234/api/v1/articles", id)
		client := resty.New()
		resp, err := client.R().
			Get(url)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 404 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 404, resp.StatusCode())
		}
	})

}
