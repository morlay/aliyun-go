package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMyGroupsRequest struct {
	requests.RpcRequest
	ContactGroups string `position:"Query" name:"ContactGroups"`
	ServiceId     int64  `position:"Query" name:"ServiceId"`
	Type          string `position:"Query" name:"Type"`
	GroupName     string `position:"Query" name:"GroupName"`
	BindUrl       string `position:"Query" name:"BindUrl"`
}

func (req *CreateMyGroupsRequest) Invoke(client *sdk.Client) (resp *CreateMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "CreateMyGroups", "cms", "")
	resp = &CreateMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
	GroupId      int64
}
