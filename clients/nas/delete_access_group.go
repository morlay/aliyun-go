package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAccessGroupRequest struct {
	requests.RpcRequest
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (req *DeleteAccessGroupRequest) Invoke(client *sdk.Client) (resp *DeleteAccessGroupResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteAccessGroup", "nas", "")
	resp = &DeleteAccessGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAccessGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
