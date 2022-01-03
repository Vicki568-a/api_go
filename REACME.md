#мелисса от вики 

#REST API

GET /chat ---list of chat--- 200, 404, 500
GET /chat/:id ---chat by id--- 200, 404, 500
POST /chat/:id ---create chat --- 204,4xx,Header Location :url
PUT /chat/:id ---fully update user --- 200/204,404,400,500
PATCH /chat/id ---partially update user --- 200/204,404,400,500
DELETE /chat/id ---Delete user id --- 204,404,400
GET /users ---list of user--- 200, 404, 500
GET /users/:id ---user by id--- 200, 404, 500
POST /users/:id ---create user --- 204,4xx,Header Location :url
PUT /users/:id ---fully update user --- 200/204,404,400,500
PATCH /users/id ---partially update user --- 200/204,404,400,500
DELETE /users/id ---Delete user id --- 204,404,400


