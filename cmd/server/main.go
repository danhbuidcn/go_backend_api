package main

import "github.com/danhbuidcn/go_backend_api/internal/initialize"

func main() {
	// r := routers.NewRouter()

	// r.Run(":8002") // listen and serve on 0.0.0.0:8002 (for windows "localhost:8002")
	initialize.Run()
}
