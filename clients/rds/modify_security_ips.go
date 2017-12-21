package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityIpsRequest struct {
	DBInstanceIPArrayName      string `position:"Query" name:"DBInstanceIPArrayName"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	ModifyMode                 string `position:"Query" name:"ModifyMode"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                string `position:"Query" name:"ClientToken"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	SecurityIps                string `position:"Query" name:"SecurityIps"`
	DBInstanceIPArrayAttribute string `position:"Query" name:"DBInstanceIPArrayAttribute"`
	DBInstanceId               string `position:"Query" name:"DBInstanceId"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
}

func (r ModifySecurityIpsRequest) Invoke(client *sdk.Client) (response *ModifySecurityIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySecurityIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifySecurityIps", "rds", "")

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
	TaskId    string
}
