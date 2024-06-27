.PHONY: test fuzz lint

lint:
	golangci-lint run -c .golangci.yml ./...

test:
	go test -coverprofile=profile.out -coverpkg=github.com/krazybee/goldmark,github.com/krazybee/goldmark/ast,github.com/krazybee/goldmark/extension,github.com/krazybee/goldmark/extension/ast,github.com/krazybee/goldmark/parser,github.com/krazybee/goldmark/renderer,github.com/krazybee/goldmark/renderer/html,github.com/krazybee/goldmark/text,github.com/krazybee/goldmark/util ./...

cov: test
	go tool cover -html=profile.out

fuzz:
	cd ./fuzz && go test -fuzz=Fuzz
