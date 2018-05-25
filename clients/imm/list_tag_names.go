package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTagNamesRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *ListTagNamesRequest) Invoke(client *sdk.Client) (resp *ListTagNamesResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListTagNames", "imm", "")
	resp = &ListTagNamesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagNamesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Tags      ListTagNamesTagsItemList
}

type ListTagNamesTagsItem struct {
	TagName common.String
	Num     common.Integer
}

type ListTagNamesTagsItemList []ListTagNamesTagsItem

func (list *ListTagNamesTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagNamesTagsItem)
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
