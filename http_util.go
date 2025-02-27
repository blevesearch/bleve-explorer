//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package main

//go:generate go-bindata-assetfs -pkg=main ./static/...
//go:generate go fmt .

import (
	"net/http"

	"github.com/gorilla/mux"
)

func staticFileRouter(r *mux.Router, static http.Handler) *mux.Router {
	// static
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		myFileHandler{static}))
	// bootstrap ui insists on loading templates from this path
	r.PathPrefix("/template/").Handler(http.StripPrefix("/template/",
		myFileHandler{static}))

	// application pages
	appPages := []string{
		"/overview",
		"/search",
		"/indexes",
		"/analysis",
		"/monitor",
	}

	for _, p := range appPages {
		// if you try to use index.html it will redirect...poorly
		r.PathPrefix(p).Handler(RewriteURL("/", static))
	}

	r.Handle("/", http.RedirectHandler("/static/index.html", 302))

	return r
}

type myFileHandler struct {
	h http.Handler
}

func (mfh myFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if *staticEtag != "" {
		w.Header().Set("Etag", *staticEtag)
	}
	mfh.h.ServeHTTP(w, r)
}

func RewriteURL(to string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = to
		h.ServeHTTP(w, r)
	})
}

func muxVariableLookup(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}

func docIDLookup(req *http.Request) string {
	return muxVariableLookup(req, "docID")
}

func indexNameLookup(req *http.Request) string {
	return muxVariableLookup(req, "indexName")
}
