package client

import (
	"encoding/csv"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/abdulmoeid7112/impact-analysis-api-svc/config"
)

// DatasetApiClient - http client for dataset api
type DatasetApiClient struct {
	Client *http.Client
}

// NewDatasetApiClient - initializes new client
func NewDatasetApiClient() *DatasetApiClient {
	return &DatasetApiClient{
		Client: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{Timeout: 30 * time.Second}).DialContext,
			},
		},
	}
}

// GetDataSet - fetch dataset for analysis
func (c *DatasetApiClient) GetDataSet() ([][]string, error) {
	requestURL, err := url.Parse(viper.GetString(config.DatasetApiURL))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse api url")
	}

	resp, err := c.Client.Do(&http.Request{
		Method: http.MethodGet,
		URL:    requestURL,
	})
	if err != nil {
		log().Errorf("failed to fetch dataset, err: %+v", err)
		return nil, errors.Wrap(err, "failed to fetch dataset")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.Wrap(err, "failed to get valid response")
	}

	reader := csv.NewReader(resp.Body)
	reader.Comma = ','

	dataset, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return dataset, nil
}
