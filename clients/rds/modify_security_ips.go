package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySecurityIpsRequest struct {
	requests.RpcRequest
	DBInstanceIPArrayName      string `position:"Query" name:"DBInstanceIPArrayName"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	ModifyMode                 string `position:"Query" name:"ModifyMode"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                string `position:"Query" name:"ClientToken"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	SecurityIps                string `position:"Query" name:"SecurityIps"`
	SecurityGroupId            string `position:"Query" name:"SecurityGroupId"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	WhitelistNetworkType       string `position:"Query" name:"WhitelistNetworkType"`
	DBInstanceIPArrayAttribute string `position:"Query" name:"DBInstanceIPArrayAttribute"`
	DBInstanceId               string `position:"Query" name:"DBInstanceId"`
}

func (req *ModifySecurityIpsRequest) Invoke(client *sdk.Client) (resp *ModifySecurityIpsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifySecurityIps", "rds", "")
	resp = &ModifySecurityIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityIpsResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskId    common.String
}
