package interfaceUseCase

type IJWTUseCase interface{
	GetRestStatForAccessToken(*int) (*string,error) 
	GetUserStatForGeneratingAccessToken(*string) (*string,error)
}