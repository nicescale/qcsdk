package qcsdk

import (
	"github.com/nicescale/qcsdk/types"
)

func (api *Api) DescribeZones(filters ...Params) ([]*types.Zone, error) {
	req := api.NewRequest("DescribeZones")
	mergeFilterParams(req, []string{"zones", "status"}, filters)

	ret := types.DescribeZonesResponse{}
	err := api.SendRequest(req, &ret)
	if err != nil {
		return nil, err
	}

	return ret.Zones, nil
}
