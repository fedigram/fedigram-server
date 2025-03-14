// Copyright (c) 2019-present,  NebulaChat Studio (https://nebula.chat).
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

package photos

import (
	"github.com/golang/glog"
	"github.com/fedigram/fedigram-server/mtproto"
	"github.com/fedigram/fedigram-server/pkg/grpc_util"
	"github.com/fedigram/fedigram-server/pkg/logger"
	"github.com/fedigram/fedigram-server/service/document/client"
	"golang.org/x/net/context"
	"time"
)

// photos.uploadProfilePhoto#d50f9c88 file:InputFile caption:string geo_point:InputGeoPoint crop:InputPhotoCrop = photos.Photo;
func (s *PhotosServiceImpl) PhotosUploadProfilePhotoLayer46(ctx context.Context, request *mtproto.TLPhotosUploadProfilePhotoLayer46) (*mtproto.Photos_Photo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("photos.uploadProfilePhoto#d50f9c88 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	file := request.GetFile()
	// uuid := helper.NextSnowflakeId()

	result, err := document_client.UploadPhotoFile(md.AuthId, file) // uuid, file.GetData2().GetId(), file.GetData2().GetParts(), file.GetData2().GetName(), file.GetData2().GetMd5Checksum())
	if err != nil {
		glog.Errorf("uploadPhoto error: %v", err)
		return nil, err
	}

	s.UserModel.SetUserPhotoID(md.UserId, result.PhotoId)

	// TODO(@benqi): sync update userProfilePhoto

	// fileData := mediaData.GetFile().GetData2()
	photo := &mtproto.TLPhotoLayer86{Data2: &mtproto.Photo_Data{
		Id:          result.PhotoId,
		HasStickers: false,
		AccessHash:  result.AccessHash, //photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
		Date:        int32(time.Now().Unix()),
		Sizes:       result.SizeList,
	}}

	photos := &mtproto.TLPhotosPhoto{Data2: &mtproto.Photos_Photo_Data{
		Photo: photo.To_Photo(),
		Users: []*mtproto.User{},
	}}

	//updateUserPhoto := &mtproto.TLUpdateUserPhoto{Data2: &mtproto.Update_Data{
	//	UserId:   md.UserId,
	//	Date:     int32(time.Now().Unix()),
	//	Photo:    photo2.MakeUserProfilePhoto(result.PhotoId, result.SizeList),
	//	Previous: mtproto.ToBool(false),
	//}}
	// sync_client.GetSyncClient().PushToUserUpdateShortData(md.UserId, updateUserPhoto.To_Update())

	glog.Infof("photos.uploadProfilePhoto#d50f9c88 - reply: %s", logger.JsonDebugData(photos))
	return photos.To_Photos_Photo(), nil
}
