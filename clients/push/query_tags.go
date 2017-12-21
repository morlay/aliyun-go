package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTagsRequest struct {
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
}

func (r QueryTagsRequest) Invoke(client *sdk.Client) (response *QueryTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "QueryTags", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryTagsResponse struct {
	RequestId string
	TagInfos  QueryTagsTagInfoList
}

type QueryTagsTagInfo struct {
	TagName string
}

type QueryTagsTagInfoList []QueryTagsTagInfo

func (list *QueryTagsTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagsTagInfo)
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
