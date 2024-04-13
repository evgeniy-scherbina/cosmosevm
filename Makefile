mock-expected-keepers:
	mockgen -source=x/monobank/types/expected_keepers.go \
		-package testutil \
		-destination=x/monobank/testutil/expected_keepers_mocks.go

	mockgen -source=x/privatbank/types/expected_keepers.go \
		-package testutil \
		-destination=x/privatbank/testutil/expected_keepers_mocks.go
