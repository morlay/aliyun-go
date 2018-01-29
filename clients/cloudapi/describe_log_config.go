package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLogConfigRequest struct {
	requests.RpcRequest
	LogType string `position:"Query" name:"LogType"`
}

func (req *DescribeLogConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLogConfigResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeLogConfig", "apigateway", "")
	resp = &DescribeLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLogConfigResponse struct {
	responses.BaseResponse
	RequestId string
	LogInfos  DescribeLogConfigLogInfoList
}

type DescribeLogConfigLogInfo struct {
	RegionId    string
	SlsProject  string
	SlsLogStore string
	LogType     string
}

type DescribeLogConfigLogInfoList []DescribeLogConfigLogInfo

func (list *DescribeLogConfigLogInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLogConfigLogInfo)
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
