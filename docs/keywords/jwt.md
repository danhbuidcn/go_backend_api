Token, đặc biệt là **JWT (JSON Web Token)**, thường được sinh ra để thay thế cho session trong một số tình huống nhất định, nhất là khi ứng dụng cần tính linh hoạt và hiệu suất cao hơn. Dưới đây là 3 tình huống điển hình khi **JWT** thay thế **session**:

### 1. **Ứng dụng phân tán (Distributed Systems)**
   - **Tình huống**: Khi ứng dụng được triển khai trên nhiều server hoặc microservices khác nhau, việc duy trì trạng thái session trên mỗi server có thể phức tạp và khó quản lý.
   - **Tại sao sử dụng JWT**: Với JWT, không cần lưu trữ trạng thái phiên (session) trên server. Thông tin người dùng và quyền truy cập được mã hóa trong token, cho phép các server khác nhau đều có thể xác thực người dùng mà không cần truy cập vào kho lưu trữ phiên tập trung.
   - **Cách thay thế**: Token JWT chứa thông tin đăng nhập hoặc quyền truy cập, được client (ví dụ: trình duyệt hoặc ứng dụng) gửi kèm trong mỗi yêu cầu HTTP, giúp xác thực và ủy quyền nhanh chóng trên bất kỳ server nào mà không cần đồng bộ hóa session.

### 2. **Ứng dụng API không trạng thái (Stateless APIs)**
   - **Tình huống**: Khi phát triển các **RESTful API** hoặc **GraphQL API**, chúng thường được thiết kế theo mô hình không trạng thái (stateless), nghĩa là mỗi yêu cầu đến server phải tự chứa đầy đủ thông tin để xác thực.
   - **Tại sao sử dụng JWT**: JWT cho phép API xử lý các yêu cầu mà không cần lưu trạng thái của phiên trên server. Mỗi yêu cầu đều đi kèm JWT chứa thông tin xác thực và quyền hạn, giúp API dễ dàng xác thực người dùng mà không cần lưu trạng thái phiên.
   - **Cách thay thế**: Thay vì lưu session trong bộ nhớ server, mỗi lần đăng nhập, một token JWT được tạo và gửi đến client. Sau đó, client gửi JWT này trong header của mỗi yêu cầu API, giúp xác thực nhanh chóng mà không cần server duy trì trạng thái.

### 3. **Ứng dụng di động hoặc SPA (Single Page Application)**
   - **Tình huống**: Với các ứng dụng di động hoặc SPA như React, Vue.js, Angular, thường yêu cầu việc xác thực mượt mà, không phụ thuộc vào phiên của server, để có thể hoạt động linh hoạt với các yêu cầu từ nhiều thiết bị khác nhau.
   - **Tại sao sử dụng JWT**: JWT có thể được lưu trữ an toàn trên client (trong localStorage hoặc sessionStorage) và gửi đi cùng mỗi yêu cầu. Điều này phù hợp với các ứng dụng di động hoặc SPA, nơi mà client cần lưu trạng thái đăng nhập trong một khoảng thời gian dài và truy cập tài nguyên qua API.
   - **Cách thay thế**: Sau khi người dùng đăng nhập, JWT được tạo ra và lưu trữ trên client. Sau đó, tất cả các yêu cầu từ client đến server sẽ kèm theo JWT trong **header** để server xác thực và phản hồi.

---

### **Tóm lại, JWT phù hợp khi:**
- Cần quản lý phiên một cách **phi tập trung** (không cần lưu trạng thái trên server).
- Ứng dụng có tính chất **phân tán** hoặc bao gồm nhiều **microservices**.
- Muốn xây dựng các API **không trạng thái**.
- Cần xác thực trên các nền tảng **di động** hoặc **SPA** mà không muốn lưu session trên server.
