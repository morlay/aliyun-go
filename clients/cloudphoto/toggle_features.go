package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ToggleFeaturesRequest struct {
	DisabledFeaturess *ToggleFeaturesDisabledFeaturesList `position:"Query" type:"Repeated" name:"DisabledFeatures"`
	StoreName         string                              `position:"Query" name:"StoreName"`
	EnabledFeaturess  *ToggleFeaturesEnabledFeaturesList  `position:"Query" type:"Repeated" name:"EnabledFeatures"`
}

func (r ToggleFeaturesRequest) Invoke(client *sdk.Client) (response *ToggleFeaturesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ToggleFeaturesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ToggleFeatures", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ToggleFeaturesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ToggleFeaturesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ToggleFeaturesResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
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
