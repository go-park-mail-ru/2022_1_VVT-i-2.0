package address

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Repository interface {
	SuggestStreet(address *models.SuggestStreetRepoReq) (*models.SuggestStreetRepoResp, error)
	SuggestHouse(addres *models.SuggestHouseRepoReq) (*models.SuggestHouseRepoResp, error)
	GetCity(city *models.GetCityRepoReq) (*models.GetCityRepoResp, error)
	GetStreet(street *models.GetStreetRepoReq) (*models.GetStreetRepoResp, error)
	GetHouse(house *models.GetHouseRepoReq) (*models.GetHouseRepoResp, error)
	GetTopUserAddrs(req *models.GetTopUserAddrsRepoReq) (*models.GetTopUserAddrsRepoResp, error)
	SuggestUserAddrs(req *models.SuggestUserAddrsRepoReq) (*models.SuggestUserAddrsRepoResp, error)
}
