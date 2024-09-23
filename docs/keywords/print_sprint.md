### Giải thích về các hàm **Print** trong Go:

Các hàm **Print** in trực tiếp ra màn hình

1. **`fmt.Print()`**: In ra các tham số liền nhau, không có dấu cách hay xuống dòng.
   ```go
   fmt.Print("Hello", "World")
   // Output: HelloWorld
   ```

2. **`fmt.Println()`**: In các tham số với dấu cách giữa chúng và tự động xuống dòng sau khi in.
   ```go
   fmt.Println("Hello", "World")
   // Output: Hello World\n
   ```

3. **`fmt.Printf()`**: In với định dạng tùy chỉnh, không tự động thêm dấu cách hoặc xuống dòng.
   ```go
   fmt.Printf("Hello %s", "World")
   // Output: Hello World
   ```

### Các hàm **Sprint** trong Go:

Các hàm **Sprint** trả về chuỗi đã định dạng

1. **`fmt.Sprint()`**: Tương tự như `fmt.Print()`, nhưng trả về chuỗi đã nối.
   ```go
   s := fmt.Sprint("Hello", "World")
   fmt.Println(s)
   // Output: HelloWorld
   ```

2. **`fmt.Sprintln()`**: Tương tự như `fmt.Println()`, nhưng trả về chuỗi có dấu cách giữa các tham số và tự động thêm dấu xuống dòng.
   ```go
   s := fmt.Sprintln("Hello", "World")
   fmt.Print(s)
   // Output: Hello World\n
   ```

3. **`fmt.Sprintf()`**: Tương tự như `fmt.Printf()`, nhưng trả về chuỗi đã định dạng thay vì in ra màn hình.
   ```go
   s := fmt.Sprintf("Hello %s", "World")
   fmt.Println(s)
   // Output: Hello World
   ```
