package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	SuccessNum     common.String
	SuccessDetails DetectTagSuccessDetailsItemList
	FailDetails    DetectTagFailDetailsItemList
}

type DetectTagSuccessDetailsItem struct {
	SrcUri common.String
	Tags   DetectTagTagsItemList
}

type DetectTagTagsItem struct {
	TagId         common.String
	TagLevel      common.String
	TagName       common.String
	ParentTagId   common.String
	ParentTagName common.String
	TagScore      common.String
}

type DetectTagFailDetailsItem struct {
	SrcUri common.String
	Reason common.String
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
