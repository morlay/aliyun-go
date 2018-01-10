package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMyGroupsRequest struct {
	requests.RpcRequest
	SelectContactGroups string `position:"Query" name:"SelectContactGroups"`
	InstanceId          string `position:"Query" name:"InstanceId"`
	GroupId             int64  `position:"Query" name:"GroupId"`
	Type                string `position:"Query" name:"Type"`
	GroupName           string `position:"Query" name:"GroupName"`
	BindUrl             string `position:"Query" name:"BindUrl"`
}

func (req *GetMyGroupsRequest) Invoke(client *sdk.Client) (resp *GetMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "GetMyGroups", "cms", "")
	resp = &GetMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Group        GetMyGroupsGroup
}

type GetMyGroupsGroup struct {
	GroupId       int64
	GroupName     string
	ServiceId     int64
	BindUrl       string
	Type          string
	ContactGroups GetMyGroupsContactGroupList
}

type GetMyGroupsContactGroup struct {
	Name string
}

type GetMyGroupsContactGroupList []GetMyGroupsContactGroup

func (list *GetMyGroupsContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetMyGroupsContactGroup)
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
