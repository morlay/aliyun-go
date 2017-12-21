package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityIpsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	SecurityIps          string `position:"Query" name:"SecurityIps"`
}

func (r ModifySecurityIpsRequest) Invoke(client *sdk.Client) (response *ModifySecurityIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySecurityIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "ModifySecurityIps", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifySecurityIpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySecurityIpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySecurityIpsResponse struct {
	RequestId string
}
