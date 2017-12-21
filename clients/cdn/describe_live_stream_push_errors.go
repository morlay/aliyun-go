package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamPushErrorsRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamPushErrorsRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamPushErrorsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamPushErrors", "", "")
	resp = &DescribeLiveStreamPushErrorsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamPushErrorsResponse struct {
	responses.BaseResponse
	RequestId           string
	PushErrorsModelList DescribeLiveStreamPushErrorsPushErrorsModelList
}

type DescribeLiveStreamPushErrorsPushErrorsModel struct {
	ErrorCode string
}

type DescribeLiveStreamPushErrorsPushErrorsModelList []DescribeLiveStreamPushErrorsPushErrorsModel

func (list *DescribeLiveStreamPushErrorsPushErrorsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamPushErrorsPushErrorsModel)
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
