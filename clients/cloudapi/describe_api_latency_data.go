package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiLatencyDataRequest struct {
	ApiId     string `position:"Query" name:"ApiId"`
	GroupId   string `position:"Query" name:"GroupId"`
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

func (r DescribeApiLatencyDataRequest) Invoke(client *sdk.Client) (response *DescribeApiLatencyDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeApiLatencyDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiLatencyData", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeApiLatencyDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApiLatencyDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApiLatencyDataResponse struct {
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
