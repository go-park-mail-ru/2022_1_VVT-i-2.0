package address

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Repository interface {
	SuggestStreet(address *models.SuggestStreetRepoInput) (*models.SuggestStreetRepoAnsw, error)
	SuggestHouse(addres *models.SuggestHouseRepoInput) (*models.SuggestHouseRepoAnsw, error)
	GetCity(city string) (*models.GetCityRepoAnsw, error)
	GetStreet(street *models.GetStreetRepoInput) (*models.GetStreetRepoAnsw, error)
	GetHouse(house *models.GetHouseRepoInput) (*models.GetHouseRepoAnsw, error)
	// SuggestHouse(address *models.SuggestInputRepo) (*models.SuggestAnswerRepo, error)
}
