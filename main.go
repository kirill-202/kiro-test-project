package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	fmt.Fprintf(w, `<html>
<head>
	<style>
		body { 
			font-family: Arial, sans-serif; 
			background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
			margin: 0;
			padding: 20px;
			min-height: 100vh;
		}
		h1 { 
			color: white; 
			text-align: center; 
			text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
			margin-bottom: 30px;
		}
		.hello-text {
			color: white;
			font-size: 18px;
			font-weight: bold;
			text-shadow: 1px 1px 2px rgba(0,0,0,0.7);
			line-height: 1.8;
			text-align: center;
		}
	</style>
</head>
<body>
	<h1>Hello World App</h1>
	<div class="hello-text">`)
	
	for i := 1; i <= 25; i++ {
		fmt.Fprintf(w, "Hello World<br>")
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