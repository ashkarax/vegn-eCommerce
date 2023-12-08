package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/ashkarax/vegn-eCommerce/pkg/aws"
	"github.com/go-playground/validator/v10"
)

type DishUseCase struct {
	DishRepo interfaceRepository.IDishRepo
}

func NewDishUseCase(dishRepo interfaceRepository.IDishRepo) interfaceUseCase.IDishUseCase {
	return &DishUseCase{DishRepo: dishRepo}
}

func (r *DishUseCase) NewDish(dishData *requestmodels.DishReq) (*responsemodels.DishRes, error) {
	var resDishData responsemodels.DishRes
	BucketFolder := "vegn-ecommerce-api/vegn-dishes/"
	var imageURLs []string

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(dishData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "RestaurantID":
					resDishData.RestaurantID = 0000000000000000
				case "Name":
					resDishData.Name = "should be a valid Name. "
				case "CategoryId":
					resDishData.CategoryId = 0000000000000000
				case "Description":
					resDishData.Description = "should be a valid Description. "
				case "CuisineType":
					resDishData.CuisineType = "should be a valid CuisineType "
				case "Price":
					resDishData.Price = 000000.00000000
				case "PortionSize":
					resDishData.PortionSize = "should have two or more digit"
				case "DietaryInformation":
					resDishData.DietaryInformation = " should only have a maximum of 10 words "
				case "resDishData.Calories":
					resDishData.Calories = 00000000000
				case "resDishData.Protein":
					resDishData.Protein = 00000000000
				case "resDishData.Carbohydrates":
					resDishData.Carbohydrates = 00000000000
				case "resDishData.Fat":
					resDishData.Fat = 00000000000
				case "SpiceLevel":
					resDishData.SpiceLevel = "should only have a maximum of 10 words"
				case "AllergenInformation":
					resDishData.AllergenInformation = "should only have a maximum of 30 words"
				case "RecommendedPairings":
					resDishData.RecommendedPairings = "should only have a maximum of 30 words"
				case "SpecialFeatures":
					resDishData.SpecialFeatures = "should only have a maximum of 10 words"
				case "ImageURL1":
					resDishData.ImageURL1 = "required"
				case "ImageURL2":
					resDishData.ImageURL2 = "required"
				case "ImageURL3":
					resDishData.ImageURL3 = "required"
				case "PreparationTime":
					resDishData.PreparationTime = "should only have a maximum of 15 words"
				case "PromotionDiscount":
					resDishData.PromotionDiscount = "should only have a maximum of 15 words"
				case "StoryOrigin":
					resDishData.StoryOrigin = "should only have a maximum of 100 words"

				}
			}
			fmt.Println(err)
			return &resDishData, err
		}
	}

	numFiles := len(dishData.Image)
	if numFiles != 3 {
		return &resDishData, errors.New("you have to upload exactly 3 images,nothing more....nothing less")
	}

	for _, image := range dishData.Image {
		if image.Size > 5*1024*1024 { // 5 MB limit
			return &resDishData, errors.New("image size exceeds the limit (5MB)")
		}
	}

	// Validate file types
	allowedTypes := map[string]struct{}{
		"image/jpeg": {},
		"image/png":  {},
		"image/gif":  {},
	}

	for _, image := range dishData.Image {
		// Open the file to read the header
		file, err := image.Open()
		if err != nil {
			return &resDishData, err
		}
		defer file.Close()

		// Read the first 512 bytes to determine the content type
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			return &resDishData, err
		}

		// Reset the file position after reading
		_, err = file.Seek(0, 0)
		if err != nil {
			return &resDishData, err
		}

		// Get the content type based on the file content
		contentType := http.DetectContentType(buffer)

		// Check if the content type is allowed
		if _, ok := allowedTypes[contentType]; !ok {
			return &resDishData, errors.New("unsupported file type,should be a jpeg,png or gif")
		}
	}

	sess, errInit := aws.AWSSessionInitializer()
	if errInit != nil {
		fmt.Println(errInit)
		return &resDishData, errInit
	}

	for i, image := range dishData.Image {
		imageURL, err := aws.AWSImageUploader(image, sess, &BucketFolder)
		if err != nil {
			fmt.Printf("Error uploading image %d: %v\n", i+1, err)
			return &resDishData, err
		}
		imageURLs = append(imageURLs, *imageURL)
	}

	dishData.ImageURL1 = imageURLs[0]
	dishData.ImageURL2 = imageURLs[1]
	dishData.ImageURL3 = imageURLs[2]

	insertErr := r.DishRepo.AddNewDish(dishData)
	if insertErr != nil {
		fmt.Println(insertErr)
		return &resDishData, insertErr
	}
	return &resDishData, nil

}

