## Interface là gì?

- **Interface** trong Go là một loại **abstract type** (kiểu trừu tượng) mà chỉ định hành vi (behavior) của một đối tượng thông qua một tập hợp các phương thức (methods).
- Interface không định nghĩa cụ thể các phương thức, mà chỉ yêu cầu đối tượng nào implement interface phải cung cấp định nghĩa cho các phương thức đó.
- **Go's interfaces are implicit**, nghĩa là một kiểu (type) tự động implement một interface nếu nó định nghĩa tất cả các phương thức mà interface yêu cầu, không cần từ khóa `implements` như các ngôn ngữ khác.

##  Bối cảnh

Giả sử bạn đang xây dựng một ứng dụng với hai loại thanh toán: **CreditCard** và **PayPal**. 
Mục tiêu là xử lý các khoản thanh toán mà không phải quan tâm đến chi tiết cụ thể của từng loại thanh toán.

### **Sử dụng Interface**:

Chúng ta sẽ định nghĩa một **interface** để mô tả hành vi chung cho các loại thanh toán khác nhau.

```go
// Interface Payment
type Payment interface {
    Pay(amount float64) string
}

// CreditCard implement interface Payment
type CreditCard struct {}

func (c CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid %v using Credit Card", amount)
}

// PayPal implement interface Payment
type PayPal struct {}

func (p PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid %v using PayPal", amount)
}

// Function to process payment
func ProcessPayment(p Payment, amount float64) {
    fmt.Println(p.Pay(amount))
}

func main() {
    // Chúng ta có thể truyền bất kỳ kiểu thanh toán nào (CreditCard, PayPal) vào ProcessPayment
    ProcessPayment(CreditCard{}, 100.0)
    ProcessPayment(PayPal{}, 150.0)
}
```

### **Không sử dụng Interface**:

Bây giờ, ta sẽ viết lại chương trình trên mà **không sử dụng interface**.

```go
// CreditCard without using interface
type CreditCard struct {}

func (c CreditCard) PayWithCreditCard(amount float64) string {
    return fmt.Sprintf("Paid %v using Credit Card", amount)
}

// PayPal without using interface
type PayPal struct {}

func (p PayPal) PayWithPayPal(amount float64) string {
    return fmt.Sprintf("Paid %v using PayPal", amount)
}

func main() {
    // Phải gọi phương thức cụ thể của từng loại thanh toán
    creditCard := CreditCard{}
    fmt.Println(creditCard.PayWithCreditCard(100.0))

    payPal := PayPal{}
    fmt.Println(payPal.PayWithPayPal(150.0))
}
```

### **Được khi không sử dụng Interface** (cụ thể):

1. **Code đơn giản hơn**: Bạn không cần khai báo interface và không cần lo lắng về việc tất cả các loại thanh toán phải implement interface đó. Điều này giúp code trở nên gọn gàng hơn nếu bạn chỉ cần xử lý từng loại thanh toán riêng lẻ.
    - **Ví dụ**: Trong ví dụ trên, thay vì sử dụng interface, bạn chỉ cần gọi trực tiếp các phương thức như `PayWithCreditCard` hoặc `PayWithPayPal`. Code này khá đơn giản và không yêu cầu abstraction.

2. **Hiệu suất tốt hơn**: Interface thêm một chút chi phí cho việc sử dụng abstraction, mặc dù nó rất nhỏ. Tuy nhiên, khi không sử dụng interface, bạn có thể xử lý logic nhanh hơn vì không cần phải sử dụng dynamic dispatch để gọi các phương thức.
    - **Ví dụ**: Không có chi phí từ việc triển khai interface, code của bạn có thể chạy nhanh hơn một chút, đặc biệt là khi bạn chỉ có một hoặc hai loại thanh toán cố định.

### **Mất khi không sử dụng Interface** (cụ thể):

1. **Mất khả năng mở rộng**: Nếu muốn thêm nhiều loại thanh toán khác trong tương lai (ví dụ: **Bitcoin**, **Stripe**, v.v.), bạn phải thay đổi logic trong hàm `main` để thêm các phương thức cụ thể cho từng loại. Điều này sẽ làm code bị phình to và phức tạp khi thêm nhiều loại thanh toán.
    - **Ví dụ**: Bạn sẽ phải viết thêm nhiều phương thức như `PayWithBitcoin`, `PayWithStripe`, v.v. Điều này khiến code trở nên cồng kềnh và khó bảo trì khi số lượng loại thanh toán tăng lên.

2. **Giảm tính tái sử dụng**: Nếu bạn muốn viết một hàm dùng chung cho các loại thanh toán khác nhau mà không cần biết đó là **CreditCard**, **PayPal** hay bất kỳ loại nào khác, việc không sử dụng interface sẽ buộc bạn phải viết các hàm riêng cho từng loại.
    - **Ví dụ**: Bạn không thể sử dụng một hàm `ProcessPayment` để xử lý chung tất cả các loại thanh toán. Mỗi loại thanh toán sẽ cần một hàm xử lý riêng, điều này làm giảm tính tái sử dụng code.

3. **Khó kiểm thử hơn**: Interface giúp dễ dàng thay thế các thành phần thật bằng các **mock object** khi kiểm thử. Nếu không sử dụng interface, bạn sẽ phải tạo ra nhiều thành phần cụ thể cho từng loại thanh toán và không thể mock chung cho tất cả.
    - **Ví dụ**: Nếu sử dụng interface, bạn có thể mock `Payment` trong quá trình kiểm thử. Nhưng nếu không có interface, bạn phải tạo các mock cụ thể cho từng loại thanh toán như `MockCreditCard`, `MockPayPal`, điều này làm kiểm thử trở nên phức tạp.

4. **Tăng coupling (ràng buộc)**: Không có interface khiến các thành phần của hệ thống bị ràng buộc chặt với nhau. Điều này làm giảm tính linh hoạt khi bạn muốn thay thế hoặc mở rộng chức năng.
    - **Ví dụ**: `main` sẽ bị ràng buộc với từng loại thanh toán cụ thể như `CreditCard`, `PayPal`, và nếu bạn muốn thay đổi cách thanh toán hoạt động, bạn phải sửa đổi hàm `main`, điều này vi phạm nguyên tắc **mở rộng - đóng gói** trong lập trình.

### Kết luận

- **ĐƯỢC**: Khi bạn làm việc với một ứng dụng đơn giản, ít thay đổi và chỉ cần xử lý một vài hành vi cố định, việc không sử dụng interface có thể giúp code gọn nhẹ và dễ hiểu hơn.
- **MẤT**: Khi bạn cần khả năng mở rộng, tái sử dụng code, dễ kiểm thử và giảm sự ràng buộc giữa các thành phần, interface sẽ là lựa chọn tốt hơn, và việc không sử dụng nó sẽ khiến code của bạn trở nên cồng kềnh và khó quản lý trong dài hạn.
