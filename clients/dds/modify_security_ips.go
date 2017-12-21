package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityIpsRequest struct {
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ModifyMode               string `position:"Query" name:"ModifyMode"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	SecurityIps              string `position:"Query" name:"SecurityIps"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	SecurityIpGroupName      string `position:"Query" name:"SecurityIpGroupName"`
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	SecurityIpGroupAttribute string `position:"Query" name:"SecurityIpGroupAttribute"`
}

func (r ModifySecurityIpsRequest) Invoke(client *sdk.Client) (response *ModifySecurityIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySecurityIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifySecurityIps", "dds", "")

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
