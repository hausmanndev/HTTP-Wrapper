package service

import (
	"fmt"
	"net/http"
)

func handleStatusCode(resp *http.Response) error {
	switch resp.StatusCode {
	//        200               201                 204
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return nil

	case http.StatusBadRequest:
		return fmt.Errorf("client: bad request (400)")

	case http.StatusUnauthorized:
		return fmt.Errorf("client: unauthorized (401)")

	case http.StatusForbidden:
		return fmt.Errorf("client: forbidden (403)")

	case http.StatusNotFound:
		return fmt.Errorf("client: not found (404)")

	case http.StatusTooManyRequests:
		return fmt.Errorf("client: too many requests (429)")

	case http.StatusServiceUnavailable:
		return fmt.Errorf("client: service unavailable (503)")

	default:
		return fmt.Errorf("client: unexpected status code: %d", resp.StatusCode)
	}
}
