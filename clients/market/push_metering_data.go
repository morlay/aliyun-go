package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PushMeteringDataRequest struct {
	requests.RpcRequest
	Metering string `position:"Query" name:"Metering"`
}

func (req *PushMeteringDataRequest) Invoke(client *sdk.Client) (resp *PushMeteringDataResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "PushMeteringData", "yunmarket", "")
	resp = &PushMeteringDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushMeteringDataResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
