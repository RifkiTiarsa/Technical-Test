INSERT INTO merchants (name, balance) 
VALUES 
('TopUp Merchant', 1000000),
('E-Store', 1000000),
('Gift Cards', 1000000);

INSERT INTO products (merchant_id, name, nominal, price) 
VALUES 
(4, 'OVO 10rb', 10000, 11000),
(4, 'OVO 20rb', 20000, 21000),
(5, 'OVO 50rb', 50000, 51000),
(6, 'OVO 100rb', 100000, 101000);

INSERT INTO topups (customer_id, merchant_id, product_id, payment_method, status) 
VALUES 
(1, 1, 1, 'BCA', 'pending'),
(2, 1, 2, 'BNI', 'pending'),
(3, 3, 4, 'BRI', 'pending');

