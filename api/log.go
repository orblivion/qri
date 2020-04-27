package api

import (
	"fmt"
	"net/http"

	util "github.com/qri-io/apiutil"
	"github.com/qri-io/qri/lib"
	"github.com/qri-io/qri/logbook"
	"github.com/qri-io/qri/repo"
)

// LogHandlers wraps a LogMethods with http.HandlerFuncs
type LogHandlers struct {
	lm       lib.LogMethods
	rm       lib.RemoteMethods
	readOnly bool
}

// NewLogHandlers allocates a LogHandlers pointer
func NewLogHandlers(inst *lib.Instance) *LogHandlers {
	req := lib.NewLogMethods(inst)
	rem := lib.NewRemoteMethods(inst)
	h := LogHandlers{
		lm: *req,
		rm: *rem,
	}
	return &h
}

// LogHandler is the endpoint for dataset logs
func (h *LogHandlers) LogHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		util.EmptyOkHandler(w, r)
	case "GET":
		h.logHandler(w, r)
	default:
		util.NotFoundHandler(w, r)
	}
}

// DatasetLogItem aliases the type from logbook
type DatasetLogItem = logbook.DatasetLogItem

func (h *LogHandlers) logHandler(w http.ResponseWriter, r *http.Request) {
	args, err := DatasetRefFromPath(r.URL.Path[len("/history"):])
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if args.Name == "" && args.Path == "" {
		util.WriteErrResponse(w, http.StatusBadRequest, fmt.Errorf("name of dataset or path needed"))
		return
	}

	lp := lib.ListParamsFromRequest(r)
	lp.Peername = args.Peername

	local := r.FormValue("local") == "true"
	remoteName := r.FormValue("remote")

	if local && remoteName != "" {
		util.WriteErrResponse(w, http.StatusBadRequest, fmt.Errorf("cannot use the 'local' and 'remote' params at the same time"))
		return
	}

	res := []DatasetLogItem{}
	if remoteName == "" {
		params := &lib.LogParams{
			Ref:        args.String(),
			ListParams: lp,
		}
		if err := h.lm.Log(params, &res); err != nil {
			if err == repo.ErrNoHistory {
				util.WriteErrResponse(w, http.StatusUnprocessableEntity, err)
				return
			}
			if err != repo.ErrNotFound {
				util.WriteErrResponse(w, http.StatusInternalServerError, err)
				return
			}
			if local && err == repo.ErrNotFound {
				util.WriteErrResponse(w, http.StatusInternalServerError, err)
				return
			}
		} else {
			if err := util.WritePageResponse(w, res, r, params.Page()); err != nil {
				log.Infof("error list dataset history response: %s", err.Error())
			}
			return
		}
	}
	p := &lib.FetchParams{
		Ref:        args.String(),
		RemoteName: remoteName,
	}
	if err := h.rm.Fetch(p, &res); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	util.WritePageResponse(w, res, r, lp.Page())
}
