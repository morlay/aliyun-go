package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type TransformToPrePaidRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *TransformToPrePaidRequest) Invoke(client *sdk.Client) (resp *TransformToPrePaidResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "TransformToPrePaid", "redisa", "")
	resp = &TransformToPrePaidResponse{}
	err = client.DoAction(req, resp)
	return
}

type TransformToPrePaidResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.String
	EndTime   common.String
}
