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
    "golang.org/x/net/context"
    "github.com/fedigram/fedigram-server/pkg/grpc_util"
    "github.com/fedigram/fedigram-server/pkg/logger"
    "github.com/fedigram/fedigram-server/mtproto"
    "github.com/fedigram/fedigram-server/messenger/biz_server/biz/core"
    "github.com/fedigram/fedigram-server/messenger/biz_server/biz/core/message"
    "github.com/fedigram/fedigram-server/messenger/biz_server/biz/core/update"
    "github.com/fedigram/fedigram-server/messenger/biz_server/biz/base"
)

// messages.importChatInvite#6c50051c hash:string = Updates;
func (s *MessagesServiceImpl) MessagesImportChatInvite(ctx context.Context, request *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("messages.importChatInvite - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    chatLogic, err := s.ChatModel.NewChatLogicByLink(request.GetHash())
    if err != nil {
        // TODO(@benqi): do chat checkChatInvite
        glog.Errorf("messages.importChatInvite#6c50051c - error: {%v}", err)
        return nil, err
    }
    peer := &base.PeerUtil{
        PeerType: base.PEER_CHAT,
        PeerId:   chatLogic.GetChatId(),
    }

    chatLogic.AddChatUser(chatLogic.GetCreator(), md.UserId)

    addUserMessage := chatLogic.MakeJoinedByLinkMessage(chatLogic.GetCreator())
    randomId := core.GetUUID()

    resultCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (*mtproto.Updates, error) {
        syncUpdates := updates.NewUpdatesLogic(md.UserId)

        updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
            Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
        }}
        syncUpdates.AddUpdate(updateChatParticipants.To_Update())
        syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(outBox.OwnerId))
        syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chatLogic.GetChatParticipantIdList()))
        syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

        syncUpdates.AddUpdateMessageId(outBox.MessageId, outBox.RandomId)

        return syncUpdates.ToUpdates(), nil
    }

    syncNotMeCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (int64, *mtproto.Updates, error) {
        syncUpdates := updates.NewUpdatesLogic(md.UserId)

        updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
            Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
        }}
        syncUpdates.AddUpdate(updateChatParticipants.To_Update())
        syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(outBox.OwnerId))
        syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chatLogic.GetChatParticipantIdList()))
        syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

        return md.AuthId, syncUpdates.ToUpdates(), nil
    }

    pushCB := func(pts, ptsCount int32, inBox *message.MessageBox2) (*mtproto.Updates, error) {
        pushUpdates := updates.NewUpdatesLogic(md.UserId)

        updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
            Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
        }}
        pushUpdates.AddUpdate(updateChatParticipants.To_Update())
        pushUpdates.AddUpdateNewMessage(pts, ptsCount, inBox.ToMessage(inBox.OwnerId))
        pushUpdates.AddUsers(s.UserModel.GetUserListByIdList(inBox.OwnerId, chatLogic.GetChatParticipantIdList()))
        pushUpdates.AddChat(chatLogic.ToChat(inBox.OwnerId))

        return pushUpdates.ToUpdates(), nil
    }

    replyUpdates, _ := s.MessageModel.SendMessage(
        md.UserId,
        peer,
        randomId,
        addUserMessage,
        resultCB,
        syncNotMeCB,
        pushCB)

    glog.Infof("messages.importChatInvite - reply: {%v}", replyUpdates)
    return replyUpdates, nil
}
