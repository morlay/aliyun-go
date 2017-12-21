package cloudauth

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitMaterialsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                        `position:"Query" name:"ResourceOwnerId"`
	Materials       *SubmitMaterialsMaterialList `position:"Query" type:"Repeated" name:"Material"`
	VerifyToken     string                       `position:"Query" name:"VerifyToken"`
}

func (req *SubmitMaterialsRequest) Invoke(client *sdk.Client) (resp *SubmitMaterialsResponse, err error) {
	req.InitWithApiInfo("Cloudauth", "2017-11-17", "SubmitMaterials", "cloudauth", "")
	resp = &SubmitMaterialsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitMaterialsMaterial struct {
	MaterialType string `name:"MaterialType"`
	Value        string `name:"Value"`
}

type SubmitMaterialsResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      SubmitMaterialsData
}

type SubmitMaterialsData struct {
	VerifyStatus SubmitMaterialsVerifyStatus
}

type SubmitMaterialsVerifyStatus struct {
	StatusCode      int
	TrustedScore    float32
	SimilarityScore float32
}

type SubmitMaterialsMaterialList []SubmitMaterialsMaterial

func (list *SubmitMaterialsMaterialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMaterialsMaterial)
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
