var ErrInvalidUser = errors.New("user name cannot be blank")

type Hook struct{}

func (v *Hook) Exec(ctx context.Context, user *test_repo.UserModel) error {
	if user.Name == "" {
		return ErrInvalidUser
	}

	return nil
}

r := repo.New(client)
r.Events().AppendPrePutHook(&Hook{})