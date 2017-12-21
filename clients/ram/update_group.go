package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateGroupRequest struct {
	requests.RpcRequest
	NewGroupName string `position:"Query" name:"NewGroupName"`
	NewComments  string `position:"Query" name:"NewComments"`
	GroupName    string `position:"Query" name:"GroupName"`
}

func (req *UpdateGroupRequest) Invoke(client *sdk.Client) (resp *UpdateGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateGroup", "", "")
	resp = &UpdateGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateGroupResponse struct {
	responses.BaseResponse
	RequestId string
	Group     UpdateGroupGroup
}

type UpdateGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}
