package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeDBInstanceEngineVersionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UpgradeDBInstanceEngineVersionRequest) Invoke(client *sdk.Client) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpgradeDBInstanceEngineVersionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceEngineVersion", "rds", "")

	resp := struct {
		*responses.BaseResponse
		UpgradeDBInstanceEngineVersionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpgradeDBInstanceEngineVersionResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpgradeDBInstanceEngineVersionResponse struct {
	RequestId string
	TaskId    string
}
