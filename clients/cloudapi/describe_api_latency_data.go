package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiLatencyDataRequest struct {
	requests.RpcRequest
	ApiId     string `position:"Query" name:"ApiId"`
	GroupId   string `position:"Query" name:"GroupId"`
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

func (req *DescribeApiLatencyDataRequest) Invoke(client *sdk.Client) (resp *DescribeApiLatencyDataResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiLatencyData", "apigateway", "")
	resp = &DescribeApiLatencyDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiLatencyDataResponse struct {
	responses.BaseResponse
	RequestId    string
	CallLatencys DescribeApiLatencyDataMonitorItemList
}

type DescribeApiLatencyDataMonitorItem struct {
	ItemTime  string
	ItemValue string
}

type DescribeApiLatencyDataMonitorItemList []DescribeApiLatencyDataMonitorItem

func (list *DescribeApiLatencyDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiLatencyDataMonitorItem)
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
