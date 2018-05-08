package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteDBInstanceRequest) Invoke(client *sdk.Client) (resp *DeleteDBInstanceResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DeleteDBInstance", "dds", "")
	resp = &DeleteDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDBInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
