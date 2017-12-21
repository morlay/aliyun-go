package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiErrorDataRequest struct {
	ApiId     string `position:"Query" name:"ApiId"`
	GroupId   string `position:"Query" name:"GroupId"`
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

func (r DescribeApiErrorDataRequest) Invoke(client *sdk.Client) (response *DescribeApiErrorDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeApiErrorDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiErrorData", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeApiErrorDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApiErrorDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApiErrorDataResponse struct {
	RequestId    string
	ClientErrors DescribeApiErrorDataMonitorItemList
	ServerErrors DescribeApiErrorDataMonitorItemList
}

type DescribeApiErrorDataMonitorItem struct {
	ItemTime  string
	ItemValue string
}

type DescribeApiErrorDataMonitorItemList []DescribeApiErrorDataMonitorItem

func (list *DescribeApiErrorDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorDataMonitorItem)
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
