package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListGroupsRequest struct {
	requests.RpcRequest
	Marker   string `position:"Query" name:"Marker"`
	MaxItems int    `position:"Query" name:"MaxItems"`
}

func (req *ListGroupsRequest) Invoke(client *sdk.Client) (resp *ListGroupsResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListGroups", "", "")
	resp = &ListGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListGroupsResponse struct {
	responses.BaseResponse
	RequestId   string
	IsTruncated bool
	Marker      string
	Groups      ListGroupsGroupList
}

type ListGroupsGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}

type ListGroupsGroupList []ListGroupsGroup

func (list *ListGroupsGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListGroupsGroup)
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
