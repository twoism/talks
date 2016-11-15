func (s *Server) Get(ctx context.Context, req *test_repo.GetRequest) (*test_repo.User, error) {
	user, err := s.cfg.UserRepo().GetByID(ctx, req.Id)

	if err != nil {
		if err == repo.ErrNotFound {
			return nil, ErrUserNotFound
		}

		return nil, wrapErrorIf(err)
	}

	return user.ToProto(), nil // HL
}