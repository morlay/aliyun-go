package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveDrdsInstanceRequest struct {
	requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *RemoveDrdsInstanceRequest) Invoke(client *sdk.Client) (resp *RemoveDrdsInstanceResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "RemoveDrdsInstance", "", "")
	resp = &RemoveDrdsInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveDrdsInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
