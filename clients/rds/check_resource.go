package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckResourceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SpecifyCount         string `position:"Query" name:"SpecifyCount"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	Engine               string `position:"Query" name:"Engine"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	DBInstanceUseType    string `position:"Query" name:"DBInstanceUseType"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
}

func (r CheckResourceRequest) Invoke(client *sdk.Client) (response *CheckResourceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckResourceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CheckResource", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CheckResourceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckResourceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckResourceResponse struct {
	RequestId    string
	SpecifyCount string
	Resources    CheckResourceResourceList
}

type CheckResourceResource struct {
	DBInstanceAvailable string
	Engine              string
	EngineVersion       string
}

type CheckResourceResourceList []CheckResourceResource

func (list *CheckResourceResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckResourceResource)
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
