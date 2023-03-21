mock-ver:
	./mockery --version

mockgen-habits:
	# service mocks
	./mockery \
		--dir=../../internal/habits/service \
		--name=Store \
		--filename=service_mock.go \
		--output=../mocks/habits/service/mocks \
		--outpkg=habits_test_mock \
