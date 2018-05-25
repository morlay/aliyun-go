package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBInstanceProxyConfigurationRequest struct {
	requests.RpcRequest
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	ProxyConfigurationKey   string `position:"Query" name:"ProxyConfigurationKey"`
	ProxyConfigurationValue string `position:"Query" name:"ProxyConfigurationValue"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBInstanceProxyConfigurationRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceProxyConfigurationResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceProxyConfiguration", "rds", "")
	resp = &ModifyDBInstanceProxyConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceProxyConfigurationResponse struct {
	responses.BaseResponse
	RequestId common.String
}
