package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListContactGroupRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	PageSize         int    `position:"Query" name:"PageSize"`
	PageNumber       int    `position:"Query" name:"PageNumber"`
}

func (req *ListContactGroupRequest) Invoke(client *sdk.Client) (resp *ListContactGroupResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListContactGroup", "cms", "")
	resp = &ListContactGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListContactGroupResponse struct {
	responses.BaseResponse
	Success       bool
	Code          string
	Message       string
	NextToken     int
	Total         int
	RequestId     string
	ContactGroups ListContactGroupContactGroupList
}

type ListContactGroupContactGroupList []string

func (list *ListContactGroupContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
