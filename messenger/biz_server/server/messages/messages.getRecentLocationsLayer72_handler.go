/*
 *  Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package messages

import (
	"github.com/golang/glog"
	"github.com/fedigram/fedigram-server/mtproto"
	"github.com/fedigram/fedigram-server/pkg/grpc_util"
	"github.com/fedigram/fedigram-server/pkg/logger"
	"golang.org/x/net/context"
)

// messages.getRecentLocations#249431e2 peer:InputPeer limit:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetRecentLocationsLayer72(ctx context.Context, request *mtproto.TLMessagesGetRecentLocationsLayer72) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getRecentLocations#249431e2 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesSearchGlobal logic
	messages := &mtproto.TLMessagesMessages{Data2: &mtproto.Messages_Messages_Data{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}}

	glog.Infof("messages.getRecentLocations#249431e2 - reply: %s", logger.JsonDebugData(messages))
	return messages.To_Messages_Messages(), nil
}
