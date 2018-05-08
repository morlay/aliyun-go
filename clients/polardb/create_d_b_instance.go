package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
}

func (req *CreateDBInstanceRequest) Invoke(client *sdk.Client) (resp *CreateDBInstanceResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateDBInstance", "polardb", "")
	resp = &CreateDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDBInstanceResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DBClusterId  common.String
	DBInstanceId common.String
	OrderId      common.String
}
