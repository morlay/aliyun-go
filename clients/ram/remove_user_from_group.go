package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveUserFromGroupRequest struct {
	requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
	UserName  string `position:"Query" name:"UserName"`
}

func (req *RemoveUserFromGroupRequest) Invoke(client *sdk.Client) (resp *RemoveUserFromGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "RemoveUserFromGroup", "", "")
	resp = &RemoveUserFromGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveUserFromGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
