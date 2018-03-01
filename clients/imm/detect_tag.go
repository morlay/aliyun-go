package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetectTagRequest struct {
	requests.RpcRequest
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
}

func (req *DetectTagRequest) Invoke(client *sdk.Client) (resp *DetectTagResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DetectTag", "imm", "")
	resp = &DetectTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetectTagResponse struct {
	responses.BaseResponse
	RequestId      string
	SuccessNum     string
	SuccessDetails DetectTagSuccessDetailsItemList
	FailDetails    DetectTagFailDetailsItemList
}

type DetectTagSuccessDetailsItem struct {
	SrcUri string
	Tags   DetectTagTagsItemList
}

type DetectTagTagsItem struct {
	TagId         string
	TagLevel      string
	TagName       string
	ParentTagId   string
	ParentTagName string
	TagScore      string
}

type DetectTagFailDetailsItem struct {
	SrcUri string
	Reason string
}

type DetectTagSuccessDetailsItemList []DetectTagSuccessDetailsItem

func (list *DetectTagSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectTagSuccessDetailsItem)
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

type DetectTagFailDetailsItemList []DetectTagFailDetailsItem

func (list *DetectTagFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectTagFailDetailsItem)
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

type DetectTagTagsItemList []DetectTagTagsItem

func (list *DetectTagTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectTagTagsItem)
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
