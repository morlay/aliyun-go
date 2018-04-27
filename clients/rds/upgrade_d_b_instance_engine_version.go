package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeDBInstanceEngineVersionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UpgradeDBInstanceEngineVersionRequest) Invoke(client *sdk.Client) (resp *UpgradeDBInstanceEngineVersionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceEngineVersion", "rds", "")
	resp = &UpgradeDBInstanceEngineVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeDBInstanceEngineVersionResponse struct {
	responses.BaseResponse
	RequestId string
	TaskId    string
}
