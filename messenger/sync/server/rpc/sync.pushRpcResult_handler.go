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

package rpc

import (
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/fedigram/fedigram-server/mtproto"
    "github.com/fedigram/fedigram-server/pkg/logger"
    "github.com/fedigram/fedigram-server/mtproto/rpc"
)

// sync.pushRpcResult server_id:int auth_key_id:long req_msg_id:long result:bytes = Bool;
func (s *SyncServiceImpl) SyncPushRpcResult(ctx context.Context, request *mtproto.TLSyncPushRpcResult) (*mtproto.Bool, error) {
    glog.Infof("sync.pushRpcResult - request: {%s}", logger.JsonDebugData(request))
    authKeyId := request.GetAuthKeyId()
    clientMsgId := request.GetReqMsgId()
    cntl := zrpc.NewController()
    pushData := request.GetResult()
    serverId := request.GetServerId()
    s.pushUpdatesToSession(syncTypeRpcResult, 0, authKeyId, clientMsgId, cntl, pushData, serverId, 0, 0)
    glog.Infof("sync.pushRpcResult#1bf9b15e - reply: {true}",)
    return mtproto.ToBool(true), nil
}

