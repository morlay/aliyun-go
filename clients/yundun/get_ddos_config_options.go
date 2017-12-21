package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetDdosConfigOptionsRequest struct {
}

func (r GetDdosConfigOptionsRequest) Invoke(client *sdk.Client) (response *GetDdosConfigOptionsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetDdosConfigOptionsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "GetDdosConfigOptions", "", "")

	resp := struct {
		*responses.BaseResponse
		GetDdosConfigOptionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetDdosConfigOptionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetDdosConfigOptionsResponse struct {
	RequestId                  string
	RequestThresholdOptions1   GetDdosConfigOptionsRequestThresholdOptionList
	RequestThresholdOptions2   GetDdosConfigOptionsRequestThresholdOptionList
	ConnectionThresholdOptions GetDdosConfigOptionsConnectionThresholdOptionList
	QpsOptions1                GetDdosConfigOptionsQpsOptions1List
	QpsOptions2                GetDdosConfigOptionsQpsOptions2List
}

type GetDdosConfigOptionsRequestThresholdOption struct {
	Bps int64
	Pps int64
}

type GetDdosConfigOptionsConnectionThresholdOption struct {
	Sipconn int64
	Sipnew  int64
}

type GetDdosConfigOptionsRequestThresholdOptionList []GetDdosConfigOptionsRequestThresholdOption

func (list *GetDdosConfigOptionsRequestThresholdOptionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDdosConfigOptionsRequestThresholdOption)
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

type GetDdosConfigOptionsConnectionThresholdOptionList []GetDdosConfigOptionsConnectionThresholdOption

func (list *GetDdosConfigOptionsConnectionThresholdOptionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDdosConfigOptionsConnectionThresholdOption)
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

type GetDdosConfigOptionsQpsOptions1List []string

func (list *GetDdosConfigOptionsQpsOptions1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type GetDdosConfigOptionsQpsOptions2List []string

func (list *GetDdosConfigOptionsQpsOptions2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
