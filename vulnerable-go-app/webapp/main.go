package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Middleware for logging
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s %s from %s\n", start.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func main() {
	mockDB := map[string]string{
		"1": "admin",
		"2": "pawan",
		"secret_data": "FLAG{K8S_TOP10_MASTER_KEY}",
	}

	// 1. Index Page
	http.HandleFunc("/", logger(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<h1>Hello by <a href='https://peachycloudsecurity.com'>peachycloudsecurity.com</a></h1>")
		fmt.Fprint(w, "<p>Welcome to the Vulnerable K8s Lab. Test endpoints: <code>/db</code>, <code>/config</code>, <code>/exec</code></p>")
	}))

	// 2. SQLi Endpoint
	http.HandleFunc("/db", logger(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Printf("[LOG] SQL Injection attempt with ID: %s\n", id)

		if strings.Contains(strings.ToUpper(id), "OR") || strings.Contains(strings.ToUpper(id), "UNION") {
			fmt.Fprintf(w, "User Found: %s (SQLi Successful!)", mockDB["secret_data"])
			return
		}

		if val, ok := mockDB[id]; ok {
			fmt.Fprintf(w, "User Found: %s", val)
		} else {
			fmt.Fprintf(w, "User Not Found")
		}
	}))

	// 3. Path Traversal (K01)
	http.HandleFunc("/config", logger(func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Query().Get("source")
		fmt.Printf("[LOG] File Access attempt: %s\n", file)
		data, _ := os.ReadFile(file) 
		w.Write(data)
	}))

	// 4. Command Injection (K08)
	http.HandleFunc("/exec", logger(func(w http.ResponseWriter, r *http.Request) {
		cmdArg := r.URL.Query().Get("run")
		fmt.Printf("[LOG] Command Exec attempt: %s\n", cmdArg)
		out, _ := exec.Command("sh", "-c", cmdArg).CombinedOutput()
		fmt.Fprintf(w, "%s", out)
	}))

	fmt.Println("------------------------------------------")
	fmt.Println("🚀 Server: http://localhost:8080")
	fmt.Println("🔗 peachycloudsecurity.com")
	fmt.Println("📡 Logging: ACTIVE")
	fmt.Println("------------------------------------------")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
