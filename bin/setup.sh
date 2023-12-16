if ! command -v air > /dev/null 2>&1; then
  echo "Installing air"
  go install github.com/cosmtrek/air@latest
fi

if ! command -v gotest > /dev/null 2>&1; then
  echo "Installing gotest"
  go install github.com/rakyll/gotest@latest
fi

if ! command -v golangci-lint > /dev/null 2>&1; then
  echo "Installing golangci-lint"
  GOLANGCI_LINT_VERSION=$(curl -s https://api.github.com/repos/golangci/golangci-lint/releases/latest | jq -r .tag_name)
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin $GOLANGCI_LINT_VERSION
fi
