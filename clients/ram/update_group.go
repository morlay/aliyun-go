package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Group     UpdateGroupGroup
}

type UpdateGroupGroup struct {
	GroupName  common.String
	Comments   common.String
	CreateDate common.String
	UpdateDate common.String
}
