package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterVideoResourcesRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r DescribeCasterVideoResourcesRequest) Invoke(client *sdk.Client) (response *DescribeCasterVideoResourcesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCasterVideoResourcesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterVideoResources", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCasterVideoResourcesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCasterVideoResourcesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCasterVideoResourcesResponse struct {
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
