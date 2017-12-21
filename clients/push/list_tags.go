package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListTagsRequest struct {
	AppKey int64 `position:"Query" name:"AppKey"`
}

func (r ListTagsRequest) Invoke(client *sdk.Client) (response *ListTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "ListTags", "", "")

	resp := struct {
		*responses.BaseResponse
		ListTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListTagsResponse struct {
	RequestId string
	TagInfos  ListTagsTagInfoList
}

type ListTagsTagInfo struct {
	TagName string
}

type ListTagsTagInfoList []ListTagsTagInfo

func (list *ListTagsTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagsTagInfo)
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
