package anilistintegration

import (
	"github.com/metafates/mangal/constant"
	"github.com/spf13/viper"
)

type Anilist struct {
	token string
}

func New() *Anilist {
	return &Anilist{}
}

func (a *Anilist) id() string {
	return viper.GetString(constant.AnilistID)
}

func (a *Anilist) secret() string {
	return viper.GetString(constant.AnilistSecret)
}

func (a *Anilist) code() string {
	return viper.GetString(constant.AnilistCode)
}

func (a *Anilist) AuthURL() string {
	return "https://anilist.co/api/v2/oauth/authorize?client_id=" + a.id() + "&response_type=code&redirect_uri=https://anilist.co/api/v2/oauth/pin"
}
