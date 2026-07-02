package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type YaSuite struct {
	suite.Suite
	client *http.Client
}

func (s *YaSuite) SetupSuite() {
	s.client = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func (s *YaSuite) TearDownSuite() {
	// если ресурсов нет, можно оставить пустым
}

func (s *YaSuite) TestYaRuReachable() {
	req, err := http.NewRequest(http.MethodGet, "https://ya.ru", nil)
	require.NoError(s.T(), err)

	resp, err := s.client.Do(req)
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	require.True(s.T(),
		resp.StatusCode == http.StatusOK ||
			resp.StatusCode == http.StatusMovedPermanently ||
			resp.StatusCode == http.StatusFound,
		"unexpected status code: %d",
		resp.StatusCode,
	)
}

func (s *YaSuite) TestYaRuResponseTime() {
	start := time.Now()

	req, err := http.NewRequest(http.MethodGet, "https://ya.ru", nil)
	require.NoError(s.T(), err)

	resp, err := s.client.Do(req)
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	elapsed := time.Since(start)

	require.True(s.T(),
		resp.StatusCode == http.StatusOK ||
			resp.StatusCode == http.StatusMovedPermanently ||
			resp.StatusCode == http.StatusFound,
		"unexpected status code: %d",
		resp.StatusCode,
	)

	require.Less(s.T(), elapsed, 5*time.Second, "response too slow: %s", elapsed)
}

func TestYaSuite(t *testing.T) {
	suite.Run(t, new(YaSuite))
}
