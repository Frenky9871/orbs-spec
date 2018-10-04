// AUTO GENERATED FILE (by membufc proto compiler v0.0.19)
package protocol

import (
)

/////////////////////////////////////////////////////////////////////////////
// enums

type ErrorCodes uint16

const (
	ERROR_CODE_RESERVED ErrorCodes = 0
	ERROR_CODE_OUT_OF_SYNC ErrorCodes = 1
	ERROR_CODE_INVALID ErrorCodes = 2
)

func (n ErrorCodes) String() string {
	switch n {
	case ERROR_CODE_RESERVED:
		return "ERROR_CODE_RESERVED"
	case ERROR_CODE_OUT_OF_SYNC:
		return "ERROR_CODE_OUT_OF_SYNC"
	case ERROR_CODE_INVALID:
		return "ERROR_CODE_INVALID"
	}
	return "UNKNOWN"
}

type ExecutionResult uint16

const (
	EXECUTION_RESULT_RESERVED ExecutionResult = 0
	EXECUTION_RESULT_SUCCESS ExecutionResult = 1
	EXECUTION_RESULT_ERROR_SMART_CONTRACT ExecutionResult = 2
	EXECUTION_RESULT_ERROR_INPUT ExecutionResult = 3
	EXECUTION_RESULT_ERROR_UNEXPECTED ExecutionResult = 4
)

func (n ExecutionResult) String() string {
	switch n {
	case EXECUTION_RESULT_RESERVED:
		return "EXECUTION_RESULT_RESERVED"
	case EXECUTION_RESULT_SUCCESS:
		return "EXECUTION_RESULT_SUCCESS"
	case EXECUTION_RESULT_ERROR_SMART_CONTRACT:
		return "EXECUTION_RESULT_ERROR_SMART_CONTRACT"
	case EXECUTION_RESULT_ERROR_INPUT:
		return "EXECUTION_RESULT_ERROR_INPUT"
	case EXECUTION_RESULT_ERROR_UNEXPECTED:
		return "EXECUTION_RESULT_ERROR_UNEXPECTED"
	}
	return "UNKNOWN"
}

type TransactionStatus uint16

const (
	TRANSACTION_STATUS_RESERVED TransactionStatus = 0
	TRANSACTION_STATUS_COMMITTED TransactionStatus = 1
	TRANSACTION_STATUS_DUPLICATE_TRANSACTION_ALREADY_COMMITTED TransactionStatus = 2
	TRANSACTION_STATUS_PENDING TransactionStatus = 3
	TRANSACTION_STATUS_DUPLICATE_TRANSACTION_ALREADY_PENDING TransactionStatus = 4
	TRANSACTION_STATUS_PRE_ORDER_VALID TransactionStatus = 5
	TRANSACTION_STATUS_NO_RECORD_FOUND TransactionStatus = 6
	TRANSACTION_STATUS_REJECTED_UNSUPPORTED_VERSION TransactionStatus = 7
	TRANSACTION_STATUS_REJECTED_VIRTUAL_CHAIN_MISMATCH TransactionStatus = 8
	TRANSACTION_STATUS_REJECTED_TIMESTAMP_WINDOW_EXCEEDED TransactionStatus = 9
	TRANSACTION_STATUS_REJECTED_SIGNATURE_MISMATCH TransactionStatus = 10
	TRANSACTION_STATUS_REJECTED_UNKNOWN_SIGNER_SCHEME TransactionStatus = 11
	TRANSACTION_STATUS_REJECTED_GLOBAL_PRE_ORDER TransactionStatus = 12
	TRANSACTION_STATUS_REJECTED_VIRTUAL_CHAIN_PRE_ORDER TransactionStatus = 13
	TRANSACTION_STATUS_REJECTED_SMART_CONTRACT_PRE_ORDER TransactionStatus = 14
	TRANSACTION_STATUS_REJECTED_TIMESTAMP_AHEAD_OF_NODE_TIME TransactionStatus = 15
	TRANSACTION_STATUS_REJECTED_CONGESTION TransactionStatus = 16
)

