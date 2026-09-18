// Copyright © 2026. Citrix Systems, Inc.

package citrixclient_test

// Tests for GetAllPagesWithRetry in citrixdaas_client.go

import (
	"fmt"
	"net/http"
	"testing"

	citrixclient "github.com/citrix/citrix-daas-rest-go/client"
	"github.com/stretchr/testify/assert"
)

type mockItem struct {
	id string
}

// mockPage describes what a single Execute call should return for a given incoming token.
type mockPage struct {
	items     []mockItem
	next      string
	fail      bool
	returnNil bool
}

// mockCollection mirrors a generated *ResponseModelCollection paging with ContinuationToken.
type mockCollection struct {
	items []mockItem
	token string
}

func (m *mockCollection) GetItems() []mockItem         { return m.items }
func (m *mockCollection) SetItems(items []mockItem)    { m.items = items }
func (m *mockCollection) GetContinuationToken() string { return m.token }

// mockNextTokenCollection mirrors a Global App Configuration response paging with NextToken.
type mockNextTokenCollection struct {
	items []mockItem
	token string
}

func (m *mockNextTokenCollection) GetItems() []mockItem      { return m.items }
func (m *mockNextTokenCollection) SetItems(items []mockItem) { m.items = items }
func (m *mockNextTokenCollection) GetNextToken() string      { return m.token }

func executePage(pages map[string]mockPage, token string, calls *int) (*mockCollection, *http.Response, error) {
	*calls++
	page, ok := pages[token]
	if !ok {
		return nil, nil, fmt.Errorf("unexpected token %q", token)
	}
	if page.fail {
		return nil, nil, fmt.Errorf("boom on token %q", token)
	}
	if page.returnNil {
		// A degenerate 200 with a nil body; the helper must not panic on it.
		return nil, &http.Response{StatusCode: 200}, nil
	}
	return &mockCollection{items: page.items, token: page.next}, &http.Response{StatusCode: 200}, nil
}

// mockContinuationRequest is a paginable request builder using the ContinuationToken convention.
type mockContinuationRequest struct {
	token string
	pages map[string]mockPage
	calls *int
}

func (r mockContinuationRequest) ContinuationToken(token string) mockContinuationRequest {
	r.token = token
	return r
}

func (r mockContinuationRequest) Execute() (*mockCollection, *http.Response, error) {
	return executePage(r.pages, r.token, r.calls)
}

// mockNextTokenRequest is a paginable request builder using the NextToken convention.
type mockNextTokenRequest struct {
	token string
	pages map[string]mockPage
	calls *int
}

func (r mockNextTokenRequest) NextToken(token string) mockNextTokenRequest {
	r.token = token
	return r
}

func (r mockNextTokenRequest) Execute() (*mockNextTokenCollection, *http.Response, error) {
	*r.calls++
	page := r.pages[r.token]
	return &mockNextTokenCollection{items: page.items, token: page.next}, &http.Response{StatusCode: 200}, nil
}

// mockNonPaginableRequest exposes no token setter, so the endpoint cannot be paged.
type mockNonPaginableRequest struct {
	pages map[string]mockPage
	calls *int
}

func (r mockNonPaginableRequest) Execute() (*mockCollection, *http.Response, error) {
	return executePage(r.pages, "", r.calls)
}

func itemIds(items []mockItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.id)
	}
	return ids
}

func TestGetAllPagesWithRetry_AggregatesAllPages(t *testing.T) {
	calls := 0
	request := mockContinuationRequest{
		calls: &calls,
		pages: map[string]mockPage{
			"":   {items: []mockItem{{"a"}, {"b"}}, next: "t1"},
			"t1": {items: []mockItem{{"c"}}, next: "t2"},
			"t2": {items: []mockItem{{"d"}}, next: ""},
		},
	}

	result, httpResp, err := citrixclient.GetAllPagesWithRetry[*mockCollection](request, &citrixclient.CitrixDaasClient{})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, httpResp.StatusCode)
	assert.Equal(t, 3, calls, "should fetch one page per continuation token")
	assert.Equal(t, []string{"a", "b", "c", "d"}, itemIds(result.GetItems()))
}

