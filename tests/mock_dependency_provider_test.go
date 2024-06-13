package tests

type mockDependencyProvider struct{}

func newMockProvider() *mockDependencyProvider {
	return &mockDependencyProvider{}
}
