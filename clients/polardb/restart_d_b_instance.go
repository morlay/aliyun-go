package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RestartDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RestartDBInstanceRequest) Invoke(client *sdk.Client) (resp *RestartDBInstanceResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "RestartDBInstance", "polardb", "")
	resp = &RestartDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RestartDBInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
