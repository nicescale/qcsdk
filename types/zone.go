package types

type Zone struct {
	Status string `json:"status"`
	ZoneID string `json:"zone_id"`
}

type DescribeZonesResponse struct {
	ResponseStatus
	Total int     `json:"total_count"`
	Zones []*Zone `json:"zone_set"`
}
