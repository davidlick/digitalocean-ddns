package digitalocean

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type client struct {
	baseUrl    string
	domain     string
	token      string
	httpClient *http.Client
}

func NewClient(baseUrl, token, domain string) *client {
	return &client{
		baseUrl: baseUrl,
		domain:  domain,
		token:   token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *client) GetARecordID() (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v2/domains/%s/records", c.baseUrl, c.domain), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	var records RecordsResponse
	err = c.doRequest(req, &records)
	if err != nil {
		return "", err
	}

	for _, record := range records.DomainRecords {
		if record.Name == "vpn" && record.Type == "A" {
			return strconv.Itoa(int(record.Id)), nil
		}
	}

	return "", ErrNotFound
}

func (c *client) SetARecord(id, ip string) error {
	body := fmt.Sprintf(`{"data":%q}`, strings.ReplaceAll(ip, "\n", ""))
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v2/domains/%s/records/%s", c.baseUrl, c.domain, id), bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Content-Type", "application/json")

	var resp map[string]interface{}
	err = c.doRequest(req, &resp)
	return err
}

func (c *client) doRequest(req *http.Request, output interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusInternalServerError {
		return fmt.Errorf("%w: statuscode %d", ErrServerError, resp.StatusCode)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("%w: statuscode %d", ErrClientError, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	return err
}
