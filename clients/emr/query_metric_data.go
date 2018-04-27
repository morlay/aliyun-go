package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMetricDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Period          int64  `position:"Query" name:"Period"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	StartTimeStamp  int64  `position:"Query" name:"StartTimeStamp"`
	MetricName      string `position:"Query" name:"MetricName"`
	HostRole        string `position:"Query" name:"HostRole"`
	EndTimeStamp    int64  `position:"Query" name:"EndTimeStamp"`
}

func (req *QueryMetricDataRequest) Invoke(client *sdk.Client) (resp *QueryMetricDataResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryMetricData", "", "")
	resp = &QueryMetricDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMetricDataResponse struct {
	responses.BaseResponse
	RequestId  string
	Datapoints QueryMetricDataCmsDataPointList
}

type QueryMetricDataCmsDataPoint struct {
	Role      string
	Maximum   float32
	Minimum   float32
	Average   float32
	Timestamp int64
}

type QueryMetricDataCmsDataPointList []QueryMetricDataCmsDataPoint

func (list *QueryMetricDataCmsDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMetricDataCmsDataPoint)
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