func (n TransactionStatus) String() string {
	switch n {
	case TRANSACTION_STATUS_RESERVED:
		return "TRANSACTION_STATUS_RESERVED"
	case TRANSACTION_STATUS_COMMITTED:
		return "TRANSACTION_STATUS_COMMITTED"
	case TRANSACTION_STATUS_DUPLICATE_TRANSACTION_ALREADY_COMMITTED:
		return "TRANSACTION_STATUS_DUPLICATE_TRANSACTION_ALREADY_COMMITTED"
	case TRANSACTION_STATUS_PENDING:
		return "TRANSACTION_STATUS_PENDING"
	case TRANSACTION_STATUS_DUPLICATE_TRANSACTION_ALREADY_PENDING:
		return "TRANSACTION_STATUS_DUPLICATE_TRANSACTION_ALREADY_PENDING"
	case TRANSACTION_STATUS_PRE_ORDER_VALID:
		return "TRANSACTION_STATUS_PRE_ORDER_VALID"
	case TRANSACTION_STATUS_NO_RECORD_FOUND:
		return "TRANSACTION_STATUS_NO_RECORD_FOUND"
	case TRANSACTION_STATUS_REJECTED_UNSUPPORTED_VERSION:
		return "TRANSACTION_STATUS_REJECTED_UNSUPPORTED_VERSION"
	case TRANSACTION_STATUS_REJECTED_VIRTUAL_CHAIN_MISMATCH:
		return "TRANSACTION_STATUS_REJECTED_VIRTUAL_CHAIN_MISMATCH"
	case TRANSACTION_STATUS_REJECTED_TIMESTAMP_WINDOW_EXCEEDED:
		return "TRANSACTION_STATUS_REJECTED_TIMESTAMP_WINDOW_EXCEEDED"
	case TRANSACTION_STATUS_REJECTED_SIGNATURE_MISMATCH:
		return "TRANSACTION_STATUS_REJECTED_SIGNATURE_MISMATCH"
	case TRANSACTION_STATUS_REJECTED_UNKNOWN_SIGNER_SCHEME:
		return "TRANSACTION_STATUS_REJECTED_UNKNOWN_SIGNER_SCHEME"
	case TRANSACTION_STATUS_REJECTED_GLOBAL_PRE_ORDER:
		return "TRANSACTION_STATUS_REJECTED_GLOBAL_PRE_ORDER"
	case TRANSACTION_STATUS_REJECTED_VIRTUAL_CHAIN_PRE_ORDER:
		return "TRANSACTION_STATUS_REJECTED_VIRTUAL_CHAIN_PRE_ORDER"
	case TRANSACTION_STATUS_REJECTED_SMART_CONTRACT_PRE_ORDER:
		return "TRANSACTION_STATUS_REJECTED_SMART_CONTRACT_PRE_ORDER"
	case TRANSACTION_STATUS_REJECTED_TIMESTAMP_AHEAD_OF_NODE_TIME:
		return "TRANSACTION_STATUS_REJECTED_TIMESTAMP_AHEAD_OF_NODE_TIME"
	case TRANSACTION_STATUS_REJECTED_CONGESTION:
		return "TRANSACTION_STATUS_REJECTED_CONGESTION"
	}
	return "UNKNOWN"
}

type RequestStatus uint16

const (
	REQUEST_STATUS_RESERVED RequestStatus = 0
	REQUEST_STATUS_COMPLETED RequestStatus = 1
	REQUEST_STATUS_IN_PROCESS RequestStatus = 2
	REQUEST_STATUS_NOT_FOUND RequestStatus = 3
	REQUEST_STATUS_REJECTED RequestStatus = 4
	REQUEST_STATUS_CONGESTION RequestStatus = 5
	REQUEST_STATUS_SYSTEM_ERROR RequestStatus = 6
)

func (n RequestStatus) String() string {
	switch n {
	case REQUEST_STATUS_RESERVED:
		return "REQUEST_STATUS_RESERVED"
	case REQUEST_STATUS_COMPLETED:
		return "REQUEST_STATUS_COMPLETED"
	case REQUEST_STATUS_IN_PROCESS:
		return "REQUEST_STATUS_IN_PROCESS"
	case REQUEST_STATUS_NOT_FOUND:
		return "REQUEST_STATUS_NOT_FOUND"
	case REQUEST_STATUS_REJECTED:
		return "REQUEST_STATUS_REJECTED"
	case REQUEST_STATUS_CONGESTION:
		return "REQUEST_STATUS_CONGESTION"
	case REQUEST_STATUS_SYSTEM_ERROR:
		return "REQUEST_STATUS_SYSTEM_ERROR"
	}
	return "UNKNOWN"
}

