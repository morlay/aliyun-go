package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiHistoryRequest struct {
	GroupId        string `position:"Query" name:"GroupId"`
	ApiId          string `position:"Query" name:"ApiId"`
	StageName      string `position:"Query" name:"StageName"`
	HistoryVersion string `position:"Query" name:"HistoryVersion"`
}

func (r DescribeApiHistoryRequest) Invoke(client *sdk.Client) (response *DescribeApiHistoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeApiHistoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiHistory", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeApiHistoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApiHistoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApiHistoryResponse struct {
	RequestId              string
	RegionId               string
	GroupId                string
	GroupName              string
	StageName              string
	ApiId                  string
	ApiName                string
	Description            string
	HistoryVersion         string
	Status                 string
	Visibility             string
	AuthType               string
	ResultType             string
	ResultSample           string
	FailResultSample       string
	DeployedTime           string
	ErrorCodeSamples       DescribeApiHistoryErrorCodeSampleList
	SystemParameters       DescribeApiHistorySystemParameterList
	CustomSystemParameters DescribeApiHistorySystemParameterList
	ConstantParameters     DescribeApiHistoryConstantParameterList
	RequestParameters      DescribeApiHistoryRequestParameterList
	ServiceParameters      DescribeApiHistoryServiceParameterList
	ServiceParametersMap   DescribeApiHistoryServiceParameterMapList
	RequestConfig          DescribeApiRequestConfig
	ServiceConfig          DescribeApiServiceConfig
	OpenIdConnectConfig    DescribeApiOpenIdConnectConfig
}

type DescribeApiHistoryErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeApiHistoryErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorCodeSample)
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

type DescribeApiHistorySystemParameterList []DescribeApiSystemParameter

func (list *DescribeApiHistorySystemParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiSystemParameter)
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

type DescribeApiHistoryConstantParameterList []DescribeApiConstantParameter

func (list *DescribeApiHistoryConstantParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiConstantParameter)
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

type DescribeApiHistoryRequestParameterList []DescribeApiRequestParameter

func (list *DescribeApiHistoryRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiRequestParameter)
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

type DescribeApiHistoryServiceParameterList []DescribeApiServiceParameter

func (list *DescribeApiHistoryServiceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameter)
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

type DescribeApiHistoryServiceParameterMapList []DescribeApiServiceParameterMap

func (list *DescribeApiHistoryServiceParameterMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameterMap)
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
