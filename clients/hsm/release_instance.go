package hsm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
}

func (req *ReleaseInstanceRequest) Invoke(client *sdk.Client) (resp *ReleaseInstanceResponse, err error) {
	req.InitWithApiInfo("hsm", "2018-01-11", "ReleaseInstance", "hsm", "")
	resp = &ReleaseInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
