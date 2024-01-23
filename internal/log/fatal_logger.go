/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package log

import "github.com/sirupsen/logrus"

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
