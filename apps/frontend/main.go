package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	version    = os.Getenv("APP_VERSION")
	backendURL = os.Getenv("BACKEND_URL")
)

func main() {
	if version == "" {
		version = "v1"
	}
	if backendURL == "" {
		backendURL = "https://api.iqbalhakim.ink"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head><title>iqbalhakim.ink</title>
<style>
  body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; background: #f0f4f8; }
  .card { background: white; padding: 2rem 3rem; border-radius: 12px; box-shadow: 0 4px 20px rgba(0,0,0,0.1); text-align: center; }
  h1 { color: #2d3748; } p { color: #718096; }
  #backend { margin-top: 1rem; padding: 1rem; background: #ebf8ff; border-radius: 8px; font-family: monospace; font-size: 0.9rem; }
</style>
</head>
<body>
<div class="card">
  <h1>iqbalhakim.ink</h1>
  <p>Frontend <strong>%s</strong> &mdash; pod: <code>%s</code></p>
  <div id="backend">loading backend...</div>
</div>
<script>
  fetch('%s')
    .then(r => r.json())
    .then(d => document.getElementById('backend').innerText = 'backend: ' + JSON.stringify(d))
    .catch(e => document.getElementById('backend').innerText = 'backend error: ' + e);
</script>
</body>
</html>`, version, hostname, backendURL)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("frontend %s listening on :%s\n", version, port)
	http.ListenAndServe(":"+port, nil)
}
