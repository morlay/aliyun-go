package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type IndexTagRequest struct {
	requests.RpcRequest
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	Force   string `position:"Query" name:"Force"`
}

func (req *IndexTagRequest) Invoke(client *sdk.Client) (resp *IndexTagResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "IndexTag", "imm", "")
	resp = &IndexTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type IndexTagResponse struct {
	responses.BaseResponse
	RequestId       string
	SetId           string
	SuccessIndexNum string
	FailDetails     IndexTagFailDetailsItemList
	SuccessDetails  IndexTagSuccessDetailsItemList
}

type IndexTagFailDetailsItem struct {
	SrcUri string
	Reason string
}

type IndexTagSuccessDetailsItem struct {
	SrcUri string
	Tags   IndexTagTagsItemList
}

type IndexTagTagsItem struct {
	TagId         string
	TagLevel      string
	TagName       string
	ParentTagId   string
	ParentTagName string
	TagScore      string
}

type IndexTagFailDetailsItemList []IndexTagFailDetailsItem

func (list *IndexTagFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexTagFailDetailsItem)
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

type IndexTagSuccessDetailsItemList []IndexTagSuccessDetailsItem

func (list *IndexTagSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexTagSuccessDetailsItem)
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

type IndexTagTagsItemList []IndexTagTagsItem

func (list *IndexTagTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexTagTagsItem)
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
