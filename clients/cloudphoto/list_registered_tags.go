package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListRegisteredTagsRequest struct {
	requests.RpcRequest
	StoreName string                      `position:"Query" name:"StoreName"`
	Langs     *ListRegisteredTagsLangList `position:"Query" type:"Repeated" name:"Lang"`
}

func (req *ListRegisteredTagsRequest) Invoke(client *sdk.Client) (resp *ListRegisteredTagsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListRegisteredTags", "cloudphoto", "")
	resp = &ListRegisteredTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRegisteredTagsResponse struct {
	responses.BaseResponse
	Code           common.String
	Message        common.String
	RequestId      common.String
	Action         common.String
	RegisteredTags ListRegisteredTagsRegisteredTagList
}

type ListRegisteredTagsRegisteredTag struct {
	TagKey    common.String
	TagValues ListRegisteredTagsTagValueList
}

type ListRegisteredTagsTagValue struct {
	Lang common.String
	Text common.String
}

type ListRegisteredTagsLangList []string

func (list *ListRegisteredTagsLangList) UnmarshalJSON(data []byte) error {
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

type ListRegisteredTagsRegisteredTagList []ListRegisteredTagsRegisteredTag

func (list *ListRegisteredTagsRegisteredTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegisteredTagsRegisteredTag)
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

type ListRegisteredTagsTagValueList []ListRegisteredTagsTagValue

func (list *ListRegisteredTagsTagValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegisteredTagsTagValue)
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
