Trong Go, `*var` và `&var` có ý nghĩa khác nhau liên quan đến con trỏ (pointer). Dưới đây là sự khác biệt giữa hai khái niệm này:

1. **`&var` (địa chỉ của biến):**
   - `&` là toán tử lấy địa chỉ của một biến.
   - Khi bạn sử dụng `&var`, bạn sẽ nhận được địa chỉ của biến `var`, tức là con trỏ trỏ đến biến đó.
   - **Ví dụ:**
     ```go
     var x int = 10
     var p *int = &x // p là con trỏ trỏ đến địa chỉ của x
     fmt.Println(p)  // In ra địa chỉ của x
     ```

2. **`*var` (giá trị tại địa chỉ con trỏ):**
   - `*` là toán tử dereference, tức là truy xuất giá trị mà con trỏ đang trỏ tới.
   - Khi bạn sử dụng `*var`, bạn lấy giá trị được lưu tại địa chỉ mà con trỏ `var` đang trỏ tới.
   - **Ví dụ:**
     ```go
     var x int = 10
     var p *int = &x    // p là con trỏ trỏ đến x
     fmt.Println(*p)    // Truy xuất giá trị của x thông qua con trỏ p
     ```