func TestGetAllPagesWithRetry_SinglePage(t *testing.T) {
	calls := 0
	request := mockContinuationRequest{
		calls: &calls,
		pages: map[string]mockPage{
			"": {items: []mockItem{{"a"}}, next: ""},
		},
	}

	result, _, err := citrixclient.GetAllPagesWithRetry[*mockCollection](request, &citrixclient.CitrixDaasClient{})

	assert.NoError(t, err)
	assert.Equal(t, 1, calls, "a single page must not trigger extra requests")
	assert.Equal(t, []string{"a"}, itemIds(result.GetItems()))
}

func TestGetAllPagesWithRetry_ReturnsErrorFromLaterPage(t *testing.T) {
	calls := 0
	request := mockContinuationRequest{
		calls: &calls,
		pages: map[string]mockPage{
			"":   {items: []mockItem{{"a"}}, next: "t1"},
			"t1": {fail: true},
		},
	}

	_, _, err := citrixclient.GetAllPagesWithRetry[*mockCollection](request, &citrixclient.CitrixDaasClient{})

	assert.Error(t, err)
	assert.Equal(t, 2, calls, "paging should stop at the failing page")
}

func TestGetAllPagesWithRetry_HandlesNilResponseWithoutPanic(t *testing.T) {
	calls := 0
	request := mockContinuationRequest{
		calls: &calls,
		pages: map[string]mockPage{
			"":   {items: []mockItem{{"a"}}, next: "t1"},
			"t1": {returnNil: true},
		},
	}

	result, _, err := citrixclient.GetAllPagesWithRetry[*mockCollection](request, &citrixclient.CitrixDaasClient{})

	assert.NoError(t, err)
	assert.Equal(t, 2, calls, "a nil later-page response must stop paging without panicking")
	assert.Equal(t, []string{"a"}, itemIds(result.GetItems()), "pages gathered before the nil response must be retained")
}

func TestGetAllPagesWithRetry_StopsWhenTokenDoesNotAdvance(t *testing.T) {
	calls := 0
	request := mockContinuationRequest{
		calls: &calls,
		pages: map[string]mockPage{
			// The "t1" page keeps returning "t1" (as a filtered query can), which would loop forever
			// without the non-advancing-token guard.
			"":   {items: []mockItem{{"a"}}, next: "t1"},
			"t1": {items: []mockItem{{"b"}}, next: "t1"},
		},
	}

	result, _, err := citrixclient.GetAllPagesWithRetry[*mockCollection](request, &citrixclient.CitrixDaasClient{})

	assert.NoError(t, err)
	assert.Equal(t, 2, calls, "must stop once the continuation token stops advancing")
	assert.Equal(t, []string{"a", "b"}, itemIds(result.GetItems()))
}

func TestGetAllPagesWithRetry_FollowsNextToken(t *testing.T) {
	calls := 0
	request := mockNextTokenRequest{
		calls: &calls,
		pages: map[string]mockPage{
			"":   {items: []mockItem{{"a"}}, next: "n1"},
			"n1": {items: []mockItem{{"b"}}, next: ""},
		},
	}

	result, _, err := citrixclient.GetAllPagesWithRetry[*mockNextTokenCollection](request, &citrixclient.CitrixDaasClient{})

	assert.NoError(t, err)
	assert.Equal(t, 2, calls)
	assert.Equal(t, []string{"a", "b"}, itemIds(result.GetItems()))
}

func TestGetAllPagesWithRetry_StopsWhenRequestCannotPage(t *testing.T) {
	calls := 0
	request := mockNonPaginableRequest{
		calls: &calls,
		pages: map[string]mockPage{
			// The response advertises a next token, but the request has no setter to follow it.
			"": {items: []mockItem{{"a"}}, next: "t1"},
		},
	}

	result, _, err := citrixclient.GetAllPagesWithRetry[*mockCollection](request, &citrixclient.CitrixDaasClient{})

	assert.NoError(t, err)
	assert.Equal(t, 1, calls, "a request without a token setter must not loop")
	assert.Equal(t, []string{"a"}, itemIds(result.GetItems()))
}
