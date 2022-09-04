package incomsms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://apiv2.incomsms.vn"
)

type Client struct {
	*httpClient

	Sms *smsService
}

// NewClient returns a Incom SMS client
func NewClient(userName string, password string) (*Client, error) {

	if userName == "" {
		return nil, errors.New("username is required")
	}

	if password == "" {
		return nil, errors.New("password is required")
	}

	c := &Client{
		httpClient: newHTTPClient(userName, password),
	}

	c.Sms = &smsService{client: c.httpClient}

	return c, nil
}

type httpClient struct {
	baseURL    *url.URL
	credential Credential
	client     *http.Client
	logger     func(...interface{})
}

func newHTTPClient(userName string, password string) *httpClient {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &httpClient{
		credential: Credential{
			Username: userName,
			Password: password,
		},
		baseURL: baseURL,
		client:  http.DefaultClient,
	}
}

// SetBaseURL change the default OneSignal base URL
func (c *httpClient) SetBaseURL(baseURL string) error {
	sBaseURL, err := url.Parse(baseURL)
	if err != nil {
		panic(fmt.Sprintf("incorrect base url format: %s", baseURL))
	}

	c.baseURL = sBaseURL
	return nil
}

// SetHTTPClient set custom http client
func (c *httpClient) SetHTTPClient(client *http.Client) {
	c.client = client
}

// SetLogger set custom debug logger
func (c *httpClient) SetLogger(logger func(args ...interface{})) {
	c.logger = logger
}

// NewRequest create an API request.
// path is a relative URL, like "/SendMultipleMessage_V4_get".
// The value pointed to by body is JSON encoded and included as the request body.
func (c *httpClient) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(c.baseURL.String() + path)
	if err != nil {
		return nil, err
	}

	c.printDebug(fmt.Sprintf("[IncomSMS] requesting url: %s", u.String()))

	var buf io.ReadWriter
	if body != nil {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(body)
		if err != nil {
			return nil, err
		}
		buf = b

		if c.logger != nil {
			c.logger("[IncomSMS] request body: " + b.String())
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// Sends an API request and returns the API response.
// Return JSON decoded and stored in the value pointed to by v,
// or an error if an API error has occurred.
func (c *httpClient) Do(r *http.Request, v interface{}) (*http.Response, error) {
	// send the request
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = checkErrorResponse(resp)
	if err != nil {
		return resp, err
	}

	if c.logger != nil {
		var b bytes.Buffer
		b.ReadFrom(resp.Body)
		c.printDebug("response body: ", b.String())
		err = json.Unmarshal(b.Bytes(), &v)
	} else {
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&v)
	}

	// it returns EOF if http status 204 (no content available)
	if err != nil && err != io.EOF {
		return resp, err
	}
	return resp, nil
}

func (c *httpClient) printDebug(args ...interface{}) {
	if c.logger != nil {
		c.logger(args...)
	}
}

// checkErrorResponse checks the API response for errors, by http status code
// and returns them if present
func checkErrorResponse(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK, http.StatusNoContent:
		return nil
	case http.StatusInternalServerError:
		return errors.New("internal server error")
	default:
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("couldn't read response body: %v", err)
		}
		return errors.New(string(b))
	}
}
