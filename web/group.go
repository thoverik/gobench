package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/thoverik/gobench/ent"
	"github.com/thoverik/gobench/ent/group"
)

func groupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groupID, err := strconv.Atoi(chi.URLParam(r, "groupID"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		group, err := db.Group.
			Query().
			Where(group.ID(groupID)).
			Only(r.Context())

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), webKey("group"), group)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func listGroups(w http.ResponseWriter, r *http.Request) {
	gs, err := db.Group.
		Query().
		All(r.Context())
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.RenderList(w, r, newGroupListResponse(gs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	group, ok := ctx.Value(webKey("group")).(*ent.Group)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	if err := render.Render(w, r, newGroupResponse(group)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func getGroupGraphs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	group, ok := ctx.Value(webKey("group")).(*ent.Group)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	graphs, err := group.
		QueryGraphs().
		All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.RenderList(w, r, newGraphListResponse(graphs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
