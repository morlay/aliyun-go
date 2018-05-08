package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PickLocalNumberRequest struct {
	requests.RpcRequest
	InstanceId       string                              `position:"Query" name:"InstanceId"`
	CandidateNumbers *PickLocalNumberCandidateNumberList `position:"Query" type:"Repeated" name:"CandidateNumber"`
	CalleeNumber     string                              `position:"Query" name:"CalleeNumber"`
}

func (req *PickLocalNumberRequest) Invoke(client *sdk.Client) (resp *PickLocalNumberResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "PickLocalNumber", "ccc", "")
	resp = &PickLocalNumberResponse{}
	err = client.DoAction(req, resp)
	return
}

type PickLocalNumberResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      PickLocalNumberData
}

type PickLocalNumberData struct {
	Callee PickLocalNumberCallee
	Caller PickLocalNumberCaller
}

type PickLocalNumberCallee struct {
	Number   common.String
	Province common.String
	City     common.String
}

type PickLocalNumberCaller struct {
	Number   common.String
	Province common.String
	City     common.String
}

type PickLocalNumberCandidateNumberList []string

func (list *PickLocalNumberCandidateNumberList) UnmarshalJSON(data []byte) error {
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
