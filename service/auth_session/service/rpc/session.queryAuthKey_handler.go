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
	"github.com/fedigram/fedigram-server/pkg/grpc_util"
	"github.com/fedigram/fedigram-server/pkg/logger"
	"github.com/fedigram/fedigram-server/mtproto"
	"golang.org/x/net/context"
)

// session.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (s *SessionServiceImpl) SessionQueryAuthKey(ctx context.Context, request *mtproto.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("session.queryAuthKey - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	reply, err := s.AuthSessionModel.QueryAuthKey(request.GetAuthKeyId())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	glog.Infof("session.queryAuthKey - reply: {%s}", logger.JsonDebugData(reply))
	return reply, nil
}
