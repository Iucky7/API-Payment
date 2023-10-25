package constant

const (
	CREATE_MERCHANT = "INSERT INTO merchant(id,name)VALUES($1,$2)RETURNING id,name"
	LIST_MERCHANT = "SELECT * FROM merchant"


	CREATE_PAYMENT = "INSERT INTO payment(id,merchant_id,bank_account,amount) VALUES ($1,$2,$3,$4) RETURNING id,merchant_id,bank_account,amount"
	LIST_PAYMENT = "SELECT * FROM payment"


	CREATE_USER_CREDENTIAL = "INSERT INTO user_credential(id,username,password)VALUES($1,$2,$3) RETURNING id,username,password"
	LIST_USER_CREDENTIAL = "SELECT * FROM user_credential"
	GET_USER_CREDENTIAL_BY_USERNAME = "Select * FROM user_credential where username = $1"
)