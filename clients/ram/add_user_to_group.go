package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddUserToGroupRequest struct {
	requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
	UserName  string `position:"Query" name:"UserName"`
}

func (req *AddUserToGroupRequest) Invoke(client *sdk.Client) (resp *AddUserToGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "AddUserToGroup", "", "")
	resp = &AddUserToGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddUserToGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
