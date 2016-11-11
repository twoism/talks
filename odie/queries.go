var users []test_repo.User
err := repo.Query(ctx).Into(&users).
	Where().Attr("name").BeginsWith("c").Limit(10).Response()

err := repo.Query(ctx).Into(&users).
	Where().Attr("_id").In(1, 2, 3).Limit(3).Response()
