package ctg_medsenger_bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"time"
)

type (
	CtgCredentials struct {
		Id  string `json:"id"`
		Key string `json:"key"`
	}

	tokenResponse struct {
		Success bool    `json:"success"`
		Token   *string `json:"token"`
		Message *string `json:"message"`
	}

	releaseTokenResponse struct {
		Success bool    `json:"success"`
		Message *string `json:"message"`
	}

	token struct {
		token     string
		createdAt time.Time
	}

	CtgRecord struct {
		Id          uuid.UUID  `json:"uuid"`
		PatientId   string     `json:"id"`
		CreatedAt   time.Time  `json:"date"`
		ReceivedAt  time.Time  `json:"received"`
		DiagnosedAt *time.Time `json:"diagnosed"`
	}

	recordsResponse struct {
		Success bool        `json:"success"`
		Message *string     `json:"message"`
		Result  []CtgRecord `json:"result"`
	}
)

type CtgClient struct {
	host        string
	credentials *CtgCredentials
	token       token
}

func NewCtgClient(host string, credentials *CtgCredentials) *CtgClient {
	return &CtgClient{host: host, credentials: credentials}
}

func (c *CtgClient) urlAppendingPath(path string) *url.URL {
	return &url.URL{Scheme: "https", Host: c.host, Path: path}
}

func (c *CtgClient) fetchToken() (*string, error) {
	reqUrl := c.urlAppendingPath("/api/login")
	encodedData, err := json.Marshal(c.credentials)
	if err != nil {
		return nil, err
	}
	httpResponse, err := http.Post(reqUrl.String(), "application/json", bytes.NewBuffer(encodedData))
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetToken: response status code is not OK: %s", httpResponse.Status)
	}
	defer httpResponse.Body.Close()
	var response tokenResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, fmt.Errorf("GetToken: response is not successful: %s", *response.Message)
	}
	return response.Token, nil
}

func (c *CtgClient) getToken() (*string, error) {
	now := time.Now()

	// token lives 60 minutes
	if c.token.token == "" || now.Sub(c.token.createdAt) > 55*time.Minute {
		tokenStr, err := c.fetchToken()
		if err != nil {
			return nil, err
		}
		c.token = token{token: *tokenStr, createdAt: now}
	}

	return &c.token.token, nil
}

func (c *CtgClient) releaseToken() error {
	if c.token.token == "" {
		return errors.New("ReleaseToken: token is empty")
	}
	reqUrl := c.urlAppendingPath("/api/login")
	req, err := http.NewRequest(http.MethodPost, reqUrl.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-access-token", c.token.token)
	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("ReleaseToken: response status code is not OK: %s", httpResponse.Status)
	}
	defer httpResponse.Body.Close()
	var response releaseTokenResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return err
	}
	if !response.Success {
		return fmt.Errorf("ReleaseToken: response is not successful: %s", *response.Message)
	}
	return nil
}

func (c *CtgClient) GetRecordsList(from time.Time, to time.Time) ([]CtgRecord, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	reqUrl := c.urlAppendingPath("/api/list")
	q := reqUrl.Query()
	q.Set("begin", from.Format(time.RFC3339))
	q.Set("end", to.Format(time.RFC3339))
	reqUrl.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-access-token", *token)
	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetRecordsList: response status code is not OK: %s", httpResponse.Status)
	}
	defer httpResponse.Body.Close()
	var response recordsResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, fmt.Errorf("GetRecordsList: response is not successful: %s", *response.Message)
	}
	return response.Result, nil
}

func (c *CtgClient) GetRecordPDF(id uuid.UUID) ([]byte, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	reqUrl := c.urlAppendingPath("/api/pdf")
	q := reqUrl.Query()
	q.Set("uuid", id.String())
	reqUrl.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-access-token", *token)
	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	return io.ReadAll(httpResponse.Body)
}
