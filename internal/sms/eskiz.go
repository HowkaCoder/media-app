package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const baseURL = "https://notify.eskiz.uz/api"

type Eskiz struct {
	email    string
	password string
	client   *http.Client

	mu    sync.RWMutex
	token string
}

func NewEskiz(email, password string) *Eskiz {
	return &Eskiz{
		email:    email,
		password: password,
		client:   &http.Client{Timeout: 10 * time.Second},
	}
}

type tokenResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func (e *Eskiz) Login(ctx context.Context) error {
	form := url.Values{}
	form.Set("email", e.email)
	form.Set("password", e.password)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/auth/login", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return e.updateToken(req)
}

func (e *Eskiz) Refresh(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, baseURL+"/auth/refresh", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+e.getToken())

	return e.updateToken(req)
}

func (e *Eskiz) StartRefresher(ctx context.Context, every time.Duration) {
	ticker := time.NewTicker(every)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := e.Refresh(ctx); err != nil {
					log.Printf("sms: refresh token: %v", err)
				}
			}
		}
	}()
}

func (e *Eskiz) Send(ctx context.Context, phone, text string) error {
	form := url.Values{}
	form.Set("mobile_phone", strings.TrimPrefix(phone, "+"))
	form.Set("message", text)
	form.Set("from", "4546")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/message/sms/send", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+e.getToken())

	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sms: send failed: %s: %s", resp.Status, body)
	}
	return nil
}

func (e *Eskiz) updateToken(req *http.Request) error {
	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sms: auth failed: %s: %s", resp.Status, body)
	}

	var tr tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return err
	}
	if tr.Data.Token == "" {
		return fmt.Errorf("sms: empty token in response")
	}

	e.mu.Lock()
	e.token = tr.Data.Token
	e.mu.Unlock()
	return nil
}

func (e *Eskiz) getToken() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.token
}
