# docker build . -t name-demo
# docker images
# docker run -p 8005:8080 name-demo
# # Sử dụng Go 1.23 từ official image
# FROM golang:1.23-alpine

# WORKDIR /build

# # Sao chép mã nguồn vào image
# COPY . .

# # Build ứng dụng
# RUN go build -o nameDemo .

# # Tạo thư mục dist và di chuyển binary vào đó
# WORKDIR /dist
# RUN cp /build/nameDemo /dist/

# # Mở cổng 8009
# EXPOSE 8009

# # Chạy ứng dụng
# CMD ["/dist/nameDemo"]
#-------------------------
# docker build . -t name-demo-1
# docker run -p 8004:8080 name-demo-1
# Sử dụng golang:1.23-alpine làm môi trường build ứng dụng
FROM golang:1.23-alpine AS builder

# Đặt thư mục làm việc để build ứng dụng
WORKDIR /build

# Sao chép mã nguồn từ máy chủ vào image
COPY . .

# Build ứng dụng thành file thực thi 'nameDemo1'
RUN go build -o nameDemo1 .

# Sử dụng scratch, môi trường nhẹ nhất, chỉ chứa những gì cần thiết để chạy ứng dụng
FROM scratch

# Sao chép file thực thi từ môi trường build sang môi trường scratch
COPY --from=builder /build/nameDemo1 /

# Thiết lập file thực thi chạy khi container khởi động
ENTRYPOINT ["/nameDemo1"]
