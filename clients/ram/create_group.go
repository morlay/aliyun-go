package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateGroupRequest struct {
	requests.RpcRequest
	Comments  string `position:"Query" name:"Comments"`
	GroupName string `position:"Query" name:"GroupName"`
}

func (req *CreateGroupRequest) Invoke(client *sdk.Client) (resp *CreateGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateGroup", "", "")
	resp = &CreateGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateGroupResponse struct {
	responses.BaseResponse
	RequestId string
	Group     CreateGroupGroup
}

type CreateGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
}
