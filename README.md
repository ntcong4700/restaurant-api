# Bakend-Restaurant

### Project Structure

````plaintext
restaurant-api/
├── config/
│   └── database.go        # Cấu hình và kết nối MySQL (db.go cũ)
├── models/
│   ├── menu.go            # Entity món ăn (Struct)
│   ├── order.go           # Entity hóa đơn/đơn hàng (Struct)
│   └── order_detail.go    # Entity chi tiết món ăn (Struct)
├── repositories/
│   ├── menu_repo.go       # Các câu lệnh SQL truy vấn Menu
│   ├── order_repo.go      # Các câu lệnh SQL liên quan đến đặt món, tính tiền
│   └── report_repo.go     # Các câu lệnh SQL tính toán Doanh thu (Umsatz)
├── handlers/
│   ├── menu_handler.go    # Tiếp nhận request API thực đơn và trả về JSON
│   ├── order_handler.go   # Tiếp nhận request API gọi món, thanh toán
│   └── report_handler.go  # Tiếp nhận request API xem doanh thu Umsatz
├── routes/
│   └── routes.go          # Quy định các đường dẫn API (URL Routing)
├── .env                   # Lưu thông tin bảo mật (Mật khẩu DB, Port)
├── go.mod                 # Quản lý các dependency gói thư viện Go
├── go.sum                 # Bản checksum của các gói thư viện
└── main.go                # Entrypoint - Điểm khởi chạy duy nhất của ứng dụng
````