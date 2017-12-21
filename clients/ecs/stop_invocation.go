package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopInvocationRequest struct {
	ResourceOwnerId      int64                         `position:"Query" name:"ResourceOwnerId"`
	InvokeId             string                        `position:"Query" name:"InvokeId"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                        `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                         `position:"Query" name:"OwnerId"`
	InstanceIds          *StopInvocationInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
}

func (r StopInvocationRequest) Invoke(client *sdk.Client) (response *StopInvocationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopInvocationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "StopInvocation", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		StopInvocationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopInvocationResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopInvocationResponse struct {
	RequestId string
}

type StopInvocationInstanceIdList []string

func (list *StopInvocationInstanceIdList) UnmarshalJSON(data []byte) error {
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
