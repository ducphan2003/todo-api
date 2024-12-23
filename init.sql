CREATE TABLE users (
    id SERIAL PRIMARY KEY, -- Khóa chính tự tăng
    name VARCHAR(255) NOT NULL UNIQUE, -- Tên, không cho phép null, duy nhất
    password TEXT NOT NULL, -- Mật khẩu, không cho phép null
    salt TEXT NOT NULL, -- Salt, không cho phép null
    status VARCHAR(50) DEFAULT 'active', -- Trạng thái, mặc định là 'active'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Tự động ghi thời gian tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Tự động ghi thời gian cập nhật
    deleted_at TIMESTAMP DEFAULT NULL -- Thời gian xóa, cho phép null
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY, -- Khóa chính tự tăng
    user_id INT NOT NULL, -- Khóa ngoại tham chiếu đến users.id
    title VARCHAR(255) NOT NULL, -- Tiêu đề, không cho phép null
    description TEXT, -- Mô tả, cho phép null
    progress VARCHAR(100), -- Tiến độ, cho phép null
    priority VARCHAR(50), -- Mức độ ưu tiên
    status VARCHAR(50) DEFAULT 'active', -- Trạng thái, mặc định là 'active'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời gian tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời gian cập nhật
    deleted_at TIMESTAMP DEFAULT NULL, -- Thời gian xóa, cho phép null

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) 
        ON DELETE CASCADE ON UPDATE CASCADE -- Khóa ngoại với bảng users
);
