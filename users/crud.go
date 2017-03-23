package users

import (
	"context"

	"github.com/ccutch/data-pit/models"
	"github.com/ccutch/datastore-model"
)

func GetListOfUsers(ctx context.Context) interface{} {
	var us []*models.User

	query := db.From(new(models.User))
	db.NewDatastore(ctx).Query(query).All(&us)
	res := []*UserResponse{}
	for _, u := range us {
		res = append(res, NewUserResponse(ctx, u))
	}

	return res
}

func CreateUser(ctx context.Context) interface{} {
	return new(UserResponse)
}
