package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAccessGroupRequest struct {
	requests.RpcRequest
	Description     string `position:"Query" name:"Description"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (req *ModifyAccessGroupRequest) Invoke(client *sdk.Client) (resp *ModifyAccessGroupResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "ModifyAccessGroup", "nas", "")
	resp = &ModifyAccessGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAccessGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
