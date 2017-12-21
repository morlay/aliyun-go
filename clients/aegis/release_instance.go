package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseInstanceRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *ReleaseInstanceRequest) Invoke(client *sdk.Client) (resp *ReleaseInstanceResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "ReleaseInstance", "", "")
	resp = &ReleaseInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
