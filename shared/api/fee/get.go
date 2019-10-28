package fee

import (
    "errors"
    "math/big"

    "github.com/rocket-pool/smartnode/shared/services"
    "github.com/rocket-pool/smartnode/shared/utils/eth"
)


// Fee get response type
type FeeGetResponse struct {
    CurrentUserFeePerc float64  `json:"currentUserFeePerc"`
    TargetUserFeePerc float64   `json:"targetUserFeePerc"`
}


// Get user fee
func GetUserFee(p *services.Provider) (*FeeGetResponse, error) {

    // Response
    response := &FeeGetResponse{}

    // Open database
    if err := p.DB.Open(); err != nil {
        return nil, err
    }
    defer p.DB.Close()

    // Get current user fee
    userFee := new(*big.Int)
    if err := p.CM.Contracts["rocketNodeSettings"].Call(nil, userFee, "getFeePerc"); err != nil {
        return nil, errors.New("Error retrieving node user fee percentage setting: " + err.Error())
    } else {
        response.CurrentUserFeePerc = eth.WeiToEth(*userFee) * 100
    }

    // Get target user fee
    targetUserFeePerc := new(float64)
    *targetUserFeePerc = -1
    if err := p.DB.Get("user.fee.target", targetUserFeePerc); err != nil {
        return nil, errors.New("Error retrieving target node user fee percentage: " + err.Error())
    } else {
        response.TargetUserFeePerc = *targetUserFeePerc
    }

    // Return response
    return response, nil

}

