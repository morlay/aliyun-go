package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMyGroupsRequest struct {
	requests.RpcRequest
	ContactGroups string `position:"Query" name:"ContactGroups"`
	GroupId       string `position:"Query" name:"GroupId"`
	ServiceId     int64  `position:"Query" name:"ServiceId"`
	Type          string `position:"Query" name:"Type"`
	GroupName     string `position:"Query" name:"GroupName"`
	BindUrls      string `position:"Query" name:"BindUrls"`
}

func (req *UpdateMyGroupsRequest) Invoke(client *sdk.Client) (resp *UpdateMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "UpdateMyGroups", "cms", "")
	resp = &UpdateMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
}
