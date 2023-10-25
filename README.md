# api-payment

-- database structure can be seen in config/database/init.sql

-- env can be set according to local needs

-- In this payment api, users are required to register first by accessing the localhost:8089/api/v1/register endpoint by entering a name and password.

-- then after registering the user logs in first using the username and password that has been registered by accessing the localhost:8089/api/v1/login endpoint and the user will get a JWT token, which will be used before making the payment process or creating a new merchant. if you don't do the login process the user will not be able to make a payment because it has not been authorized.

-- After logging in and getting a token, the user can make a payment by accessing the localhost:8089/api/v1/payment endpoint by using the bearer token authorization with the token obtained at login.

-- if the user wants to logout the user can access the endpoint localhost:8089/api/v1/logout
