/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/andrey4d/ytapiserver/internal/log"
	"github.com/andrey4d/ytapiserver/templ/layout"
)

var PageAttributes *layout.BaseAttributes

func GetPageAttributes() *layout.BaseAttributes {
	logger := log.New()
	if PageAttributes == nil {
		PageAttributes = &layout.BaseAttributes{
			BodyClass:   "container",
			BodyId:      "main-container",
			AuthorEmail: "andrey4d.dev@gmail.com",
		}

		logger.LogModuleInfo("page_attributes").Info("new attributes")
		return PageAttributes
	}
	logger.LogModuleInfo("page_attributes").Info("reuse attributes")
	return PageAttributes
}
