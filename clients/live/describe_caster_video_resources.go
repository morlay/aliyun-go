package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterVideoResourcesRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCasterVideoResourcesRequest) Invoke(client *sdk.Client) (resp *DescribeCasterVideoResourcesResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterVideoResources", "live", "")
	resp = &DescribeCasterVideoResourcesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterVideoResourcesResponse struct {
	responses.BaseResponse
	RequestId      string
	Total          int
	VideoResources DescribeCasterVideoResourcesVideoResourceList
}

type DescribeCasterVideoResourcesVideoResource struct {
	MaterialId    string
	ResourceId    string
	ResourceName  string
	LocationId    string
	LiveStreamUrl string
	RepeatNum     int
	VodUrl        string
	BeginOffset   int
	EndOffset     int
}

type DescribeCasterVideoResourcesVideoResourceList []DescribeCasterVideoResourcesVideoResource

func (list *DescribeCasterVideoResourcesVideoResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterVideoResourcesVideoResource)
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
