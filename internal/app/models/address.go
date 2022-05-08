package models

type SuggestReq struct {
	Address string `json:"address" valid:"address,required"`
}

type SuggestResp struct {
	Suggests    []string `json:"suggests"`
	AddressFull bool     `json:"end"`
}

type SuggestUcaseReq struct {
	Address string
}

type SuggestUcaseResp struct {
	Suggests    []string
	AddressFull bool
}

type SuggestStreetRepoInput struct {
	CityId int64
	Street string
}

type SuggestStreetRepoAnsw struct {
	StreetSuggests []string
}

type SuggestHouseRepoInput struct {
	CityId   int64
	StreetId int64
	House    string
}

type SuggestHouseRepoAnsw struct {
	HouseSuggests []string
}

type GetStreetRepoInput struct {
	CityId int64
	Street string
}

type GetStreetRepoAnsw struct {
	StreetId int64
	Name     string
}

type GetCityRepoAnsw struct {
	CityId int64
	Name   string
}

type GetHouseRepoInput struct {
	StreetId int64
	House    string
}

type GetHouseRepoAnsw struct {
	House string
}
