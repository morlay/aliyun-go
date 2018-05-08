package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDrdsInstanceDescriptionRequest struct {
	requests.RpcRequest
	Description    string `position:"Query" name:"Description"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *ModifyDrdsInstanceDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyDrdsInstanceDescriptionResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsInstanceDescription", "", "")
	resp = &ModifyDrdsInstanceDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDrdsInstanceDescriptionResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
