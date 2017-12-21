package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceMaintainTimeRequest struct {
	requests.RpcRequest
	MaintainTime         string `position:"Query" name:"MaintainTime"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBInstanceMaintainTimeRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceMaintainTimeResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBInstanceMaintainTime", "polardb", "")
	resp = &ModifyDBInstanceMaintainTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceMaintainTimeResponse struct {
	responses.BaseResponse
	RequestId string
}
