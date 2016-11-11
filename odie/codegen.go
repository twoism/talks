func (b *QueryBuilder) Where() *expression.Builder {
	return b.bldr.Where()
}

func (b *QueryBuilder) Into(r *[]users.UserModel) *QueryBuilder {
	b.target = r
	b.bldr.Into(b.Target())

	return b
}