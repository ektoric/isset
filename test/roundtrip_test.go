package test

import (
	"encoding/json"
	"github.com/ektoric/isset/pkg/isset"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientRequest struct {
	TInt    isset.Int    `json:"tint"`
	TFloat  isset.Float  `json:"tfloat"`
	TBool   isset.Bool   `json:"tbool"`
	TString isset.String `json:"tstring"`
}

type ServerResponse struct {
	TInt    *isset.Int    `json:"tint,omitempty"`
	TFloat  *isset.Float  `json:"tfloat,omitempty"`
	TBool   *isset.Bool   `json:"tbool,omitempty"`
	TString *isset.String `json:"tstring,omitempty"`
}

func rspFromReq(req ClientRequest) ServerResponse {
	rsp := ServerResponse{}
	if req.TInt.IsSet() {
		rsp.TInt = &req.TInt
	}
	if req.TFloat.IsSet() {
		rsp.TFloat = &req.TFloat
	}
	if req.TBool.IsSet() {
		rsp.TBool = &req.TBool
	}
	if req.TString.IsSet() {
		rsp.TString = &req.TString
	}
	return rsp
}

func (suite *JsonRoundTripTypeTester) TestJsonMarshal() {
	body := []byte(`
    {
        "tfloat": 0,
        "tbool": false,
        "tstring": null
    }
    `)
	value := ClientRequest{} // clear it
	err := json.Unmarshal(body, &value)
	suite.Require().NoError(err)
	suite.Assert().False(value.TInt.IsSet())
	suite.Assert().False(value.TInt.IsNull())
	suite.Assert().True(value.TFloat.IsSet())
	suite.Assert().False(value.TFloat.IsNull())
	suite.Assert().True(value.TBool.IsSet())
	suite.Assert().False(value.TBool.IsNull())
	suite.Assert().True(value.TString.IsSet())
	suite.Assert().True(value.TString.IsNull())

	// roundtrip it
	rsp := rspFromReq(value)
	body, err = json.Marshal(rsp)
	suite.Require().NoError(err)
	suite.Assert().Equal(body, []byte(`{"tfloat":0,"tbool":false,"tstring":null}`))
}

type JsonRoundTripTypeTester struct {
	suite.Suite
}

// Testing entry point.
func TestJsonRoundTrip(t *testing.T) {
	suite.Run(t, new(JsonRoundTripTypeTester))
}
