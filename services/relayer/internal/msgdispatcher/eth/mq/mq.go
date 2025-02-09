package mq

/*
 * Copyright 2021 ConsenSys Software Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import "github.com/consensys/gpact/messaging/relayer/internal/messages"

// Handlers store all mq handler.
var Handlers map[string]map[string]func(msg messages.Message)

// init inits the Handlers.
func init() {
	Handlers = make(map[string]map[string]func(msg messages.Message))
	initV1()
}
