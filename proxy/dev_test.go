//+build dev

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package proxy

import (
	"testing"

	"github.com/LingmoOS/go-gir-generator/glib-2.0"
	"github.com/LingmoOS/golang-github-lingmo-go-lib/log"
)

func TestMain(t testing.T) {
	logger.SetLogLevel(log.LevelDebug)
	SetupProxy()
	logger.Info("start loop...")
	glib.StartLoop()
}
