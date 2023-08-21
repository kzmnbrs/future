//go:build tools
// +build tools

package future

//go:generate go install github.com/vektra/mockery/v2@v2.32.4
//go:generate mockery --name=Future --output=. --outpkg=future --filename=future_mock.go --structname=Mock
