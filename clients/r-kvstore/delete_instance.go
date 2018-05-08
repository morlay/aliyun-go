package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteInstanceRequest) Invoke(client *sdk.Client) (resp *DeleteInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteInstance", "redisa", "")
	resp = &DeleteInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
