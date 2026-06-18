package grafana

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	token   string
	http    *http.Client
}

type Dashboard struct {
	UID  string
	Name string
}

var Dashboards = []Dashboard{
	{UID: "fcstask-http", Name: "🌐 HTTP"},
	{UID: "fcstask-postgres", Name: "🐘 Postgres"},
	{UID: "fcstask-go-runtime", Name: "⚙️ Go Runtime"},
	{UID: "fcstask-business", Name: "💼 Business"},
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		baseURL: baseURL,
		token:   token,
		http:    &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) RenderDashboard(ctx context.Context, uid string) ([]byte, error) {
	url := fmt.Sprintf("%s/render/d/%s?orgId=1&from=now-1h&to=now&width=1200&height=800&theme=dark", c.baseURL, uid)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("grafana render returned %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
