# Project go_backend_api

Golang Programming Course with Real Project CRM  eCommerce

## Getting Started

[Course - Go Backend Architecture](https://www.youtube.com/playlist?list=PLw0w5s5b9NK6qiL9Xzki-mGbq_V8dBQkY)


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
    + `Inversion of Control (Ioc)`: 
        [LINK 1](https://stackoverflow.com/questions/3058/what-is-inversion-of-control),
        [LINK 2](https://martinfowler.com/bliki/InversionOfControl.html)
    + `Dependency Injection (DI)`
    + [wire](https://github.com/google/wire) :
        ```
        go install github.com/google/wire/cmd/wire@latest
        go get github.com/google/wire/cmd/wire
        cd internal/wire/
        wire
        ```

- `19`: Migrating Schema with GORM to MYSQL and DUMP Database
    + mysql command to go model
    ```sql
    -- db up
    DROP TABLE IF EXISTS `go_crm_user`;
    -- new
    CREATE TABLE `go_crm_user` (
    `usr_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
    `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
    `usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone Number',
    `usr_username` varchar(30) NOT NULL DEFAULT '' COMMENT 'Username',
    `usr_password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
    `usr_created_at` int NOT NULL DEFAULT '0' COMMENT 'Creation Time',
    `usr_updated_at` int NOT NULL DEFAULT '0' COMMENT 'Update Time',
    `usr_create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Creation IP',
    `usr_last_login_at` int NOT NULL DEFAULT '0' COMMENT 'Last Login Time',
    `usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last Login IP',
    `usr_login_times` int NOT NULL DEFAULT '0' COMMENT 'Login Times',
    `usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`usr_id`),
    KEY `idx_email` (`usr_email`),
    KEY `idx_phone` (`usr_phone`),
    KEY `idx_username` (`usr_username`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account';

    DESC go_crm_user;
    ```
    + [gen in gorm](https://gorm.io/gen/index.html):
    ```bash
    docker-compose up
    go get -u gorm.io/gen
    make run

    // command to export database to .sql file
    docker exec -it mysql_shopdevgo mysqldump -uroot -proot --databases shopdevgo --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > migrations/shopdevgo.sql
    ```

- `20`: Chiến đấu DOCKER, sai lầm cách build này Level 0, 1
    + comand : [gin](https://github.com/gin-gonic/gin)
    ```bash
    cd demo_docker
    go mod init example.com/demo_docker
    touch main.go
    touch Dockerfile
    go get -u github.com/gin-gonic/gin
    touch test.http
    go run .

    GET http://localhost:8080/ping
    ```
    + build Dockerfile from [hub.docker](https://hub.docker.com/_/golang/tags?name=alpine)

- `21`: Chiến đấu với Docker Link, Docker Compose build Project Level 2
    + build go with docker and link to mysql container
    ```bash
    docker-compose up -d # run mysql container

    docker build . -t backend_api # build go project
    docker network connect bridge mysql_shopdevgo # connect bridge to mysql_shopdevgo network
    docker run --link mysql_shopdevgo:mysql_shopdevgo -p 8002:8002 backend_api

    docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql
    docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis

    docker start backend_shopdevgo # star backend_api if it's unhealthy
    ```
- `22`: Công ty đề nghị chuyển [GORM](https://gorm.io/index.html) sang [SQLC](https://sqlc.dev/) như thế nào?
    + command
    ```bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    cd sqlc_practice
    sqlc generate
    go run cmd/main.go
    ```

- `23`: Goose (A database migration tool) hiệu suất cao của dân BACKEND

- `24`: [User Register Send OTP To Email Template](https://backend.anonystick.com/golang/go-24.html)

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
