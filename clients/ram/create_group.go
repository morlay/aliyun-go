package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Group     CreateGroupGroup
}

type CreateGroupGroup struct {
	GroupName  common.String
	Comments   common.String
	CreateDate common.String
}
