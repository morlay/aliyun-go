package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteGroupRequest struct {
	requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

func (req *DeleteGroupRequest) Invoke(client *sdk.Client) (resp *DeleteGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteGroup", "", "")
	resp = &DeleteGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
