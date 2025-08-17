package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	colors := []string{
		"#FF6B6B", "#4ECDC4", "#45B7D1", "#96CEB4", "#FFEAA7",
		"#DDA0DD", "#98D8C8", "#F7DC6F", "#BB8FCE", "#85C1E9",
		"#F8C471", "#82E0AA", "#F1948A", "#85C1E9", "#D7BDE2",
		"#A3E4D7", "#F9E79F", "#D5A6BD", "#AED6F1", "#A9DFBF",
		"#F5B7B1", "#D2B4DE", "#AED6F1", "#A9CCE3", "#FADBD8",
	}
	
	fmt.Fprintf(w, `<html>
<head>
	<style>
		body { font-family: 'Courier New', monospace; background: #1a1a1a; padding: 20px; }
		h1 { color: #00ff00; text-align: center; }
		.hello-line { font-size: 18px; margin: 5px 0; font-weight: bold; }
	</style>
</head>
<body>
	<h1>🌈 Colorful Hello World App 🌈</h1>
	<div>`)
	
	for i := 1; i <= 25; i++ {
		color := colors[i-1]
		fmt.Fprintf(w, `<div class="hello-line" style="color: %s;">Hello World #%d</div>`, color, i)
	}
	
	fmt.Fprintf(w, `	</div>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}