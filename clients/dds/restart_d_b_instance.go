package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RestartDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
}

func (req *RestartDBInstanceRequest) Invoke(client *sdk.Client) (resp *RestartDBInstanceResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "RestartDBInstance", "dds", "")
	resp = &RestartDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RestartDBInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
