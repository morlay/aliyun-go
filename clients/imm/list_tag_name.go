package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTagNameRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *ListTagNameRequest) Invoke(client *sdk.Client) (resp *ListTagNameResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListTagName", "imm", "")
	resp = &ListTagNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagNameResponse struct {
	responses.BaseResponse
	RequestId common.String
	Tags      ListTagNameTagsItemList
}

type ListTagNameTagsItem struct {
	TagName common.String
	Num     common.Integer
}

type ListTagNameTagsItemList []ListTagNameTagsItem

func (list *ListTagNameTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagNameTagsItem)
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
