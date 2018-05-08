package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBInstanceSSLRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBInstanceSSLRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceSSLResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceSSL", "rds", "")
	resp = &ModifyDBInstanceSSLResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceSSLResponse struct {
	responses.BaseResponse
	RequestId common.String
}
