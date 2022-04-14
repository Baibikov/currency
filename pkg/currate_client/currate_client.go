package currate_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type Pair struct {
	From string
	To   string

	Well float64
}

func (p *Pair) String() string {
	return strings.ToUpper(p.From+p.To)
}

type Pairs []Pair

func (p Pairs) String() string {
	pairs := make([]string, 0, len(p))
	for _, v := range p {
		pairs = append(pairs, v.String())
	}

	return strings.Join(pairs, ",")
}


type Config struct {
	Key string
}

type Client struct {
	config Config
}

func New(config Config) *Client {
	return &Client{
		config: config,
	}
}

type Response struct {
	Status int
	Message string
	Data   Pairs
}

func (c *Client) url(pairs string) string {
	return fmt.Sprintf("https://currate.ru/api/?key=%s&pairs=%s&get=rates", c.config.Key, pairs)
}

func (c *Client) GetRates(pairs Pairs) (*Response, error) {
	resp, err := http.Get(c.url(pairs.String()))
	if err != nil {
		return nil, err
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer multierr.AppendInto(&err, resp.Body.Close())

	type response struct {
		Data map[string]string `json:"data"`
		Status int `json:"status"`
		Message string `json:"message"`
	}

	var vv response
	err = json.Unmarshal(bb, &vv)
	if err != nil {
		return nil, err
	}

	curResp := &Response{
		Status: vv.Status,
		Message: vv.Message,
		Data: make(Pairs, 0, len(pairs)),
	}

	check := make(map[string]struct{})

	for _, p := range pairs {
		for k, v := range vv.Data {
			rate := p.From+p.To
			_, ok := check[rate]

			if strings.Contains(k, p.From) && strings.Contains(k, p.To) && !ok {
				well, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, errors.Wrapf(err, "convert to float64 %s", v)
				}

				curResp.Data = append(curResp.Data, Pair{
					From: p.From,
					To: p.To,

					Well: well,
				})

				check[rate] = struct{}{}
			}
		}
	}

	return curResp, nil
}