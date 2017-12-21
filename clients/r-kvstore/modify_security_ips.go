package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityIpsRequest struct {
	requests.RpcRequest
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ModifyMode               string `position:"Query" name:"ModifyMode"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	SecurityIps              string `position:"Query" name:"SecurityIps"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	SecurityIpGroupName      string `position:"Query" name:"SecurityIpGroupName"`
	InstanceId               string `position:"Query" name:"InstanceId"`
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	SecurityIpGroupAttribute string `position:"Query" name:"SecurityIpGroupAttribute"`
}

func (req *ModifySecurityIpsRequest) Invoke(client *sdk.Client) (resp *ModifySecurityIpsResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifySecurityIps", "redisa", "")
	resp = &ModifySecurityIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityIpsResponse struct {
	responses.BaseResponse
	RequestId string
}
