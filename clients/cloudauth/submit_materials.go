package cloudauth

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      SubmitMaterialsData
}

type SubmitMaterialsData struct {
	VerifyStatus SubmitMaterialsVerifyStatus
}

type SubmitMaterialsVerifyStatus struct {
	StatusCode      common.Integer
	TrustedScore    common.Float
	SimilarityScore common.Float
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
