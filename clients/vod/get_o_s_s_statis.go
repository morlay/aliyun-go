package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOSSStatisRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	StartStatisTime      string `position:"Query" name:"StartStatisTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Level                string `position:"Query" name:"Level"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	EndStatisTime        string `position:"Query" name:"EndStatisTime"`
}

func (req *GetOSSStatisRequest) Invoke(client *sdk.Client) (resp *GetOSSStatisResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetOSSStatis", "vod", "")
	resp = &GetOSSStatisResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOSSStatisResponse struct {
	responses.BaseResponse
	RequestId             string
	MaxStorageUtilization int64
	OssStatisList         GetOSSStatisOSSMetricList
}

type GetOSSStatisOSSMetric struct {
	StatTime           string
	StorageUtilization int64
}

type GetOSSStatisOSSMetricList []GetOSSStatisOSSMetric

func (list *GetOSSStatisOSSMetricList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetOSSStatisOSSMetric)
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
