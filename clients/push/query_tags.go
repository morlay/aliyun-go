package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTagsRequest struct {
	requests.RpcRequest
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
}

func (req *QueryTagsRequest) Invoke(client *sdk.Client) (resp *QueryTagsResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryTags", "", "")
	resp = &QueryTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTagsResponse struct {
	responses.BaseResponse
	RequestId common.String
	TagInfos  QueryTagsTagInfoList
}

type QueryTagsTagInfo struct {
	TagName common.String
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
