package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeDBInstanceNetWorkInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UpgradeDBInstanceNetWorkInfoRequest) Invoke(client *sdk.Client) (resp *UpgradeDBInstanceNetWorkInfoResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceNetWorkInfo", "rds", "")
	resp = &UpgradeDBInstanceNetWorkInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeDBInstanceNetWorkInfoResponse struct {
	responses.BaseResponse
	RequestId string
}
