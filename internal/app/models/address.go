package models

type SuggestReq struct {
	Address string `json:"address"`
}

//easyjson:json
type OneSuggestResp struct {
	Address string `json:"address"`
	Full    bool   `json:"end"`
}

//easyjson:json
type SuggestsResp struct {
	Suggests []OneSuggestResp `json:"suggests"`
}

type OneSuggestUcaseResp struct {
	Address string
	Full    bool
}

type SuggestsUcaseResp struct {
	Suggests []OneSuggestUcaseResp
}

type SuggestUcaseReq struct {
	Address string
	UserId  int64
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

type GetTopUserAddrsRepoReq struct {
	UserId int64
	Limit  int
}

type GetTopUserAddrsRepoResp struct {
	Addrs []string
}

type SuggestUserAddrsRepoReq struct {
	Addr       string
	StreetType string
	UserId     int64
	Limit      int
}

type SuggestUserAddrsRepoResp struct {
	Addrs []string
}
