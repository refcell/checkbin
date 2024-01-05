# Help Commands
help-me: help
help-me-please: help
help-me-please-i-am-drowning: help
help:
  just -l

# Build Commands
build:
  @go build -o bin/checkbin .

# Test Commands
test:
  @go test -v ./...

# Test the checkbin binary checksum using checkbin
check-me:
  @just build
  @./bin/checkbin ./bin/checkbin
