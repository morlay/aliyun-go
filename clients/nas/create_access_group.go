package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateAccessGroupRequest struct {
	requests.RpcRequest
	Description     string `position:"Query" name:"Description"`
	AccessGroupType string `position:"Query" name:"AccessGroupType"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (req *CreateAccessGroupRequest) Invoke(client *sdk.Client) (resp *CreateAccessGroupResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateAccessGroup", "nas", "")
	resp = &CreateAccessGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAccessGroupResponse struct {
	responses.BaseResponse
	RequestId       common.String
	AccessGroupName common.String
}
