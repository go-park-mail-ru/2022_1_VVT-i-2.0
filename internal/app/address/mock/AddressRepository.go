package mock

// import (
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
// 	"github.com/stretchr/testify/mock"
// )

// type AddrRepo struct {
// 	mock.Mock
// }

// func (r *AddrRepo) SuggestStreet(address *models.SuggestStreetRepoReq) (*models.SuggestStreetRepoResp, error) {
// 	return &models.SuggestStreetRepoResp{StreetSuggests: []string{"Измайловская улица", "Измайловский проспект"}}, nil
// }

// func (r *AddrRepo) SuggestHouse(addres *models.SuggestHouseRepoReq) (*models.SuggestHouseRepoResp, error) {
// 	return &models.SuggestHouseRepoResp{HouseSuggests: []string{"1"}}, nil
// }

// func (r *AddrRepo) GetCity(city string) (*models.GetCityRepoResp, error) {
// 	return &models.GetCityRepoResp{CityId: 1, Name: "Москва"}, nil
// }

// func (r *AddrRepo) GetStreet(street *models.GetStreetRepoReq) (*models.GetStreetRepoResp, error) {
// 	return &models.GetStreetRepoResp{StreetId: 1, Name: "Измайловский проспект"}, nil
// }

// func (r *AddrRepo) GetHouse(house *models.GetHouseRepoReq) (*models.GetHouseRepoResp, error) {
// 	return &models.GetHouseRepoResp{House: "1"}, nil
// }

// type AddrRepoErr struct {
// 	mock.Mock
// }

// func (r *AddrRepoErr) SuggestStreet(address *models.SuggestStreetRepoReq) (*models.SuggestStreetRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// func (r *AddrRepoErr) SuggestHouse(addres *models.SuggestHouseRepoReq) (*models.SuggestHouseRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// func (r *AddrRepoErr) GetCity(city string) (*models.GetCityRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// func (r *AddrRepoErr) GetStreet(street *models.GetStreetRepoReq) (*models.GetStreetRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// func (r *AddrRepoErr) GetHouse(house *models.GetHouseRepoReq) (*models.GetHouseRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// type AddrRepoHouseErr struct {
// 	mock.Mock
// }

// func (r *AddrRepoHouseErr) SuggestHouse(addres *models.SuggestHouseRepoReq) (*models.SuggestHouseRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// func (r *AddrRepoHouseErr) GetHouse(house *models.GetHouseRepoReq) (*models.GetHouseRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// func (r *AddrRepoHouseErr) SuggestStreet(address *models.SuggestStreetRepoReq) (*models.SuggestStreetRepoResp, error) {
// 	return &models.SuggestStreetRepoResp{StreetSuggests: []string{"Измайловская улица", "Измайловский проспект"}}, nil
// }

// func (r *AddrRepoHouseErr) GetCity(city string) (*models.GetCityRepoResp, error) {
// 	return &models.GetCityRepoResp{CityId: 1, Name: "Москва"}, nil
// }

// func (r *AddrRepoHouseErr) GetStreet(street *models.GetStreetRepoReq) (*models.GetStreetRepoResp, error) {
// 	return &models.GetStreetRepoResp{StreetId: 1, Name: "Измайловский проспект"}, nil
// }

// type AddrRepoGetHouseErr struct {
// 	mock.Mock
// }

// func (r *AddrRepoGetHouseErr) SuggestHouse(addres *models.SuggestHouseRepoReq) (*models.SuggestHouseRepoResp, error) {
// 	return &models.SuggestHouseRepoResp{HouseSuggests: []string{"1"}}, nil
// }

// func (r *AddrRepoGetHouseErr) GetHouse(house *models.GetHouseRepoReq) (*models.GetHouseRepoResp, error) {
// 	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
// }

// func (r *AddrRepoGetHouseErr) SuggestStreet(address *models.SuggestStreetRepoReq) (*models.SuggestStreetRepoResp, error) {
// 	return &models.SuggestStreetRepoResp{StreetSuggests: []string{"Измайловская улица", "Измайловский проспект"}}, nil
// }

// func (r *AddrRepoGetHouseErr) GetCity(city string) (*models.GetCityRepoResp, error) {
// 	return &models.GetCityRepoResp{CityId: 1, Name: "Москва"}, nil
// }

// func (r *AddrRepoGetHouseErr) GetStreet(street *models.GetStreetRepoReq) (*models.GetStreetRepoResp, error) {
// 	return &models.GetStreetRepoResp{StreetId: 1, Name: "Измайловский проспект"}, nil
// }
