package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyCdnServiceRequest struct {
	requests.RpcRequest
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyCdnServiceRequest) Invoke(client *sdk.Client) (resp *ModifyCdnServiceResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyCdnService", "", "")
	resp = &ModifyCdnServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCdnServiceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
