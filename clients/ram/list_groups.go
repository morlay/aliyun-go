package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	IsTruncated bool
	Marker      common.String
	Groups      ListGroupsGroupList
}

type ListGroupsGroup struct {
	GroupName  common.String
	Comments   common.String
	CreateDate common.String
	UpdateDate common.String
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
