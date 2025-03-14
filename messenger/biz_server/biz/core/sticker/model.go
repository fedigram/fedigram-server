// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
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

package sticker

import (
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/core"
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/dal/dao"
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/dal/dao/mysql_dao"
)

type stickersDAO struct {
	*mysql_dao.StickerPacksDAO
	*mysql_dao.StickerSetsDAO
}

type StickerModel struct {
	dao *stickersDAO
}

func (m *StickerModel) InstallModel() {
	m.dao.StickerPacksDAO = dao.GetStickerPacksDAO(dao.DB_MASTER)
	m.dao.StickerSetsDAO = dao.GetStickerSetsDAO(dao.DB_MASTER)
}

func (m *StickerModel) RegisterCallback(cb interface{}) {
}

func init() {
	core.RegisterCoreModel(&StickerModel{dao: &stickersDAO{}})
}
