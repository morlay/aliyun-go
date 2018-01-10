package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMyGroupsRequest struct {
	requests.RpcRequest
	GroupId int64 `position:"Query" name:"GroupId"`
}

func (req *DeleteMyGroupsRequest) Invoke(client *sdk.Client) (resp *DeleteMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "DeleteMyGroups", "cms", "")
	resp = &DeleteMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Group        DeleteMyGroupsGroup
}

type DeleteMyGroupsGroup struct {
	GroupId       int64
	GroupName     string
	ServiceId     string
	BindUrls      string
	Type          string
	ContactGroups DeleteMyGroupsContactGroupList
}

type DeleteMyGroupsContactGroup struct {
	Name string
}

type DeleteMyGroupsContactGroupList []DeleteMyGroupsContactGroup

func (list *DeleteMyGroupsContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteMyGroupsContactGroup)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
