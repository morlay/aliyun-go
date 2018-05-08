package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
	PageNumber   common.Integer
	PageSize     common.Integer
	Total        common.Integer
	Resources    ListMyGroupsResourceList
}

type ListMyGroupsResource struct {
	GroupId       common.Long
	GroupName     common.String
	ServiceId     common.String
	BindUrls      common.String
	Type          common.String
	GmtModified   common.Long
	GmtCreate     common.Long
	ContactGroups ListMyGroupsContactGroupList
}

type ListMyGroupsContactGroup struct {
	Name common.String
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
