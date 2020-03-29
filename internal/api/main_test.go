package api_test

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/Nashmeg/Usage.git/internal/api"
)

var resolver api.Resolver

func randomString(n int) *string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	s := string(b)
	return &s
}

func randomInt64() *int64 {
	var min, max int64
	min = 9000000
	max = 90000000
	rand.Seed(time.Now().Unix())
	i := rand.Int63n(max-min) + min
	return &i
}

var ctx context.Context

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	startup(m)
	code := m.Run()
	os.Exit(code)
}

// startup sets up the database and starts the test server.
func startup(m *testing.M) {
	os.Setenv("PORT", "8082")
	//testServer := httptest.NewServer(server.Router())
	//usage.Config.UsageURL := testServer.URL + "/api/usage"
}
