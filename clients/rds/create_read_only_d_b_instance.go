package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateReadOnlyDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int    `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	VPCId                 string `position:"Query" name:"VPCId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
}

func (req *CreateReadOnlyDBInstanceRequest) Invoke(client *sdk.Client) (resp *CreateReadOnlyDBInstanceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateReadOnlyDBInstance", "rds", "")
	resp = &CreateReadOnlyDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateReadOnlyDBInstanceResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DBInstanceId     common.String
	OrderId          common.String
	ConnectionString common.String
	Port             common.String
}
