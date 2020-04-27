package rest_v1

//region create-user

//swagger:parameters createUser
type SwaggerCreateNewUserRequest struct {
	//in:body
	Body CreateNewUserRequest
}

type CreateNewUserRequest struct {
	//unique:true
	//required:true
	Name string
	//required:true
	Password string
}

//swagger:response createUser
type SwaggerCreateNewUserResponse struct {
	//in:body
	Body CreateNewUserResponse
}

type CreateNewUserResponse struct {
	ID   uint
	Name string
}

//endregion
//region get-user

//swagger:parameters getUser
type SwaggerGetUserRequest struct {
	//in: body
	//required:true
	Body GetUserRequest
}

type GetUserRequest struct {
	//required:true
	Name string
	//required:true
	Password string
}

//swagger:response getUser
type SwaggerGetUserResponse struct {
	//in:body
	Body GetUserResponse
}

type GetUserResponse struct {
	ID       uint
	Name     string
	Balance  int
	CCNumber string
}

//endregion
//region get-user-by-token

//swagger:parameters getUserByToken
type SwaggerGetUserByTokenRequest struct {
	//User token
	//in:header
	//required:true
	Authorization string
}

//swagger:response getUserByToken
type SwaggerGetUserByTokenResponse struct {
	//in:body
	Body GetUserResponse
}

//endregion
