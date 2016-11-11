type QueryBuilder struct {
	bldr *query.Builder

	target *[]users.UserModel
	events *Events

	decorators []QueryDecorator
}