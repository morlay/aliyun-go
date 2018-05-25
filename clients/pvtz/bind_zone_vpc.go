package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BindZoneVpcRequest struct {
	requests.RpcRequest
	UserClientIp string               `position:"Query" name:"UserClientIp"`
	ZoneId       string               `position:"Query" name:"ZoneId"`
	Lang         string               `position:"Query" name:"Lang"`
	Vpcss        *BindZoneVpcVpcsList `position:"Query" type:"Repeated" name:"Vpcs"`
}

func (req *BindZoneVpcRequest) Invoke(client *sdk.Client) (resp *BindZoneVpcResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "BindZoneVpc", "pvtz", "")
	resp = &BindZoneVpcResponse{}
	err = client.DoAction(req, resp)
	return
}

type BindZoneVpcVpcs struct {
	RegionId string `name:"RegionId"`
	VpcId    string `name:"VpcId"`
}

type BindZoneVpcResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type BindZoneVpcVpcsList []BindZoneVpcVpcs

func (list *BindZoneVpcVpcsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BindZoneVpcVpcs)
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
