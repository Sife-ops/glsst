package lib

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func Onboard(user User) error {
	uq, err := Query(User{
		UserId: user.UserId,
	})
	if err != nil {
		return err
	}

	switch len(uq.Items) < 1 {
	case true:
		if _, err := Put(user); err != nil {
			return err
		}
	default:
		var ul []User
		if err := attributevalue.UnmarshalListOfMaps(uq.Items, &ul); err != nil {
			return err
		}
		u := ul[0]

		switch {
		case u.Avatar != user.Avatar:
			fallthrough
		case u.DisplayName != user.DisplayName:
			fallthrough
		case u.Username != user.Username:
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
