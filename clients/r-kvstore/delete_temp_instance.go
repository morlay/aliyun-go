package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteTempInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TempInstanceId       string `position:"Domain" name:"TempInstanceId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteTempInstanceRequest) Invoke(client *sdk.Client) (resp *DeleteTempInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteTempInstance", "redisa", "")
	resp = &DeleteTempInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTempInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
