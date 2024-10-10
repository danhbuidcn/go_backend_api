### **Cookie**:

#### **Đặc điểm**:
- **Lưu trữ trên máy người dùng**: Cookie được lưu trên trình duyệt của người dùng, dùng để ghi nhớ thông tin giữa các phiên truy cập khác nhau.
- **Dung lượng hạn chế**: Cookie có dung lượng tối đa khoảng 4KB.
- **Thời gian tồn tại tùy chỉnh**: Cookie có thể có thời hạn cố định hoặc hết hạn khi trình duyệt đóng.
- **Gửi trong mỗi yêu cầu HTTP**: Cookie được tự động gửi kèm trong tất cả các yêu cầu từ trình duyệt đến server.

#### **Mục đích sử dụng**:
- **Dùng khi nào**: Khi cần lưu trữ thông tin đơn giản, ít dung lượng mà cần ghi nhớ lâu dài trên trình duyệt, ngay cả khi người dùng đóng trình duyệt.
- **Mục đích**: Dùng để lưu thông tin người dùng như thông tin đăng nhập (nhớ mật khẩu), các lựa chọn cá nhân hóa, hoặc theo dõi hành vi người dùng (tracking).

#### **Lưu ý khi dùng**:
- **Bảo mật**: Cần mã hóa dữ liệu nhạy cảm và sử dụng cookie với thuộc tính `HttpOnly` và `Secure` để tránh đánh cắp qua tấn công XSS.
- **Tuân thủ pháp luật**: Các trang web cần tuân thủ quy định bảo vệ quyền riêng tư, như **GDPR**, yêu cầu thông báo cho người dùng về việc sử dụng cookie.

#### **Hạn chế**:
- **Dung lượng nhỏ**: Giới hạn khoảng 4KB nên không thể lưu trữ dữ liệu lớn.
- **Dễ bị đánh cắp**: Nếu không bảo mật tốt, cookie có thể bị đánh cắp qua các cuộc tấn công như **XSS**.
- **Gây tăng tải dữ liệu**: Cookie được gửi trong mọi yêu cầu HTTP, có thể làm tăng kích thước dữ liệu truyền tải.
- **Dễ bị xóa/chặn**: Người dùng có thể xóa hoặc chặn cookie.

#### **Giới hạn**:
- **Dung lượng**: Khoảng 4KB cho mỗi cookie.
- **Số lượng**: Giới hạn số lượng cookie mỗi domain (~20-50 cookie/domain).
- **Không chia sẻ được giữa các domain**: Cookie chỉ áp dụng trong phạm vi một domain.

---

### **Session**:

#### **Đặc điểm**:
- **Lưu trữ trên server**: Session lưu trữ thông tin trên server, thường dùng để lưu trạng thái người dùng trong suốt phiên làm việc.
- **Thời gian tồn tại ngắn**: Session hết hạn sau một khoảng thời gian không tương tác (session timeout).
- **Liên kết với session ID**: Session được xác định bởi session ID, thường lưu trong cookie hoặc URL.

#### **Mục đích sử dụng**:
- **Dùng khi nào**: Khi cần lưu trữ thông tin trạng thái phiên làm việc của người dùng, như đăng nhập hoặc giỏ hàng, nhưng không cần lưu dài hạn.
- **Mục đích**: Dùng để duy trì phiên làm việc an toàn cho người dùng, không phụ thuộc vào trình duyệt và dễ dàng quản lý trên server.

#### **Lưu ý khi dùng**:
- **Bảo mật session ID**: Session ID cần được bảo mật thông qua cookie `HttpOnly` và `Secure` để tránh bị tấn công XSS hoặc nghe lén.
- **Tối ưu hiệu suất**: Cần có cơ chế quản lý session tốt khi ứng dụng có nhiều người dùng để tránh tải nặng lên server.

#### **Hạn chế**:
- **Phụ thuộc vào server**: Server cần xử lý và lưu trữ tất cả các session, dễ gây tải lớn khi có nhiều người dùng.
- **Không thể chia sẻ giữa các domain**: Session chỉ tồn tại trong phạm vi một domain.
- **Thời gian sống ngắn**: Session có thời gian tồn tại ngắn và hết hạn nếu người dùng không tương tác.

#### **Giới hạn**:
- **Thời gian sống**: Hết hạn sau một khoảng thời gian không tương tác.
- **Không lưu trữ được nhiều dữ liệu**: Session chủ yếu lưu trữ dữ liệu cần thiết cho phiên làm việc.
- **Không chia sẻ được giữa server khác nhau**: Cần cơ chế đồng bộ nếu hệ thống có nhiều server.

