package gorush

import (
	"net/http"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() error {
	var err error

	if PushConf.Core.SSL && PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
		err = http.ListenAndServeTLS(":"+PushConf.Core.Port, PushConf.Core.CertPath, PushConf.Core.KeyPath, routerEngine())
	} else {
		err = http.ListenAndServe(":"+PushConf.Core.Port, routerEngine())
	}

	return err
}
