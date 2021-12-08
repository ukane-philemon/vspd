// Copyright (c) 2021 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package webapi

import (
	"net/http"
	"strconv"
)

type errorKind string

const (
	errBadRequest                       = errorKind("errBadRequest")
	errInternalError                    = errorKind("errInternalError")
	errVspClosed                        = errorKind("errVspClosed")
	errFeeAlreadyReceived               = errorKind("errFeeAlreadyReceived")
	errInvalidFeeTx                     = errorKind("errInvalidFeeTx")
	errFeeTooSmall                      = errorKind("errFeeTooSmall")
	errUnknownTicket                    = errorKind("errUnknownTicket")
	errTicketCannotVote                 = errorKind("errTicketCannotVote")
	errFeeExpired                       = errorKind("errFeeExpired")
	errInvalidVoteChoices               = errorKind("errInvalidVoteChoices")
	errBadSignature                     = errorKind("errBadSignature")
	errInvalidPrivKey                   = errorKind("errInvalidPrivKey")
	errFeeNotReceived                   = errorKind("errFeeNotReceived")
	errInvalidTicket                    = errorKind("errInvalidTicket")
	errCannotBroadcastTicket            = errorKind("errCannotBroadcastTicket")
	errCannotBroadcastFee               = errorKind("errCannotBroadcastFee")
	errCannotBroadcastFeeUnknownOutputs = errorKind("errCannotBroadcastFeeUnknownOutputs")
	errInvalidTimestamp                 = errorKind("errInvalidTmestamp")

	errNotPublicKey               = errorKind("errNotPublicKey")
	errMultipleInvalidChildGenKey = errorKind("errMultipleInvalidChildGenKey")
	errZeroNetworkTicketPoolSize  = errorKind("errZeroNetworkTicketPoolSize")
	errNoVspClientSignature       = errorKind("errNoVspClientSignature")
	errInvalidTx                  = errorKind("errInvalidTx")
)

var defaultErr = map[errorKind]string{
	errBadRequest:                       "bad request",
	errInternalError:                    "internal error",
	errVspClosed:                        "vsp is closed",
	errFeeAlreadyReceived:               "fee tx already received for ticket",
	errInvalidFeeTx:                     "invalid fee tx",
	errFeeTooSmall:                      "fee too small",
	errUnknownTicket:                    "unknown ticket",
	errTicketCannotVote:                 "ticket not eligible to vote",
	errFeeExpired:                       "fee has expired",
	errInvalidVoteChoices:               "invalid vote choices",
	errBadSignature:                     "bad request signature",
	errInvalidPrivKey:                   "invalid private key",
	errFeeNotReceived:                   "no fee tx received for ticket",
	errInvalidTicket:                    "not a valid ticket tx",
	errCannotBroadcastTicket:            "ticket transaction could not be broadcast",
	errCannotBroadcastFee:               "fee transaction could not be broadcast",
	errCannotBroadcastFeeUnknownOutputs: "fee transaction could not be broadcast due to unknown outputs",
	errInvalidTimestamp:                 "old or reused timestamp",
}

func (e errorKind) Is(target error) bool {
	return target.Error() == defaultErr[e]
}

// Error maps application errorKind to HTTP status codes.
func (e errorKind) Error() string {
	switch e {
	case errBadRequest:
		return strconv.Itoa(http.StatusBadRequest)
	case errInternalError:
		return strconv.Itoa(http.StatusInternalServerError)
	case errVspClosed:
		return strconv.Itoa(http.StatusBadRequest)
	case errFeeAlreadyReceived:
		return strconv.Itoa(http.StatusBadRequest)
	case errInvalidFeeTx:
		return strconv.Itoa(http.StatusBadRequest)
	case errFeeTooSmall:
		return strconv.Itoa(http.StatusBadRequest)
	case errUnknownTicket:
		return strconv.Itoa(http.StatusBadRequest)
	case errTicketCannotVote:
		return strconv.Itoa(http.StatusBadRequest)
	case errFeeExpired:
		return strconv.Itoa(http.StatusBadRequest)
	case errInvalidVoteChoices:
		return strconv.Itoa(http.StatusBadRequest)
	case errBadSignature:
		return strconv.Itoa(http.StatusBadRequest)
	case errInvalidPrivKey:
		return strconv.Itoa(http.StatusBadRequest)
	case errFeeNotReceived:
		return strconv.Itoa(http.StatusBadRequest)
	case errInvalidTicket:
		return strconv.Itoa(http.StatusBadRequest)
	case errCannotBroadcastTicket:
		return strconv.Itoa(http.StatusInternalServerError)
	case errCannotBroadcastFee:
		return strconv.Itoa(http.StatusInternalServerError)
	case errCannotBroadcastFeeUnknownOutputs:
		return strconv.Itoa(http.StatusPreconditionRequired)
	case errInvalidTimestamp:
		return strconv.Itoa(http.StatusBadRequest)
	default:
		return strconv.Itoa(http.StatusInternalServerError)
	}
}

type ApiError struct {
	err         error
	description string
}

func (e ApiError) Unwrap() error {
	return e.err
}

func apiError(e errorKind, desc string) ApiError {
	return ApiError{err: e, description: desc}
}

// Error satisfies the error interface and prints human-readable errors for a given ErrorKind
func (e ApiError) Error() string {
	return e.description
}
