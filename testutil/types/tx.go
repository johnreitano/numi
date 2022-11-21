package types

import "github.com/johnreitano/numi/x/numi/types"

const (
	Alice  = "numi12utcwzwmrwzsjj362qsht6zapq9cd5q0gh8pgy"
	Bob    = "numi1hq9ujvq7yl2krs54xez67pfhfkpg8nlpqa75u2"
	Oliver = "numi13crpqdukn5l3gr4pzzcjzcl6fpx7rhay8uvy44"
	Olivia = "numi1tsacr4aqrrerakdlcmlzl7daplle54fj874w2s"
)

func ValidMsgCreateAndVerifyUser() *types.MsgCreateAndVerifyUser {
	return &types.MsgCreateAndVerifyUser{
		Creator:           Olivia,
		UserId:            "1bc3e020-2b02-40a7-abd8-eadc9b4250c5",
		FirstName:         "John",
		LastName:          "Doe",
		CountryCode:       "USA",
		SubnationalEntity: "California",
		City:              "San Diego",
		Bio:               "a serious man",
		Referrer:          Alice,
		AccountAddress:    Bob,
	}
}

func ValidUser() *types.User {
	return &types.User{
		Creator:           Olivia,
		UserId:            "1bc3e020-2b02-40a7-abd8-eadc9b4250c5",
		FirstName:         "John",
		LastName:          "Doe",
		CountryCode:       "USA",
		SubnationalEntity: "California",
		City:              "San Diego",
		Bio:               "a serious man",
		Referrer:          Alice,
		AccountAddress:    Bob,
	}
}
