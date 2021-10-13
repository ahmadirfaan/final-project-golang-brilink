-- Pemilihan database untuk dimasukkan ke dalam role ---
use brilink;

INSERT INTO roles (created_at, updated_at, role) values (current_time, current_time, "Admin");
INSERT INTO roles (created_at, updated_at, role) values (current_time, current_time, "Customer");
