// Copyright 2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dolthub/go-mysql-server/sql/config"
)

func TestConfigWithDefaults(t *testing.T) {
	defaults := config.NewMapConfig(map[string]string{
		"max_connections":   "1000",
		"net_write_timeout": "1",
		"net_read_timeout":  "1",
	})
	serverConf := Config{}
	serverConf, err := serverConf.WithDefaults(defaults)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1000), serverConf.MaxConnections)
	assert.Equal(t, time.Duration(1000000), serverConf.ConnReadTimeout)
	assert.Equal(t, time.Duration(1000000), serverConf.ConnWriteTimeout)
}
