package oauth

import (
	"log"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var Manager = manage.NewDefaultManager()
var ClientStore = store.NewClientStore()
var Server = server.NewDefaultServer(Manager)

func init() {
	Manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	Manager.MustTokenStorage(store.NewMemoryTokenStore())
	Manager.MapClientStorage(ClientStore)

	Server.SetAllowGetAccessRequest(true)
	Server.SetClientInfoHandler(server.ClientFormHandler)
	Manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	Server.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
}
