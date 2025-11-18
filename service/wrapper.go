package service

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IHttpWrapper interface {
	Get(url string, header map[string]string) ([]byte, error)
	Post(url string, body []byte, header map[string]string) ([]byte, error)
	Put(url string, body []byte, header map[string]string) ([]byte, error)
	Delete(url string, header map[string]string) ([]byte, error)
}

type HttpWrapper struct {
	httpClient *http.Client
}

func NewHttpWrapper(httpClient *http.Client) IHttpWrapper {
	return &HttpWrapper{
		httpClient: httpClient,
	}
}

func (h *HttpWrapper) buildGetRequest(url string, header map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("client: could not create GET request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (h *HttpWrapper) Get(url string, header map[string]string) ([]byte, error) {
	req, err := h.buildGetRequest(url, header)
	if err != nil {
		return nil, fmt.Errorf("client: failed to create GET request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client: failed to perform GET request: %w", err)
	}
	defer resp.Body.Close()

	if err := handleStatusCode(resp); err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("client: failed to read response body: %w", err)
	}

	return respBody, nil
}

func (h *HttpWrapper) buildPostRequest(url string, body []byte, header map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("client: could not create POST request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (h *HttpWrapper) Post(url string, body []byte, header map[string]string) ([]byte, error) {
	req, err := h.buildPostRequest(url, body, header)
	if err != nil {
		return nil, fmt.Errorf("client: could not create request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client: error making http request: %w", err)
	}
	defer resp.Body.Close()

	if err := handleStatusCode(resp); err != nil {
		return nil, err
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body: %w", err)
	}

	return bodyResp, nil
}

func (h *HttpWrapper) buildPutRequest(url string, body []byte, header map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("client: could not create PUT request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (h *HttpWrapper) Put(url string, body []byte, header map[string]string) ([]byte, error) {
	req, err := h.buildPutRequest(url, body, header)
	if err != nil {
		return nil, fmt.Errorf("failed to create PUT request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send PUT request: %w", err)
	}
	defer resp.Body.Close()

	if err := handleStatusCode(resp); err != nil {
		return nil, err
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body: %w", err)
	}

	return bodyResp, nil
}

func (h *HttpWrapper) buildDeleteRequest(url string, header map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("client: could not create DELETE request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (h *HttpWrapper) Delete(url string, header map[string]string) ([]byte, error) {
	req, err := h.buildDeleteRequest(url, header)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send DELETE request: %w", err)
	}
	defer resp.Body.Close()

	if err := handleStatusCode(resp); err != nil {
		return nil, err
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body: %w", err)
	}

	return bodyResp, nil
}
