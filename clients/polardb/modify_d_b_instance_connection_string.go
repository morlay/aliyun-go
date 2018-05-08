package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBInstanceConnectionStringRequest struct {
	requests.RpcRequest
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (req *ModifyDBInstanceConnectionStringRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceConnectionStringResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBInstanceConnectionString", "polardb", "")
	resp = &ModifyDBInstanceConnectionStringResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceConnectionStringResponse struct {
	responses.BaseResponse
	RequestId           common.String
	DBInstanceId        common.String
	OldConnectionString common.String
	OldPort             common.String
	NewConnectionString common.String
	NewPort             common.String
	IPType              common.String
}
