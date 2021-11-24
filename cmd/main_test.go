package main

// func TestPingRoute(t *testing.T) {
// 	config := &config.Config
// 	engine := setupServer(config)
// 	// The setupServer method, that we previously refactored
// 	// is injected into a test server
// 	server := httptest.NewServer(engine)
// 	// Shut down the server and block until all requests have gone through
// 	defer server.Close()

// 	// Make a request to our server with the {base url}/ping
// 	resp, err := http.Get(fmt.Sprintf("%s/api/ping", server.URL))

// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
// 	}

// 	val, ok := resp.Header["Content-Type"]

// 	// Assert that the "content-type" header is actually set
// 	if !ok {
// 		t.Fatalf("Expected Content-Type header to be set")
// 	}

// 	// Assert that it was set as expected
// 	if val[0] != "application/json; charset=utf-8" {
// 		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
// 	}
// }
