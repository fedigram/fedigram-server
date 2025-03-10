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

package account

import (
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/core"
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/dal/dao"
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/dal/dao/mysql_dao"
	"github.com/golang/glog"
)

type accountsDAO struct {
	//*mysql_dao.CommonDAO
	*mysql_dao.AuthUsersDAO
	*mysql_dao.UsersDAO
	*mysql_dao.DevicesDAO
	*mysql_dao.UserNotifySettingsDAO
	*mysql_dao.UserPasswordsDAO
	*mysql_dao.UserPrivacysDAO
	*mysql_dao.ReportsDAO
	*mysql_dao.WallPapersDAO
	*mysql_dao.UsernameDAO
}

type AccountModel struct {
	dao *accountsDAO
}

func (m *AccountModel) InstallModel() {
	m.dao.AuthUsersDAO = dao.GetAuthUsersDAO(dao.DB_MASTER)
	m.dao.UsersDAO = dao.GetUsersDAO(dao.DB_MASTER)
	m.dao.DevicesDAO = dao.GetDevicesDAO(dao.DB_MASTER)
	m.dao.UserNotifySettingsDAO = dao.GetUserNotifySettingsDAO(dao.DB_MASTER)
	m.dao.UserPasswordsDAO = dao.GetUserPasswordsDAO(dao.DB_MASTER)
	m.dao.UserPrivacysDAO = dao.GetUserPrivacysDAO(dao.DB_MASTER)
	m.dao.ReportsDAO = dao.GetReportsDAO(dao.DB_MASTER)
	m.dao.WallPapersDAO = dao.GetWallPapersDAO(dao.DB_MASTER)
	m.dao.UsernameDAO = dao.GetUsernameDAO(dao.DB_MASTER)
}

func (m *AccountModel) RegisterCallback(cb interface{}) {
}

func (m *AccountModel) CheckShowStatus(selfId, userId int32, isContact bool) bool {
	logic := m.MakePrivacyLogic(selfId)
	glog.Info("selfId: ", selfId, ", seenId: ", userId, ", isContact: ", isContact, ", logic: ", logic)
	return logic.GetPrivacy(PrivacyKeyType_STATUS_TIMESTAMP).IsAllow(userId, isContact)
}

func (m *AccountModel) CheckAllowChatInvites(selfId, userId int32, isContact bool) bool {
	logic := m.MakePrivacyLogic(userId)
	return logic.GetPrivacy(PrivacyKeyType_CHAT_INVITE).IsAllow(userId, isContact)
}

func (m *AccountModel) CheckAllowCalls(selfId, userId int32, isContact bool) bool {
	logic := m.MakePrivacyLogic(userId)
	return logic.GetPrivacy(PrivacyKeyType_PHONE_CALL).IsAllow(userId, isContact)
}

func init() {
	core.RegisterCoreModel(&AccountModel{dao: &accountsDAO{}})
}
