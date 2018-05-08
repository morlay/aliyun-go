package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTagsRequest struct {
	requests.RpcRequest
	AppKey int64 `position:"Query" name:"AppKey"`
}

func (req *ListTagsRequest) Invoke(client *sdk.Client) (resp *ListTagsResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "ListTags", "", "")
	resp = &ListTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagsResponse struct {
	responses.BaseResponse
	RequestId common.String
	TagInfos  ListTagsTagInfoList
}

type ListTagsTagInfo struct {
	TagName common.String
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
