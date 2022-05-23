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

type SuggestStreetRepoReq struct {
	CityId         int64
	Street         string
	StreetType     string
	SuggsLimit     int
	SearchInMiddle bool
}

type SuggestStreetRepoResp struct {
	StreetSuggests []string
}

type SuggestHouseRepoReq struct {
	CityId     int64
	StreetId   int64
	House      string
	SuggsLimit int
}

type SuggestHouseRepoResp struct {
	HouseSuggests []string
}

type GetStreetRepoReq struct {
	CityId     int64
	Street     string
	StreetType string
}

type GetStreetRepoResp struct {
	StreetId int64 `db:"id"`
	Name     string
}

type GetCityRepoReq struct {
	City string
}

type GetCityRepoResp struct {
	CityId int64  `db:"id"`
	Name   string `db:"name"`
}

type GetHouseRepoReq struct {
	StreetId int64
	House    string
}

type GetHouseRepoResp struct {
	House string
}
