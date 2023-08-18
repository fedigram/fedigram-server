// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
// Copyright (c) 2023-present, Fedigram Team. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package main

import (
	"flag"
	_ "github.com/fedigram/fedigram-server/mtproto"
	"github.com/fedigram/fedigram-server/pkg/util"
	"github.com/fedigram/fedigram-server/service/document/service"
)

func main() {
	flag.Parse()

	instance := service.NewDocumentServer()
	util.DoMainAppInstance(instance)
}
