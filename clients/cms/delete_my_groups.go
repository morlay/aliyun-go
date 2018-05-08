package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteMyGroupsRequest struct {
	requests.RpcRequest
	GroupId int64 `position:"Query" name:"GroupId"`
}

func (req *DeleteMyGroupsRequest) Invoke(client *sdk.Client) (resp *DeleteMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DeleteMyGroups", "cms", "")
	resp = &DeleteMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
	Group        DeleteMyGroupsGroup
}

type DeleteMyGroupsGroup struct {
	GroupId       common.Long
	GroupName     common.String
	ServiceId     common.String
	BindUrls      common.String
	Type          common.String
	ContactGroups DeleteMyGroupsContactGroupList
}

type DeleteMyGroupsContactGroup struct {
	Name common.String
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
