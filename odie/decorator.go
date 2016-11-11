func LogThings(logger logger.Logger) decorators.IFace {
	return func(b decorators.ExecIFace) decorators.ExecIFace {
		return decorators.ExecFunc(func(ctx context.Context, in interface{}) error {
			logger.Infof(ctx, "%+v", in)
			return b.Exec(ctx, in)
		})
	}
}

client.Decorate(LogThings(logger.DefaultLogger()))