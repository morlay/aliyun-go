package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListRegisteredTagsRequest struct {
	StoreName string                      `position:"Query" name:"StoreName"`
	Langs     *ListRegisteredTagsLangList `position:"Query" type:"Repeated" name:"Lang"`
}

func (r ListRegisteredTagsRequest) Invoke(client *sdk.Client) (response *ListRegisteredTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListRegisteredTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListRegisteredTags", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ListRegisteredTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListRegisteredTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListRegisteredTagsResponse struct {
	Code           string
	Message        string
	RequestId      string
	Action         string
	RegisteredTags ListRegisteredTagsRegisteredTagList
}

type ListRegisteredTagsRegisteredTag struct {
	TagKey    string
	TagValues ListRegisteredTagsTagValueList
}

type ListRegisteredTagsTagValue struct {
	Lang string
	Text string
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
