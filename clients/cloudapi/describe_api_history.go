package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiHistoryRequest struct {
	requests.RpcRequest
	GroupId        string `position:"Query" name:"GroupId"`
	ApiId          string `position:"Query" name:"ApiId"`
	StageName      string `position:"Query" name:"StageName"`
	HistoryVersion string `position:"Query" name:"HistoryVersion"`
}

func (req *DescribeApiHistoryRequest) Invoke(client *sdk.Client) (resp *DescribeApiHistoryResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiHistory", "apigateway", "")
	resp = &DescribeApiHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiHistoryResponse struct {
	responses.BaseResponse
	RequestId              common.String
	RegionId               common.String
	GroupId                common.String
	GroupName              common.String
	StageName              common.String
	ApiId                  common.String
	ApiName                common.String
	Description            common.String
	HistoryVersion         common.String
	Status                 common.String
	Visibility             common.String
	AuthType               common.String
	ResultType             common.String
	ResultSample           common.String
	FailResultSample       common.String
	DeployedTime           common.String
	AllowSignatureMethod   common.String
	ErrorCodeSamples       DescribeApiHistoryErrorCodeSampleList
	SystemParameters       DescribeApiHistorySystemParameterList
	CustomSystemParameters DescribeApiHistorySystemParameterList
	ConstantParameters     DescribeApiHistoryConstantParameterList
	RequestParameters      DescribeApiHistoryRequestParameterList
	ServiceParameters      DescribeApiHistoryServiceParameterList
	ServiceParametersMap   DescribeApiHistoryServiceParameterMapList
	RequestConfig          DescribeApiHistoryRequestConfig
	ServiceConfig          DescribeApiHistoryServiceConfig
	OpenIdConnectConfig    DescribeApiHistoryOpenIdConnectConfig
}

type DescribeApiHistoryErrorCodeSampleList []DescribeApiHistoryErrorCodeSample

func (list *DescribeApiHistoryErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiHistoryErrorCodeSample)
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

type DescribeApiHistorySystemParameterList []DescribeApiHistorySystemParameter

func (list *DescribeApiHistorySystemParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiHistorySystemParameter)
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

type DescribeApiHistoryConstantParameterList []DescribeApiHistoryConstantParameter

func (list *DescribeApiHistoryConstantParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiHistoryConstantParameter)
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

type DescribeApiHistoryRequestParameterList []DescribeApiHistoryRequestParameter

func (list *DescribeApiHistoryRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiHistoryRequestParameter)
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

type DescribeApiHistoryServiceParameterList []DescribeApiHistoryServiceParameter

func (list *DescribeApiHistoryServiceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiHistoryServiceParameter)
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

type DescribeApiHistoryServiceParameterMapList []DescribeApiHistoryServiceParameterMap

func (list *DescribeApiHistoryServiceParameterMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiHistoryServiceParameterMap)
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
