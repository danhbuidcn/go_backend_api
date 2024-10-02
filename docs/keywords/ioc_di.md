# Inversion of Control (Ioc) & Dependency Injection (DI)

## Inversion of Control

- Là 1 nguyên tắc thiết kế (trong OOP) để giảm sự kết nối giữa các block code với nhau

## Dependency Injection

- Là một phương pháp phổ biến trong IoC
- Các module cấp cao không phụ thuộc vào các module cấp thấp
- Ưu điểm: Giảm khảo năng phụ thuộc giữa các module. tăng khả năng linh hoạt của mã code
- Nhược điểm: tính liên kết cao, giảm khả năng mở rộng và gỡ lỗi

```go
// use dependency injection
func NewUserRepoSitory(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// non use dependency injection
func NewUserRepoSitory1() *UserRepository  {
	db,_ = gorm.Open(...) // connect to database
	return &UserRepository{db: db}
}
```