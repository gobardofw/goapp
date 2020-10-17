package bootstrap

import (
	ccrypto "github.com/gobardofw/console/crypto"
	"github.com/gobardofw/crypto"
)

func init() {
	conf := app.Config()
	if c := crypto.NewCryptography(conf.String("key", "")); c != nil {
		_container.Register("--APP-CRYPTO", c)
	} else {
		panic("failed to build crypto driver")
	}
	_cli.AddCommand(ccrypto.HashCommand(func(driver string) crypto.Crypto {
		return app.Crypto(driver)
	}, "--APP-CRYPTO"))
}
