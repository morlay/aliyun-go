package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateIpControlRequest struct {
	requests.RpcRequest
	IpControlName     string                               `position:"Query" name:"IpControlName"`
	IpControlType     string                               `position:"Query" name:"IpControlType"`
	IpControlPolicyss *CreateIpControlIpControlPolicysList `position:"Query" type:"Repeated" name:"IpControlPolicys"`
	Description       string                               `position:"Query" name:"Description"`
}

func (req *CreateIpControlRequest) Invoke(client *sdk.Client) (resp *CreateIpControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateIpControl", "apigateway", "")
	resp = &CreateIpControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateIpControlIpControlPolicys struct {
	AppId  string `name:"AppId"`
	CidrIp string `name:"CidrIp"`
}

type CreateIpControlResponse struct {
	responses.BaseResponse
	RequestId   common.String
	IpControlId common.String
}

type CreateIpControlIpControlPolicysList []CreateIpControlIpControlPolicys

func (list *CreateIpControlIpControlPolicysList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateIpControlIpControlPolicys)
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
