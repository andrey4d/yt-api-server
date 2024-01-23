/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package yaml

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ParseFile(filename string, config interface{}) error {

	cfg_data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	yaml.Unmarshal(cfg_data, config)

	return nil
}
