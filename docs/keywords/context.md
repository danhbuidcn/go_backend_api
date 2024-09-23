
- `Context` trong Go 
    + Giúp quản lý thời gian chờ (timeout), hủy tác vụ, và truyền dữ liệu giữa các hàm. 
    + Thường được dùng khi xử lý các tác vụ như gọi API, truy vấn cơ sở dữ liệu hoặc chạy nhiều goroutines để kiểm soát tốt hơn.
    ```go
    package main

    import (
        "context"
        "fmt"
        "time"
    )

    func main() {
        // Tạo context với thời gian chờ 2 giây
        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
        defer cancel()

        // Goroutine giả lập một tác vụ tốn thời gian
        go func() {
            select {
            case <-time.After(3 * time.Second):
                fmt.Println("Completed task")
            case <-ctx.Done():
                fmt.Println("Task canceled:", ctx.Err())
            }
        }()

        time.Sleep(4 * time.Second)
    }
    ```

- `context.Background()` trong Go :
    + Là một context gốc không có giá trị, không có thời gian hết hạn và không thể hủy. 
    + Thường được sử dụng làm điểm bắt đầu cho các context khác hoặc trong những trường hợp không cần quản lý thời gian hoặc hủy bỏ tác vụ. 
    + Context giúp quản lý vòng đời của các tác vụ, truyền dữ liệu giữa các hàm, và xử lý hủy tác vụ hoặc đặt thời gian hết hạn.
