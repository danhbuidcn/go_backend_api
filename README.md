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
    + run `go run cmd/server/main.go` 
    + call `curl http://0.0.0.0:8002/v1/ping/1515\?uid='12313'`
    + call `curl -H "Authorization: valid-token" http://0.0.0.0:8002/v1/user/1`

- `09`: Go Backend: GIN vs TEST MOCK
    + run `go get github.com/stretchr/testify` or `go mod tidy`
    + run test
    ```
    cd tests/basic
    go test --v --coverprofile=coverage.out
    // OR
    go tool cover --html=coverage.out -o coverage.html
    open coverage.html
    ```

- `10`: Go Backend: Cấu trúc file main.go cho DỰ ÁN LỚN
    + run `go run cmd/server/main.go` 

- `11`: Go Backend: QUẢN LÝ LOGs CHO DỰ ÁN LỚN
    + run `go get gopkg.in/natefinch/lumberjack.v2`
    + `make run`

- `12`: Go Backend: Làm việc với Mysql Pool, Tại sao lại là Pool
    + run `go get -u gorm.io/gorm` [gorm package](https://gorm.io/docs/)
    + run `go get github.com/google/uuid`
    + setup docker mysql 
    ```
    docker-compose up -d
    docker exec -it mysql_shopdevgo bash
        # mysql -u root -proot
            > use shopdevgo
            > show tables;
            > desc go_db_user;
    ```
    + run [benchmark](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
    ```
    cd tests/benchmark
    go test -bench=. -benchmem -count 5
    ```

- `13`: Go Backend: Làm việc với Redis sentinel, cluster, Tại sao lại là Sentinel

- `14`: Go Backend: Router cho TEAM LỚN

- `15`: Go Kafka Backend: Kafka thực hành về mua bán cổ phiếu với các tình huống

- `16`: Go Interview: Không sử dụng Interface có được không? ĐƯỢC vs MẤT khi không sử dụng?

- `17`: Go Backend: Interface cách triển khai nếu bạn là member

- `18`: Go Backend: Nói về Dependency Injection và sủ dụng Wire trong dự án

- `19`: Go Backend: Triển khai BlackList IP, WhiteList IP cho hệ thống api

- `20`: Go Backend: User-Register: Send OTP, VerifyOTP, BLock IP Spam...

## Resource

- [Con đường Lập Trình Viên](https://www.youtube.com/playlist?list=PLw0w5s5b9NK5fDx409WXgT06Zm4P83yiA)
- [Go - Con đường lập trình](https://www.youtube.com/playlist?list=PL0Pnqmz-onB5Xk2o46BpvGHa-c8Toe0YP)
- [Database](https://www.youtube.com/playlist?list=PLw0w5s5b9NK6EOfWmhaoGK9Jjih3CzR6c)
- [Series Phỏng Vấn](https://www.youtube.com/playlist?list=PLw0w5s5b9NK5xE6lmH85ge8dFXHheYV2o)
- [Message Queue](https://www.youtube.com/playlist?list=PLw0w5s5b9NK4yji-f3c3L7htTJjiyhRDa)
- [See more](https://www.youtube.com/@anonystick/playlists)