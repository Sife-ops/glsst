package lib

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func Onboard(u User) error {
	uq, err := Query(User{
		UserId: u.UserId,
	})
	if err != nil {
		return err
	}

	switch len(uq.Items) < 1 {
	case true:
		if _, err := Put(u); err != nil {
			return err
		}
	default:
		var ul []User
		if err := attributevalue.UnmarshalListOfMaps(uq.Items, &ul); err != nil {
			return err
		}
		u := ul[0]

		switch {
		case u.Avatar != u.Avatar:
			fallthrough
		case u.DisplayName != u.DisplayName:
			fallthrough
		case u.Username != u.Username:
			if _, err := Update(u, map[string]interface{}{
				"avatar":      u.Avatar,
				"displayname": u.DisplayName,
				"username":    u.Username,
			}); err != nil {
				return err
			}
		}
	}

	return nil
}
