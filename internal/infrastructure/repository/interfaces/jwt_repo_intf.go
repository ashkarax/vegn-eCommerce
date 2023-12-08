package interfaceRepository

type IJWTRepo interface{
	GetRestStatForGeneratingAccessToken(*int)(*string,error)
	GetUserStatForGeneratingAccessToken(*string) (*string,error)

}