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

package payments

import (
    "fmt"
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/fedigram/fedigram-server/pkg/grpc_util"
    "github.com/fedigram/fedigram-server/pkg/logger"
    "github.com/fedigram/fedigram-server/mtproto"
)

// payments.sendPaymentForm#2b8879b3 flags:# msg_id:int requested_info_id:flags.0?string shipping_option_id:flags.1?string credentials:InputPaymentCredentials = payments.PaymentResult;
func (s *PaymentsServiceImpl) PaymentsSendPaymentForm(ctx context.Context, request *mtproto.TLPaymentsSendPaymentForm) (*mtproto.Payments_PaymentResult, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("payments.sendPaymentForm - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // Sorry: not impl PaymentsSendPaymentForm logic
    glog.Warning("payments.sendPaymentForm blocked, License key from https://nebula.chat required to unlock enterprise features.")

    return nil, fmt.Errorf("not imp PaymentsSendPaymentForm")
}
