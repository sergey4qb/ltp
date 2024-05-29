package kraken_client

import "fmt"

func errUnexpectedCode(code int) error {
	return fmt.Errorf("unexpected status code: %d", code)
}

func errThirdPartyError(errString string) error {
	return fmt.Errorf("third party api error: %s", errString)
}
