func (b *GetBuilder) Response() error {
	err := b.decorated().Exec(b.bldr.Context(), b.Target())

	if b.Target() != nil && err == nil {
		b.Target().Commit() // HL
	}

	return err
}