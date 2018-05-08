package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTaskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryTaskRequest) Invoke(client *sdk.Client) (resp *QueryTaskResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "QueryTask", "redisa", "")
	resp = &QueryTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTaskResponse struct {
	responses.BaseResponse
	RequestId common.String
	Action    common.String
	Progress  common.Integer
}
