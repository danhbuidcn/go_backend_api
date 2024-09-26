## Redis

**Redis** (Remote Dictionary Server) là một **cơ sở dữ liệu NoSQL** dạng **in-memory** (lưu trữ dữ liệu trong bộ nhớ) có tốc độ cao, được sử dụng như một **key-value store**. Redis có thể hoạt động như một bộ nhớ đệm (cache), một cơ sở dữ liệu, hoặc một message broker.

### Đặc điểm chính của Redis:

- **In-memory**: Redis lưu trữ dữ liệu trong bộ nhớ RAM, giúp tăng tốc độ truy xuất dữ liệu rất nhanh chóng, nhưng cũng có khả năng lưu trữ trên đĩa.
- **Key-Value Store**: Dữ liệu trong Redis được lưu trữ dưới dạng cặp khóa-giá trị.
- **Dữ liệu đa dạng**: Redis hỗ trợ nhiều kiểu dữ liệu như chuỗi (string), danh sách (list), tập hợp (set), băm (hash), và nhiều hơn nữa.
- **Tốc độ cao**: Vì lưu trữ trong RAM, Redis có khả năng xử lý hàng triệu yêu cầu mỗi giây.
- **Persistence (lưu trữ lâu dài)**: Redis có thể cấu hình để lưu trữ dữ liệu trên đĩa, đảm bảo dữ liệu không bị mất khi tắt máy.
- **Replication và Cluster**: Redis hỗ trợ sao lưu dữ liệu và phân tán dữ liệu trên nhiều máy chủ, giúp tăng cường tính sẵn sàng và mở rộng.

### Ứng dụng của Redis:

- **Caching**: Redis được sử dụng phổ biến để lưu trữ các dữ liệu tạm thời nhằm giảm tải cho cơ sở dữ liệu chính và tăng tốc độ truy cập.
- **Session Store**: Dùng để lưu trữ session cho các ứng dụng web.
- **Message Broker**: Redis có thể làm hệ thống message broker đơn giản với cơ chế pub/sub.
- **Leaderboards**: Quản lý và xử lý dữ liệu cho các bảng xếp hạng hoặc dữ liệu thời gian thực.

## Redis Sentinel:

- **Chức năng**: Đảm bảo tính **khả dụng cao** (high availability) cho các instance Redis.
- **Nhiệm vụ chính**:
  - **Giám sát (monitoring)**: Theo dõi các instance Redis và phát hiện khi chúng không hoạt động.
  - **Chuyển đổi dự phòng (failover)**: Tự động chuyển đổi máy chủ master thành slave khi phát hiện lỗi.
  - **Thông báo (notification)**: Gửi thông báo khi phát hiện sự cố hoặc thay đổi trạng thái.
  - **Cấu hình lại (reconfiguration)**: Tự động cấu hình lại các instance slave để chúng trở thành master mới.
- **Sử dụng**: Phù hợp cho các hệ thống cần **tính sẵn sàng cao** và tự động quản lý chuyển đổi trong trường hợp lỗi.

## Redis Cluster:

- **Chức năng**: **Phân phối Redis** trên nhiều nút để đảm bảo **khả năng mở rộng** và **chịu lỗi**.
- **Nhiệm vụ chính**:
  - **Phân vùng dữ liệu**: Redis Cluster tự động phân chia dữ liệu thành các **slot** và phân phối chúng trên các nút khác nhau.
  - **Sao chép dữ liệu**: Dữ liệu của mỗi slot được sao chép trên các nút để đảm bảo dữ liệu không bị mất khi có sự cố.
  - **Tự động cân bằng**: Phân phối tải và dữ liệu một cách tự động giữa các nút khi có thêm hoặc mất nút.
  - **Tự động xử lý lỗi**: Nếu một nút bị lỗi, Cluster sẽ tự động phục hồi bằng cách sử dụng dữ liệu từ các bản sao.
- **Sử dụng**: Phù hợp cho các ứng dụng cần **khả năng mở rộng** lớn và **độ tin cậy** cao khi làm việc với lượng dữ liệu lớn.
