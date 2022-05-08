package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type AddrRepo struct {
	mock.Mock
}

func (r *AddrRepo) SuggestStreet(address *models.SuggestStreetRepoInput) (*models.SuggestStreetRepoAnsw, error) {
	return &models.SuggestStreetRepoAnsw{StreetSuggests: []string{"Измайловская улица", "Измайловский проспект"}}, nil
}

func (r *AddrRepo) SuggestHouse(addres *models.SuggestHouseRepoInput) (*models.SuggestHouseRepoAnsw, error) {
	return &models.SuggestHouseRepoAnsw{HouseSuggests: []string{"1"}}, nil
}

func (r *AddrRepo) GetCity(city string) (*models.GetCityRepoAnsw, error) {
	return &models.GetCityRepoAnsw{CityId: 1, Name: "Москва"}, nil
}

func (r *AddrRepo) GetStreet(street *models.GetStreetRepoInput) (*models.GetStreetRepoAnsw, error) {
	return &models.GetStreetRepoAnsw{StreetId: 1, Name: "Измайловский проспект"}, nil
}

func (r *AddrRepo) GetHouse(house *models.GetHouseRepoInput) (*models.GetHouseRepoAnsw, error) {
	return &models.GetHouseRepoAnsw{House: "1"}, nil
}

type AddrRepoErr struct {
	mock.Mock
}

func (r *AddrRepoErr) SuggestStreet(address *models.SuggestStreetRepoInput) (*models.SuggestStreetRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepoErr) SuggestHouse(addres *models.SuggestHouseRepoInput) (*models.SuggestHouseRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepoErr) GetCity(city string) (*models.GetCityRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepoErr) GetStreet(street *models.GetStreetRepoInput) (*models.GetStreetRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepoErr) GetHouse(house *models.GetHouseRepoInput) (*models.GetHouseRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

type AddrRepoHouseErr struct {
	mock.Mock
}

func (r *AddrRepoHouseErr) SuggestHouse(addres *models.SuggestHouseRepoInput) (*models.SuggestHouseRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepoHouseErr) GetHouse(house *models.GetHouseRepoInput) (*models.GetHouseRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepoHouseErr) SuggestStreet(address *models.SuggestStreetRepoInput) (*models.SuggestStreetRepoAnsw, error) {
	return &models.SuggestStreetRepoAnsw{StreetSuggests: []string{"Измайловская улица", "Измайловский проспект"}}, nil
}

func (r *AddrRepoHouseErr) GetCity(city string) (*models.GetCityRepoAnsw, error) {
	return &models.GetCityRepoAnsw{CityId: 1, Name: "Москва"}, nil
}

func (r *AddrRepoHouseErr) GetStreet(street *models.GetStreetRepoInput) (*models.GetStreetRepoAnsw, error) {
	return &models.GetStreetRepoAnsw{StreetId: 1, Name: "Измайловский проспект"}, nil
}

type AddrRepoGetHouseErr struct {
	mock.Mock
}

func (r *AddrRepoGetHouseErr) SuggestHouse(addres *models.SuggestHouseRepoInput) (*models.SuggestHouseRepoAnsw, error) {
	return &models.SuggestHouseRepoAnsw{HouseSuggests: []string{"1"}}, nil
}

func (r *AddrRepoGetHouseErr) GetHouse(house *models.GetHouseRepoInput) (*models.GetHouseRepoAnsw, error) {
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepoGetHouseErr) SuggestStreet(address *models.SuggestStreetRepoInput) (*models.SuggestStreetRepoAnsw, error) {
	return &models.SuggestStreetRepoAnsw{StreetSuggests: []string{"Измайловская улица", "Измайловский проспект"}}, nil
}

func (r *AddrRepoGetHouseErr) GetCity(city string) (*models.GetCityRepoAnsw, error) {
	return &models.GetCityRepoAnsw{CityId: 1, Name: "Москва"}, nil
}

func (r *AddrRepoGetHouseErr) GetStreet(street *models.GetStreetRepoInput) (*models.GetStreetRepoAnsw, error) {
	return &models.GetStreetRepoAnsw{StreetId: 1, Name: "Измайловский проспект"}, nil
}
