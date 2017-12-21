package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenCdnServiceRequest struct {
	requests.RpcRequest
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
}

func (req *OpenCdnServiceRequest) Invoke(client *sdk.Client) (resp *OpenCdnServiceResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "OpenCdnService", "", "")
	resp = &OpenCdnServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type OpenCdnServiceResponse struct {
	responses.BaseResponse
	RequestId string
}
