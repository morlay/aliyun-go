package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListRegionsRequest struct {
	requests.RpcRequest
}

func (req *ListRegionsRequest) Invoke(client *sdk.Client) (resp *ListRegionsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListRegions", "ehs", "")
	resp = &ListRegionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRegionsResponse struct {
	responses.BaseResponse
	RequestId string
	Regions   ListRegionsRegionInfoList
}

type ListRegionsRegionInfo struct {
	RegionId  string
	LocalName string
}

type ListRegionsRegionInfoList []ListRegionsRegionInfo

func (list *ListRegionsRegionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegionsRegionInfo)
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
