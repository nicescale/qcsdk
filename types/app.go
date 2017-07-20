package types

import (
	"time"
)

type User struct {
	NotifyEmail   string    `json:"notify_email"`
	UserType      int       `json:"user_type"`
	GravatarEmail string    `json:"gravatar_email"`
	Zones         []string  `json:"zones"`
	Currency      string    `json:"currency"`
	CreateTime    time.Time `json:"create_time"`
	CompanyPhone  string    `json:"company_phone"`
	UserID        string    `json:"user_id"`
	CompanyName   string    `json:"company_name"`
	UserName      string    `json:"user_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	Lang          string    `json:"lang"`
	Gender        string    `json:"gender"`
}

type DescribeUsersResponse struct {
	ResponseStatus
	TotalCount int     `json:"total_count"`
	Users      []*User `json:"user_set"`
}
