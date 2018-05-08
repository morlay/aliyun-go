package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type IndexFaceRequest struct {
	requests.RpcRequest
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	Force   string `position:"Query" name:"Force"`
}

func (req *IndexFaceRequest) Invoke(client *sdk.Client) (resp *IndexFaceResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "IndexFace", "imm", "")
	resp = &IndexFaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type IndexFaceResponse struct {
	responses.BaseResponse
	RequestId      common.String
	SetId          common.String
	SuccessDetails IndexFaceSuccessDetailsItemList
	FailDetails    IndexFaceFailDetailsItemList
	SrcUris        IndexFaceSrcUriList
}

type IndexFaceSuccessDetailsItem struct {
	SrcUri  common.String
	PhotoId common.String
	Faces   IndexFaceFacesItemList
}

type IndexFaceFacesItem struct {
	FaceId        common.String
	FaceRectangle IndexFaceFaceRectangle
	FaceAttribute IndexFaceFaceAttribute
}

type IndexFaceFaceRectangle struct {
	Top    common.String
	Left   common.String
	Width  common.String
	Height common.String
}

type IndexFaceFaceAttribute struct {
	Gender      IndexFaceGender
	Age         IndexFaceAge
	Headpose    IndexFaceHeadpose
	Eyestatus   IndexFaceEyestatus
	Blur        IndexFaceBlur
	Facequality IndexFaceFacequality
}

type IndexFaceGender struct {
	Value common.String
}

type IndexFaceAge struct {
	Value common.Integer
}

type IndexFaceHeadpose struct {
	Pitch_angle common.Float
	Roll_angle  common.Float
	Yaw_angle   common.Float
}

type IndexFaceEyestatus struct {
	Left_eye_status  IndexFaceLeft_eye_status
	Right_eye_status IndexFaceRight_eye_status
}

type IndexFaceLeft_eye_status struct {
	Normal_glass_eye_open  common.Float
	No_glass_eye_close     common.Float
	Occlusion              common.Float
	No_glass_eye_open      common.Float
	Normal_glass_eye_close common.Float
	Dark_glasses           common.Float
}

type IndexFaceRight_eye_status struct {
	Normal_glass_eye_open  common.Float
	No_glass_eye_close     common.Float
	Occlusion              common.Float
	No_glass_eye_open      common.Float
	Normal_glass_eye_close common.Float
	Dark_glasses           common.Float
}

type IndexFaceBlur struct {
	Blurness IndexFaceBlurness
}

type IndexFaceBlurness struct {
	Value     common.Float
	Threshold common.Float
}

type IndexFaceFacequality struct {
	Value     common.Float
	Threshold common.Float
}

type IndexFaceFailDetailsItem struct {
	SrcUri common.String
	Reason common.String
}

type IndexFaceSuccessDetailsItemList []IndexFaceSuccessDetailsItem

func (list *IndexFaceSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceSuccessDetailsItem)
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

type IndexFaceFailDetailsItemList []IndexFaceFailDetailsItem

func (list *IndexFaceFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceFailDetailsItem)
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

type IndexFaceSrcUriList []common.String

func (list *IndexFaceSrcUriList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type IndexFaceFacesItemList []IndexFaceFacesItem

func (list *IndexFaceFacesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceFacesItem)
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
