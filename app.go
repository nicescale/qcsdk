package qcsdk

import (
	"github.com/nicescale/qcsdk/types"
)

const defaultAppEndpoint = "https://api.qingcloud.com/app/"

type App struct{ *Api }

func NewApp(ak, sk string) *App {
	api := NewApi(ak, sk, "")
	api.SetEndpoint(defaultAppEndpoint)
	return &App{api}
}

func (app *App) DescribeUsers() ([]*types.User, error) {
	req := app.NewRequest("DescribeUsers")
	ret := types.DescribeUsersResponse{}
	err := app.SendRequest(req, &ret)
	if err != nil {
		return nil, err
	}

	return ret.Users, nil
}
