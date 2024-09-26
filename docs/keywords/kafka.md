## Apache Kafka

- **Apache Kafka**: Nền tảng xử lý luồng dữ liệu phân tán theo thời gian thực.
- **Mô hình**: Publish-Subscribe.
- **Phát triển bởi**: LinkedIn, hiện thuộc Apache Software Foundation.
- **Đặc điểm chính**:
  - **Lưu trữ lâu dài**: Dữ liệu có thể được lưu trữ trên đĩa.
  - **Phân phối dữ liệu**: Dữ liệu được phân phối và xử lý theo luồng.
  - **Tốc độ cao**: Xử lý lượng dữ liệu lớn với độ trễ thấp.
  - **Mở rộng tốt**: Có khả năng mở rộng theo chiều ngang với nhiều broker.
  - **Chịu lỗi**: Dữ liệu được sao chép để đảm bảo không mất dữ liệu khi có sự cố.
- **Ứng dụng**: Hệ thống xử lý sự kiện, luồng dữ liệu thời gian thực, logging phân tán.

## Zookeeper:

- **Zookeeper**: Hệ thống quản lý và điều phối phân tán cho các dịch vụ lớn như Apache Kafka.
- **Theo dõi và quản lý dịch vụ**: Giám sát trạng thái các broker (máy chủ Kafka) và các dịch vụ phân tán khác.
- **Chọn leader**: Đảm bảo chỉ có một leader trong hệ thống để điều phối các hoạt động.
- **Đồng bộ hóa dịch vụ**: Giúp các dịch vụ phân tán hoạt động cùng nhịp, tránh xung đột.
- **Quản lý metadata**: Lưu trữ và quản lý thông tin cấu hình và trạng thái của các dịch vụ như Kafka.
- **Phân phối khóa**: Cung cấp khóa phân tán để tránh truy cập đồng thời vào tài nguyên.