---

### **So sánh Cookie và Session**

#### **Giống nhau**:
- **Lưu trữ dữ liệu người dùng**: Cả cookie và session đều được sử dụng để lưu trữ thông tin người dùng nhằm duy trì trạng thái giữa các yêu cầu HTTP (vốn là stateless, không giữ được trạng thái giữa các lần request).
- **Sử dụng để duy trì phiên làm việc**: Đều có thể dùng để duy trì trạng thái đăng nhập hoặc các thông tin liên quan đến phiên làm việc của người dùng.
- **Có thể được sử dụng song song**: Cookie thường được dùng để lưu trữ session ID giúp xác định session đang hoạt động trên server.

#### **Khác nhau**:

| **Tiêu chí**               | **Cookie**                                               | **Session**                                              |
|----------------------------|----------------------------------------------------------|----------------------------------------------------------|
| **Nơi lưu trữ**             | Lưu trữ trên trình duyệt của người dùng (client-side).   | Lưu trữ trên server (server-side).                       |
| **Dung lượng**              | Giới hạn khoảng 4KB cho mỗi cookie.                      | Không giới hạn kích thước, tùy thuộc vào tài nguyên server. |
| **Thời gian tồn tại**       | Có thể tồn tại lâu dài, tùy theo cấu hình expiration date. | Thời gian ngắn, thường hết hạn sau khi phiên làm việc kết thúc hoặc một khoảng thời gian không hoạt động. |
| **Bảo mật**                 | Dễ bị đánh cắp qua các cuộc tấn công XSS nếu không được bảo mật tốt. | An toàn hơn vì dữ liệu được lưu trên server, chỉ session ID được truyền qua client. |
| **Truyền dữ liệu**          | Cookie được gửi kèm theo mỗi yêu cầu HTTP.               | Chỉ session ID được gửi qua client, dữ liệu session ở trên server. |
| **Quản lý dữ liệu**         | Dữ liệu không thể quá lớn vì hạn chế dung lượng và được lưu trên máy người dùng. | Quản lý trên server, lưu trữ dữ liệu lớn dễ dàng hơn.     |
| **Tùy chỉnh bởi người dùng**| Người dùng có thể xóa hoặc chặn cookie từ trình duyệt.  | Người dùng không thể xóa session trực tiếp, nhưng session sẽ tự hết hạn. |
| **Phạm vi hoạt động**       | Có thể được chia sẻ giữa các yêu cầu trong cùng một domain hoặc subdomain. | Chỉ có hiệu lực trên server và không thể chia sẻ giữa các domain hoặc subdomain. |


#### **Tóm tắt sử dụng**:
- **Cookie**: Dùng khi cần lưu thông tin lâu dài trên máy người dùng, ví dụ như lưu trữ thông tin đăng nhập, theo dõi hành vi người dùng, hoặc lưu các tùy chỉnh cá nhân.
- **Session**: Dùng khi cần bảo mật dữ liệu người dùng tốt hơn và không cần lưu trữ lâu dài. Phù hợp cho việc lưu trạng thái đăng nhập, giỏ hàng, hoặc các thông tin phiên tạm thời.

---

### **Ví dụ về quy trình hoạt động của Facebook**:

1. **Đăng nhập**: 
   - Người dùng nhập tên đăng nhập và mật khẩu. Nếu xác thực thành công, Facebook tạo ra một session trên server và lưu session ID vào cookie trên máy của người dùng.
   - Nếu người dùng chọn "Nhớ mật khẩu", một cookie riêng biệt chứa token sẽ được lưu trữ để duy trì trạng thái đăng nhập.

2. **Duy trì trạng thái**:
   - Mỗi lần người dùng truy cập Facebook hoặc thực hiện thao tác mới, session ID trong cookie sẽ được gửi lên server để Facebook xác nhận người dùng và duy trì phiên làm việc.

3. **Kết thúc phiên**:
   - Nếu người dùng không hoạt động trong một khoảng thời gian hoặc chọn đăng xuất, session sẽ bị hủy trên server. Cookie chứa session ID sẽ mất hiệu lực, và người dùng sẽ phải đăng nhập lại để tạo session mới.

4. **Tự động đăng nhập**:
   - Khi người dùng truy cập lại vào Facebook, nếu cookie "Nhớ mật khẩu" vẫn còn hiệu lực, Facebook sẽ sử dụng token từ cookie để tạo một session mới mà không yêu cầu người dùng nhập lại thông tin đăng nhập.
