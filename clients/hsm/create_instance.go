package hsm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateInstanceRequest struct {
	requests.RpcRequest
	Period          int    `position:"Query" name:"Period"`
	PeriodUnit      string `position:"Query" name:"PeriodUnit"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	Quantity        int    `position:"Query" name:"Quantity"`
	HsmDeviceType   string `position:"Query" name:"HsmDeviceType"`
	ClientToken     string `position:"Query" name:"ClientToken"`
	ZoneId          string `position:"Query" name:"ZoneId"`
	HsmOem          string `position:"Query" name:"HsmOem"`
}

func (req *CreateInstanceRequest) Invoke(client *sdk.Client) (resp *CreateInstanceResponse, err error) {
	req.InitWithApiInfo("hsm", "2018-01-11", "CreateInstance", "hsm", "")
	resp = &CreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateInstanceResponse struct {
	responses.BaseResponse
	RequestId   string
	InstanceIds CreateInstanceInstanceIdList
}

type CreateInstanceInstanceIdList []string

func (list *CreateInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
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
