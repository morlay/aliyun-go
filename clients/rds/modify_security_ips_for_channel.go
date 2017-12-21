package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityIpsForChannelRequest struct {
	requests.RpcRequest
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

func (req *ModifySecurityIpsForChannelRequest) Invoke(client *sdk.Client) (resp *ModifySecurityIpsForChannelResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifySecurityIpsForChannel", "rds", "")
	resp = &ModifySecurityIpsForChannelResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityIpsForChannelResponse struct {
	responses.BaseResponse
	RequestId string
}
