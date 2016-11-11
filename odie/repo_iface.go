type IFace interface {
	Events() *Events
	Get(ctx context.Context, id uint64) *GetBuilder
	Put(ctx context.Context, target *users.UserModel) *PutBuilder
	Delete(ctx context.Context) *DeleteBuilder
	Update(ctx context.Context) *UpdateBuilder
	Query(ctx context.Context) *QueryBuilder
}
