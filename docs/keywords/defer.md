## defer

- Việc sử dụng `defer sqlDB.Close()` ở đầu hàm thay vì viết trực tiếp ở cuối hàm có lý do quan trọng liên quan đến **quản lý tài nguyên** và **đảm bảo rằng các kết nối sẽ được đóng bất kể khi nào và bằng cách nào hàm thoát ra**.

`tests/benchmark/mysql_pool_test.go`

### Lý do chính:

1. **Đảm bảo tài nguyên luôn được giải phóng**:
   - **`defer`** đảm bảo rằng khi hàm `BenchmarkMaxOpenConns` kết thúc, cho dù nó kết thúc bình thường hay do gặp lỗi ở bất kỳ đâu trong hàm, thì **hàm `sqlDB.Close()` sẽ luôn được gọi** để đóng kết nối cơ sở dữ liệu.
   - Nếu bạn đặt `sqlDB.Close()` trực tiếp ở cuối hàm mà không sử dụng `defer`, nó sẽ chỉ được gọi nếu hàm kết thúc một cách tuần tự mà không gặp bất kỳ lỗi nào. Nếu có lỗi xảy ra trước đó, chẳng hạn như trong quá trình mở kết nối hoặc chạy benchmark, kết nối có thể không bao giờ được đóng, gây ra **rò rỉ kết nối** (connection leak).

2. **Dễ quản lý và an toàn**:
   - Đặt `defer` ngay sau khi mở kết nối giúp cho việc **quản lý tài nguyên dễ dàng** và **an toàn hơn**, vì bạn không cần phải nhớ đóng kết nối ở tất cả các trường hợp có thể xảy ra.
   - Nếu bạn đặt `sqlDB.Close()` ở cuối hàm sau `b.RunParallel`, bạn cần phải chắc chắn rằng mọi trường hợp thoát ra khỏi hàm đều đảm bảo kết nối được đóng, điều này có thể phức tạp nếu có nhiều luồng hoặc điều kiện khác nhau.

3. **Phù hợp với phong cách Go**:
   - Việc gọi `defer` ngay sau khi mở tài nguyên (như kết nối cơ sở dữ liệu) là một **phong cách lập trình phổ biến** trong Go. Nó giúp mã ngắn gọn, an toàn và dễ đọc, đồng thời tránh lỗi liên quan đến việc quên đóng tài nguyên.

### Ví dụ so sánh:

#### Sử dụng `defer`:

```go
sqlDB, err := db.DB()
if err != nil {
    log.Fatalln("failed to get sql.DB from gorm.DB: ", err)
}
defer sqlDB.Close() // Đảm bảo luôn đóng kết nối khi hàm kết thúc

b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
        insertRecord(b, db)
    }
})
```

#### Đặt `Close()` ở cuối hàm mà không dùng `defer`:

```go
sqlDB, err := db.DB()
if err != nil {
    log.Fatalln("failed to get sql.DB from gorm.DB: ", err)
}

b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
        insertRecord(b, db)
    }
})

sqlDB.Close() // Chỉ đóng kết nối nếu hàm kết thúc bình thường
```

### Điểm yếu của cách không dùng `defer`:

- Nếu một lỗi xảy ra trước khi hàm chạy đến dòng `sqlDB.Close()`, kết nối sẽ không được đóng và dẫn đến việc rò rỉ kết nối.
- Phức tạp hơn nếu trong hàm có nhiều điểm kết thúc tiềm năng, và bạn phải chắc chắn rằng tài nguyên luôn được giải phóng ở mọi điểm thoát.
