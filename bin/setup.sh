if ! command -v air > /dev/null 2>&1; then
  echo "Installing air"
  go install github.com/cosmtrek/air@latest
fi

if ! command -v gotest > /dev/null 2>&1; then
  echo "Installing gotest"
  go install github.com/rakyll/gotest@latest
fi

INSTALL_GOLANGCI_LINT=0
LATEST_GOLANGCI_LINT_VERSION=$(curl -s https://api.github.com/repos/golangci/golangci-lint/releases/latest | jq -r .tag_name)
if command -v golangci-lint > /dev/null 2>&1; then
  GOLANGCI_LINT_VERSION=v$(golangci-lint --version | awk '{print $4}')
  if [ "$GOLANGCI_LINT_VERSION" != "$LATEST_GOLANGCI_LINT_VERSION" ]; then
    echo "Warning: Your golangci-lint version is $GOLANGCI_LINT_VERSION, but the latest version is $LATEST_GOLANGCI_LINT_VERSION"

    read -p "Do you want to install the latest version ? [y/N] " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
      INSTALL_GOLANGCI_LINT=1
    fi
  fi
else
  INSTALL_GOLANGCI_LINT=1
fi

if [[ $INSTALL_GOLANGCI_LINT -eq 1 ]]; then
  echo "Installing golangci-lint"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin $LATEST_GOLANGCI_LINT_VERSION
fi
