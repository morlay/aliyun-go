package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCdnServiceRequest struct {
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyCdnServiceRequest) Invoke(client *sdk.Client) (response *ModifyCdnServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCdnServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyCdnService", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCdnServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCdnServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCdnServiceResponse struct {
	RequestId string
}
