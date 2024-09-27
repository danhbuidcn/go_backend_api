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

- `13`: Go Backend: Làm việc với [Redis sentinel, cluster](https://github.com/redis/go-redis), Tại sao lại là Sentinel
    + install redis `go get github.com/redis/go-redis/v9`
    + install debugging `go install github.com/go-delve/delve/cmd/dlv@latest`

- `14`: Go Backend: Router cho TEAM LỚN

- `15`: Go Kafka Backend: Kafka thực hành về mua bán cổ phiếu với các tình huống (xem [Kafka section](#kafka-section))
    + installation [kafka](https://github.com/segmentio/kafka-go): `go get -u github.com/segmentio/kafka-go`
    + run
    ```
    docker-compose up -d
    go run cmd/cli/kafka/kafka.go
    curl -X POST "http://127.0.0.1:8999/action/stock?msg=HPG&type=MUA"

    // create topic
    docker exec -it broker kafka-topics --create --topic user_topic_001 --bootstrap-server broker:9091 --partitions 3 --replication-factor 1
    docker exec -it broker kafka-topics --bootstrap-server broker:9091 --list

    // list consumer group
    docker exec -it broker kafka-consumer-groups --bootstrap-server broker:9091 --list

    // ensure JMX is opening
    docker exec -it kafka-ui /bin/sh
    nc -vz broker 19101
    ```
    + access `http://localhost:8080/`

- `16`: Go Interview: Không sử dụng Interface có được không? ĐƯỢC vs MẤT khi không sử dụng?
    + [Interface In Go](https://gobyexample.com/interfaces)
    + [Interfaces in Go (part I)](https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c)

- `17`: Go Backend: Interface cách triển khai nếu bạn là member

- `18`: Go Backend: Nói về Dependency Injection và sủ dụng Wire trong dự án

- `19`: Go Backend: Triển khai BlackList IP, WhiteList IP cho hệ thống api

- `20`: Go Backend: User-Register: Send OTP, VerifyOTP, BLock IP Spam...

- `21`: Go Backend: Chiến đấu với Docker Link, Docker Compose build Project Level 2, 3, 4

- `BONUS`: Go Backend: Đồng bộ dữ liệu Mysql to Kafka sử dụng Debezium với N Tables tốc độ REALTIME

- `22`: Go Backend: Tại sao nhiều công ty lại sử dụng SQLC

- `23`: Go Backend: Từ khi sử dụng GOOSE tốc độ làm việc Database nhanh gấp đôi

- `24`: Go Backend: Triển khai quy trình send OTP cho User Registration

- `25`: Go Backend: Đến lúc chia tay GORM vì sao?

- `26`: Go Backend: TEAM JAVA hỗ trợ send OTP cho TEAM GO

- `27`: Go Backend: Đến lúc làm việc với Kafka

- `28`: Go Backend: Quyết định refactor interface sau khi nhìn lén code đồng nghiệp

- `29`: Go Backend: Đến lúc trở thành Senior và làm việc với makefile, diagram mysql

## Resource

- [Con đường Lập Trình Viên](https://www.youtube.com/playlist?list=PLw0w5s5b9NK5fDx409WXgT06Zm4P83yiA)
- [Go - Con đường lập trình](https://www.youtube.com/playlist?list=PL0Pnqmz-onB5Xk2o46BpvGHa-c8Toe0YP)
- [Database](https://www.youtube.com/playlist?list=PLw0w5s5b9NK6EOfWmhaoGK9Jjih3CzR6c)
- [Series Phỏng Vấn](https://www.youtube.com/playlist?list=PLw0w5s5b9NK5xE6lmH85ge8dFXHheYV2o)
- [Message Queue](https://www.youtube.com/playlist?list=PLw0w5s5b9NK4yji-f3c3L7htTJjiyhRDa)
- [See more](https://www.youtube.com/@anonystick/playlists)

## Kafka section

- `01: kafka dùng khi nào ?`[Kafka đã thay đổi hệ thống eCommerce trở nên mạnh mẽ như thế nào so với cách cũ](https://www.youtube.com/watch?v=yK4T7Myi9N4)
- `02: kafka có 7 mấu chốt`[Kafka: Đây là 7 thứ đủ để bắt đầu cuộc chiến TOPIC, PARTITIONs và Consumer Group](https://www.youtube.com/watch?v=a7lmP5hdgB0)
- `03: kafka vs stocks`[Kafka: Ứng dụng thực tế hệ thống MUA BÁN Backend API](https://www.youtube.com/watch?v=UIFWVisug1M&t=3s)
- `04: sync data from mysql to kafka` [KỸ SƯ CAO CẤP: Cách đồng bộ dữ liệu Mysql to Kafka sử dụng Debezium với N Tables tốc độ REALTIME](https://www.youtube.com/watch?v=KqLzls2xCnQ&t=291s)

## gRPC

- [Cách sử dụng gRPC: Triển khai Order Service Go, Java! Độ trễ gần như KHÔNG CÓ](https://www.youtube.com/watch?v=x5ZVWiJIuSM)

## Packages usefull

- [godepgraph](https://github.com/alovn/godepgraph?tab=readme-ov-file): godepgraph is a packages dependency graph visualization tool for your local go module project.
