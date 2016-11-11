package decorators

import (
	"context"

	model "github.com/lyft/lyft-idl/generated/go/pb/lyft/test_repo"
	test_repo "github.com/lyft/lyft-idl/generated/go/pb/lyft/test_repo/test_repo/mongodb"
)

func Rename() test_repo.PutDecorator {
	return test_repo.NewPutDecorator(
		func(ctx context.Context, user *model.UserModel, next test_repo.PutExecFunc) error {
			user.Name = user.Name + "_pre"
			err := next(ctx, user)
			user.Name = user.Name + "_post"

			return err
		},
	)
}

err := repo.Put(ctx, &user).Decorate(Rename()).Response()
