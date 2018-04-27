package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeDBInstanceNetworkRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UpgradeDBInstanceNetworkRequest) Invoke(client *sdk.Client) (resp *UpgradeDBInstanceNetworkResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceNetwork", "rds", "")
	resp = &UpgradeDBInstanceNetworkResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeDBInstanceNetworkResponse struct {
	responses.BaseResponse
	RequestId      string
	DBInstanceName string
}
