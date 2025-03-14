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

package messages

import (
	"github.com/golang/glog"
	"github.com/fedigram/fedigram-server/pkg/grpc_util"
	"github.com/fedigram/fedigram-server/pkg/logger"
	"github.com/fedigram/fedigram-server/mtproto"
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/base"
	"golang.org/x/net/context"
)

// messages.hideReportSpam#a8f1709b peer:InputPeer = Bool;
func (s *MessagesServiceImpl) MessagesHideReportSpam(ctx context.Context, request *mtproto.TLMessagesHideReportSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.hideReportSpam#a8f1709b - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.GetPeer())
	if peer.PeerType == base.PEER_USER || peer.PeerType == base.PEER_CHAT {
		// TODO(@benqi): 入库
	}

	glog.Info("messages.hideReportSpam#a8f1709b - reply: {true}")
	return mtproto.ToBool(true), nil
}
