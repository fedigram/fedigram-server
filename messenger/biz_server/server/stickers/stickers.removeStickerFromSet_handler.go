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

package stickers

import (
    "fmt"
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/fedigram/fedigram-server/pkg/grpc_util"
    "github.com/fedigram/fedigram-server/pkg/logger"
    "github.com/fedigram/fedigram-server/mtproto"
)

// stickers.removeStickerFromSet#f7760f51 sticker:InputDocument = messages.StickerSet;
func (s *StickersServiceImpl) StickersRemoveStickerFromSet(ctx context.Context, request *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("stickers.removeStickerFromSet - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // Sorry: not impl StickersRemoveStickerFromSet logic
    glog.Warning("stickers.removeStickerFromSet blocked, License key from https://nebula.chat required to unlock enterprise features.")

    return nil, fmt.Errorf("not imp StickersRemoveStickerFromSet")
}
