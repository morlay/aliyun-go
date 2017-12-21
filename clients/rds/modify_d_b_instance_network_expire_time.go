package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceNetworkExpireTimeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	ClassicExpiredDays   int    `position:"Query" name:"ClassicExpiredDays"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBInstanceNetworkExpireTimeRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceNetworkExpireTimeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceNetworkExpireTime", "rds", "")
	resp = &ModifyDBInstanceNetworkExpireTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceNetworkExpireTimeResponse struct {
	responses.BaseResponse
	RequestId string
}
