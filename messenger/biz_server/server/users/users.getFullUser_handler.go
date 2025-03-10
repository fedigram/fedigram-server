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

package users

import (
	"github.com/golang/glog"
	"github.com/fedigram/fedigram-server/pkg/grpc_util"
	"github.com/fedigram/fedigram-server/pkg/logger"
	"github.com/fedigram/fedigram-server/mtproto"
	"golang.org/x/net/context"
)

// users.getFullUser#ca30a5b1 id:InputUser = UserFull;
func (s *UsersServiceImpl) UsersGetFullUser(ctx context.Context, request *mtproto.TLUsersGetFullUser) (*mtproto.UserFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("users.getFullUser#ca30a5b1 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		peerId int32
		id = request.GetId()
	)

	switch id.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		peerId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		peerId = id.GetData2().GetUserId()
	default:
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PEER_ID_INVALID)
		glog.Error(err)
		return nil, err
	}

	userFull := s.UserModel.GetUserFull(md.UserId, peerId)
	glog.Infof("users.getFullUser#ca30a5b1 - reply: %s", logger.JsonDebugData(userFull))
	return userFull, nil
}
