package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMyGroupsRequest struct {
	requests.RpcRequest
	SelectContactGroups string `position:"Query" name:"SelectContactGroups"`
	InstanceId          string `position:"Query" name:"InstanceId"`
	PageSize            int    `position:"Query" name:"PageSize"`
	Keyword             string `position:"Query" name:"Keyword"`
	Type                string `position:"Query" name:"Type"`
	GroupName           string `position:"Query" name:"GroupName"`
	PageNumber          int    `position:"Query" name:"PageNumber"`
	BindUrls            string `position:"Query" name:"BindUrls"`
}

func (req *ListMyGroupsRequest) Invoke(client *sdk.Client) (resp *ListMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListMyGroups", "cms", "")
	resp = &ListMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
	PageNumber   int
	PageSize     int
	Total        int
	Resources    ListMyGroupsResourceList
}

type ListMyGroupsResource struct {
	GroupId       int64
	GroupName     string
	ServiceId     string
	BindUrls      string
	Type          string
	GmtModified   int64
	GmtCreate     int64
	ContactGroups ListMyGroupsContactGroupList
}

type ListMyGroupsContactGroup struct {
	Name string
}

type ListMyGroupsResourceList []ListMyGroupsResource

func (list *ListMyGroupsResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupsResource)
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

type ListMyGroupsContactGroupList []ListMyGroupsContactGroup

func (list *ListMyGroupsContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupsContactGroup)
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
