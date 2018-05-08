package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ToggleFeaturesRequest struct {
	requests.RpcRequest
	DisabledFeaturess *ToggleFeaturesDisabledFeaturesList `position:"Query" type:"Repeated" name:"DisabledFeatures"`
	StoreName         string                              `position:"Query" name:"StoreName"`
	EnabledFeaturess  *ToggleFeaturesEnabledFeaturesList  `position:"Query" type:"Repeated" name:"EnabledFeatures"`
}

func (req *ToggleFeaturesRequest) Invoke(client *sdk.Client) (resp *ToggleFeaturesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ToggleFeatures", "cloudphoto", "")
	resp = &ToggleFeaturesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ToggleFeaturesResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
}

type ToggleFeaturesDisabledFeaturesList []string

func (list *ToggleFeaturesDisabledFeaturesList) UnmarshalJSON(data []byte) error {
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

type ToggleFeaturesEnabledFeaturesList []string

func (list *ToggleFeaturesEnabledFeaturesList) UnmarshalJSON(data []byte) error {
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
