package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenCdnServiceRequest struct {
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
}

func (r OpenCdnServiceRequest) Invoke(client *sdk.Client) (response *OpenCdnServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OpenCdnServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "OpenCdnService", "", "")

	resp := struct {
		*responses.BaseResponse
		OpenCdnServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OpenCdnServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type OpenCdnServiceResponse struct {
	RequestId string
}
