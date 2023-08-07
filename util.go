package deezus

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/oklookat/deezus/schema"
	"github.com/oklookat/vantuz"
)

func genApiPath(paths ...string) string {
	if len(paths) == 0 {
		return schema.ApiUrl
	}

	base := schema.ApiUrl + "/" + paths[0]
	for i := 1; i < len(paths); i++ {
		if len(paths[i]) == 0 {
			continue
		}
		base += "/" + paths[i]
	}

	return base
}

const _errPrefix = "deezus"

func checkResponse(resp *vantuz.Response, data schema.ErrorInResponse) error {
	if data.Error != nil {
		return fmt.Errorf("%s: %w", _errPrefix, *data.Error)
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("%s: %d", _errPrefix, resp.StatusCode)
	}
	return nil
}

func getIndexLimit(index, limit int) url.Values {
	params := url.Values{}
	if index < 1 {
		index = 0
	}
	if limit < 1 {
		limit = 1
	}
	if index > 0 {
		params.Set("index", strconv.Itoa(index))
	}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	return params
}

// 1,2,3,4.
func idsJoin(ids []schema.ID) string {
	conv := make([]string, len(ids))
	for i := range conv {
		conv[i] = ids[i].String()
	}
	return strings.Join(conv, ",")
}

func setSearchParams(params url.Values, query string, order schema.Order, strict bool) {
	params.Set("q", query)
	if len(order) > 0 {
		params.Set("order", order.String())
	}
	if strict {
		params.Set("strict", "on")
	}
}

func getAnyResp[T any](ctx context.Context, client *Client, queryParams url.Values, paths ...string) (*schema.Response[T], error) {
	data := &schema.Response[T]{}
	resp, err := httpAny(ctx, client, queryParams, data, nil, http.MethodGet, paths...)
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

func postAny(ctx context.Context, client *Client, body url.Values, paths ...string) (*schema.BoolResponse, error) {
	data := &schema.BoolResponse{}
	resp, err := httpAny(ctx, client, nil, nil, body, http.MethodPost, paths...)
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

func postAnyParams(ctx context.Context, client *Client, params url.Values, paths ...string) (*schema.BoolResponse, error) {
	data := &schema.BoolResponse{}
	resp, err := httpAny(ctx, client, params, data, nil, http.MethodPost, paths...)
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

func deleteAny(ctx context.Context, client *Client, queryParams url.Values, paths ...string) (*schema.BoolResponse, error) {
	data := &schema.BoolResponse{}
	resp, err := httpAny(ctx, client, queryParams, data, nil, http.MethodDelete, paths...)
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

func httpAny(
	ctx context.Context,
	client *Client,
	queryParams url.Values,
	resultAndError any,
	body url.Values,
	method string,
	paths ...string,
) (*vantuz.Response, error) {
	endpoint := genApiPath(paths...)

	req := client.Http.R()
	if resultAndError != nil {
		req.SetResult(resultAndError).
			SetError(resultAndError)
	}
	if body != nil {
		req.SetFormUrlValues(body)
	}

	for k, v := range queryParams {
		for _, v2 := range v {
			req.QueryParams().Add(k, v2)
		}
	}

	var (
		resp *vantuz.Response
		err  error
	)

	switch method {
	case http.MethodGet:
		resp, err = req.Get(ctx, endpoint)
	case http.MethodPost:
		resp, err = req.Post(ctx, endpoint)
	case http.MethodPut:
		resp, err = req.Put(ctx, endpoint)
	case http.MethodDelete:
		resp, err = req.Delete(ctx, endpoint)
	}

	return resp, err
}
