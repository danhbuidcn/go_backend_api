# Project go_backend_api

One Paragraph of project description goes here

## Getting Started

[course](https://github.com/danhbuidcn/go_backend_api)

## Series Backend eCommerce GOLANG (Redis, mysql, elasticSearch, kafak, nginx, mongodb...)

- `01`: Go backend: Các kiến trúc phổ phiến khởi tạo dự án BackEnd
    + [go-blueprint gihub](https://github.com/Melkeydev/go-blueprint)
    + [blueprint.dev](https://go-blueprint.dev/)

- `02`: Go backend: Setup dự án BackEnd
    + [directory_structure](./docs/directory_structure.md)
    + run command
    ```
    go mod init github.com/danhbuidcn/go_backend_api
    go mod tidy
    ```

- `03`: Go backend: GIN vs ROUTER
    + install [gin web framework](https://github.com/gin-gonic/gin) `go get -u github.com/gin-gonic/gin`
    + run `go run cmd/server/main.go` and call `curl http://0.0.0.0:8002/v1/ping/1515\?uid='12313'`

- `04`: Go Backend: GIN vs MVC
    + run `go run cmd/server/main.go` and call `curl http://0.0.0.0:8002/v1/user/1`

- `05`: Go Backend: GIN vs ERROR HANDLER

- `06`: Go backend: GIN vs LOGGER HANDLER
    + install [zap](https://github.com/uber-go/zap) `go get -u go.uber.org/zap`
    + run `go run cmd/cli/main.log.go`

- `07`: Go backend: GIN vs VIPER
    + install [viper](https://github.com/spf13/viper) `go get github.com/spf13/viper` : a complete configuration solution
    + run `go run cmd/cli/viper/main.viper.go`

- `08`: Go Backend: GIN vs MIDDLEWARES

- `09`: Go Backend: GIN vs TEST MOCK

- `10`: Go Backend: Cấu trúc file main.go cho DỰ ÁN LỚN

- `11`: Go Backend: QUẢN LÝ LOGs CHO DỰ ÁN LỚN

- `12`: Go Backend: Làm việc với Mysql Pool, Tại sao lại là Pool

- `13`: Go Backend: Làm việc với Redis sentinel, cluster, Tại sao lại là Sentinel

- `14`: Go Backend: Router cho TEAM LỚN

- `15`: Go Kafka Backend: Kafka thực hành về mua bán cổ phiếu với các tình huống

- `16`: Go Interview: Không sử dụng Interface có được không? ĐƯỢC vs MẤT khi không sử dụng?

- `17`: Go Backend: Interface cách triển khai nếu bạn là member

- `18`: Go Backend: Nói về Dependency Injection và sủ dụng Wire trong dự án

- `19`: Go Backend: Triển khai BlackList IP, WhiteList IP cho hệ thống api

- `20`: Go Backend: User-Register: Send OTP, VerifyOTP, BLock IP Spam...