func (r *DishUseCase) FetchAllDishesForRestaurant(restaurantId *int) (*[]responsemodels.DishRes, error) {
	resDishMap, err := r.DishRepo.FetchAllDishesForARestaurant(restaurantId)
	if err != nil {
		return resDishMap, err
	}
	return resDishMap, nil
}
func (r *DishUseCase) DishById(dishId *int) (*responsemodels.DishRes, error) {
	resDish, err := r.DishRepo.FetchDishById(dishId)
	if err != nil {
		return resDish, err
	}
	return resDish, nil
}
func (r *DishUseCase) UpdateDishDetails(dishData *requestmodels.DishUpdateReq, id *int) (*responsemodels.DishRes, error) {
	var resDishData responsemodels.DishRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(dishData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "RestaurantID":
					resDishData.RestaurantID = 0000000000000000
				case "Name":
					resDishData.Name = "should be a valid Name. "
				case "CategoryId":
					resDishData.CategoryId = 0000000000000000
				case "Description":
					resDishData.Description = "should be a valid Description. "
				case "CuisineType":
					resDishData.CuisineType = "should be a valid CuisineType "
				case "Price":
					resDishData.Price = 000000.00000000
				case "PortionSize":
					resDishData.PortionSize = "should have two or more digit"
				case "DietaryInformation":
					resDishData.DietaryInformation = " should only have a maximum of 10 words "
				case "resDishData.Calories":
					resDishData.Calories = 00000000000
				case "resDishData.Protein":
					resDishData.Protein = 00000000000
				case "resDishData.Carbohydrates":
					resDishData.Carbohydrates = 00000000000
				case "resDishData.Fat":
					resDishData.Fat = 00000000000
				case "SpiceLevel":
					resDishData.SpiceLevel = "should only have a maximum of 10 words"
				case "AllergenInformation":
					resDishData.AllergenInformation = "should only have a maximum of 30 words"
				case "RecommendedPairings":
					resDishData.RecommendedPairings = "should only have a maximum of 30 words"
				case "SpecialFeatures":
					resDishData.SpecialFeatures = "should only have a maximum of 10 words"
				case "PreparationTime":
					resDishData.PreparationTime = "should only have a maximum of 15 words"
				case "PromotionDiscount":
					resDishData.PromotionDiscount = "should only have a maximum of 15 words"
				case "StoryOrigin":
					resDishData.StoryOrigin = "should only have a maximum of 100 words"
				case "Availability":
					resDishData.Availability = false
				case "RemainingQuantity":
					resDishData.RemainingQuantity = 00000000000000000000000000

				}
			}
			fmt.Println(err)
			return &resDishData, err
		}
	}
	insertErr := r.DishRepo.UpdateDish(dishData, id)
	if insertErr != nil {
		fmt.Println(insertErr)
		return &resDishData, insertErr
	}
	return &resDishData, nil

}

func (r *DishUseCase) DeleteDish(dishId *string) error {
	err := r.DishRepo.DeleteDishById(dishId)
	if err != nil {
		return err
	}
	return nil
}

func (r *DishUseCase) GetAllDishesForUser() (*[]responsemodels.DishRes, error) {
	resDishMap, err := r.DishRepo.GetAllDishesForUser()
	if err != nil {
		return resDishMap, err
	}
	return resDishMap, nil
}

func (r *DishUseCase) FetchDishesByCategoryId(id *string) (*[]responsemodels.DishRes, error) {
	var resDishMap *[]responsemodels.DishRes
	
	idInt, _ := strconv.Atoi(*id)
	if idInt == 0 {
		return resDishMap, errors.New("category with id=0 does not exist")
	}

	resDishMap, err := r.DishRepo.FetchDishesByCategoryId(id)
	if err != nil {
		return resDishMap, err
	}
	return resDishMap, nil
}
