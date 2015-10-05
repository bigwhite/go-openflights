package flights

type serverClient struct {
	*IDStore
	*CodeStore
}

func newServerClient(idStore *IDStore, codeStore *CodeStore) (*serverClient, error) {
	return &serverClient{
		idStore,
		codeStore,
	}, nil
}
