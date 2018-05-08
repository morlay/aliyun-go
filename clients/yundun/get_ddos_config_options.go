package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetDdosConfigOptionsRequest struct {
	requests.RpcRequest
}

func (req *GetDdosConfigOptionsRequest) Invoke(client *sdk.Client) (resp *GetDdosConfigOptionsResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "GetDdosConfigOptions", "", "")
	resp = &GetDdosConfigOptionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetDdosConfigOptionsResponse struct {
	responses.BaseResponse
	RequestId                  common.String
	RequestThresholdOptions1   GetDdosConfigOptionsRequestThresholdOptionList
	RequestThresholdOptions2   GetDdosConfigOptionsRequestThresholdOptionList
	ConnectionThresholdOptions GetDdosConfigOptionsConnectionThresholdOptionList
	QpsOptions1                GetDdosConfigOptionsQpsOptions1List
	QpsOptions2                GetDdosConfigOptionsQpsOptions2List
}

type GetDdosConfigOptionsRequestThresholdOption struct {
	Bps common.Long
	Pps common.Long
}

type GetDdosConfigOptionsConnectionThresholdOption struct {
	Sipconn common.Long
	Sipnew  common.Long
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

type GetDdosConfigOptionsQpsOptions1List []common.String

func (list *GetDdosConfigOptionsQpsOptions1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type GetDdosConfigOptionsQpsOptions2List []common.String

func (list *GetDdosConfigOptionsQpsOptions2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
