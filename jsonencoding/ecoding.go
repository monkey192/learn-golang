package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ItemClub struct {
	ClubID            int64        `json:"clubID"`
	ClubName          string       `json:"clubName"`
	ClubOwnerID       int64        `json:"clubOwnerID"`
	ClubOwnerName     string       `json:"clubOwnerName"`
	ContractOrgID     int64        `json:"contractOrgID"`
	ContractOrgName   string       `json:"contractOrgName"`
	ClubOwnerImageURL string       `json:"clubOwnerImageUrl"`
	Description       string       `json:"description"`
	ClubImageURL      string       `json:"clubImageUrl"`
	CoverImageURL     string       `json:"coverImageUrl"`
	RewardDownloadURL string       `json:"rewardDownloadUrl"`
	ThemeColor        string       `json:"themeColor"`
	CategoryCode      int64        `json:"categoryCode"`
	CategoryName      string       `json:"categoryName"`
	ReleaseStatus     string       `json:"releaseStatus"`
	ReleaseStartDate  int64        `json:"releaseStartDate"`
	ReleaseEndDate    int64        `json:"releaseEndDate"`
	MaxMembers        int64        `json:"maxMembers"`
	ClubParticipants  int64        `json:"clubParticipants"`
	MaxMembershipNo   int64        `json:"maxMembershipNo"`
	UpdatedAt         *time.Time   `json:"updatedAt"`
}

func main() {

	dataClubResponseInfo := ItemClub{}
	var sampleResponseJSON = `
		{
			"clubID": 2826115,
			"clubName": "2022年度ペネトレーションテスト用1",
			"clubOwnerID": 124825,
			"contractOrgID": 99999,
			"contractOrgName": "",
			"clubOwnerName": "STGテストクラブオーナー2",
			"description": "",
			"clubImageUrl": "https://contents.stg.nudge-platform.com/bbb",
			"coverImageUrl": "https://contents.stg.nudge-platform.com/aaa",
			"themeColor": "#0F1F4D",
			"categoryCode": 110,
			"releaseStatus": "0",
			"releaseStartDate": 20220531,
			"releaseEndDate": 99991231,
			"maxMembers": 100,
			"clubParticipants": 8,
			"maxMembershipNo": 8,
			"createdAt": "2022-05-31 15:50:21",
			"updatedAt": "2022-07-12 15:03:55",
			"createdPgmID": "API300501",
			"updatedPgmID": "API300503"
			}
	`

	sampleResponseJSONByte := []byte(sampleResponseJSON)

	err := json.Unmarshal(sampleResponseJSONByte, &dataClubResponseInfo)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(dataClubResponseInfo)
	}
}