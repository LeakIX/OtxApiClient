package OtxApiClient

import (
	"encoding/json"
	"net/http"
)

type userService struct {
	client *Client
}

func (u *userService) GetMe() (*User, error) {
	resp, err := u.client.doHttpRequest(http.MethodGet, "/users/me", nil)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type User struct {
	SubscriberCount int64  `json:"subscriber_count"`
	FollowerCount   int64  `json:"follower_count"`
	MemberSince     string `json:"member_since"`
	AvatarUrl       string `json:"avatar_url"`
	AwardCount      int64  `json:"award_count"`
	UserId          int64  `json:"user_id"`
	Username        string `json:"username"`
	IndicatorCount  int64  `json:"indicator_count"`
	PulseCount      int64  `json:"pulse_count"`
}
