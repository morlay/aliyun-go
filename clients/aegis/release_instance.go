package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReleaseInstanceRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *ReleaseInstanceRequest) Invoke(client *sdk.Client) (resp *ReleaseInstanceResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "ReleaseInstance", "vipaegis", "")
	resp = &ReleaseInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
