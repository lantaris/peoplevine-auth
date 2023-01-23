package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"peoplevine-auth/logger"
)

type RestServer struct {
	RestRouter     *mux.Router
	RestServerConf TRestServerConf
}

// *******************************************************************************
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Debugln("REST call error: " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// *******************************************************************************
func notFound(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debugln("REST call error (endpoint not found): " + r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

// *******************************************************************************
func (RS RestServer) Start() error {
	var (
		err error = nil
	)
	logger.Logger.Debugln("Starting REST server: " + RS.RestServerConf.Name)
	RS.RestRouter = mux.NewRouter().StrictSlash(true)
	for _, service := range RS.RestServerConf.Endpoints {
		logger.Logger.Debugln("REST add endpoint '" + service.Name + "':[" + service.Endpoint + ":" + service.Type + "]")
		RS.RestRouter.HandleFunc(service.Endpoint, service.Callback).Methods(service.Type).Name(service.Name)
	}

	logger.Logger.Infoln("Starting rest server listen on:", RS.RestServerConf.Listen)
	RS.RestRouter.NotFoundHandler = http.HandlerFunc(notFound)
	err = http.ListenAndServe(RS.RestServerConf.Listen, RS.RestRouter)
	if err != nil {
		logger.Logger.Errorln("Error starting REST server: [" + RS.RestServerConf.Name + "]:" + err.Error())
	}
	return err
}
