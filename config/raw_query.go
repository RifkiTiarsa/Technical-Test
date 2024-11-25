package config

const (
	InsertCustomer        = `INSERT INTO customers (name, email, password, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	GetCustomerById       = `SELECT id, name, email, password, balance, created_at, updated_at FROM customers WHERE id=$1`
	GetCustomerByUsername = `SELECT id, name, email, password, balance, created_at, updated_at FROM customers WHERE email=$1`

	InsertMerchant = `INSERT INTO merchants (name, balance, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	InsertProduct  = `INSERT INTO products (merchant_id, name, nominal, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	GetProductById = `SELECT id, merchant_id, name, nominal, price, created_at, updated_at FROM products WHERE id=$1`
	GetAllProducts = `SELECT id, merchant_id, name, nominal, price, created_at, updated_at FROM products`
	UpdateProduct  = `UPDATE products SET name=$1, nominal=$2, price=$3, updated_at=$4 WHERE id=$5`
	DeleteProduct  = `DELETE FROM products WHERE id=$1`

	InsertTopup           = `INSERT INTO topups (customer_id, merchant_id, product_id, payment_method, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	GetTopupById          = `SELECT id, customer_id, merchant_id, product_id, payment_method, status, created_at, updated_at FROM topups WHERE id=$1`
	UpdateStatusTopup     = `UPDATE topups SET status=$1 where id = $2`
	UpdateBalanceCustomer = `UPDATE customers SET balance= balance + $1 WHERE id=$2`
	UpdateBalanceMerchant = `UPDATE merchants SET balance= balance - $1 WHERE id=$2`
)
