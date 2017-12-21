package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiTrafficDataRequest struct {
	ApiId     string `position:"Query" name:"ApiId"`
	GroupId   string `position:"Query" name:"GroupId"`
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

func (r DescribeApiTrafficDataRequest) Invoke(client *sdk.Client) (response *DescribeApiTrafficDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeApiTrafficDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiTrafficData", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeApiTrafficDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApiTrafficDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApiTrafficDataResponse struct {
	RequestId     string
	CallUploads   DescribeApiTrafficDataMonitorItemList
	CallDownloads DescribeApiTrafficDataMonitorItemList
}

type DescribeApiTrafficDataMonitorItem struct {
	ItemTime  string
	ItemValue string
}

type DescribeApiTrafficDataMonitorItemList []DescribeApiTrafficDataMonitorItem

func (list *DescribeApiTrafficDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiTrafficDataMonitorItem)
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
