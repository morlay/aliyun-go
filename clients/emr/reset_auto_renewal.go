package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResetAutoRenewalRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                    `position:"Query" name:"ResourceOwnerId"`
	ClusterId            string                                   `position:"Query" name:"ClusterId"`
	EcsResetAutoRenewDos *ResetAutoRenewalEcsResetAutoRenewDoList `position:"Query" type:"Repeated" name:"EcsResetAutoRenewDo"`
}

func (req *ResetAutoRenewalRequest) Invoke(client *sdk.Client) (resp *ResetAutoRenewalResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResetAutoRenewal", "", "")
	resp = &ResetAutoRenewalResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetAutoRenewalEcsResetAutoRenewDo struct {
	EcsId                string `name:"EcsId"`
	EcsAutoRenew         string `name:"EcsAutoRenew"`
	EcsAutoRenewDuration string `name:"EcsAutoRenewDuration"`
	EmrAutoRenew         string `name:"EmrAutoRenew"`
	EmrAutoRenewDuration string `name:"EmrAutoRenewDuration"`
}

type ResetAutoRenewalResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type ResetAutoRenewalEcsResetAutoRenewDoList []ResetAutoRenewalEcsResetAutoRenewDo

func (list *ResetAutoRenewalEcsResetAutoRenewDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResetAutoRenewalEcsResetAutoRenewDo)
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
