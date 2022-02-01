//go:build all || unit
// +build all unit

package hedera

import (
	"context"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	protobuf "google.golang.org/protobuf/proto"

	"github.com/hashgraph/hedera-protobufs-go/services"
	"google.golang.org/grpc"
)

var wg sync.WaitGroup

func DisabledTestUnitMock(t *testing.T) {
	responses := []interface{}{
		&services.TransactionResponse{
			NodeTransactionPrecheckCode: services.ResponseCodeEnum_OK,
		},
		&services.Response{
			Response: &services.Response_TransactionGetReceipt{
				TransactionGetReceipt: &services.TransactionGetReceiptResponse{
					Header: &services.ResponseHeader{
						Cost:         0,
						ResponseType: services.ResponseType_COST_ANSWER,
					},
				},
			},
		},
		&services.Response{
			Response: &services.Response_TransactionGetReceipt{
				TransactionGetReceipt: &services.TransactionGetReceiptResponse{
					Header: &services.ResponseHeader{
						Cost:         0,
						ResponseType: services.ResponseType_ANSWER_ONLY,
					},
					Receipt: &services.TransactionReceipt{
						Status: services.ResponseCodeEnum_SUCCESS,
						AccountID: &services.AccountID{Account: &services.AccountID_AccountNum{
							AccountNum: 234,
						}},
					},
				},
			},
		},
	}

	client, server := NewMockClientAndServer(responses)

	newKey, err := PrivateKeyGenerateEd25519()
	require.NoError(t, err)

	newBalance := NewHbar(2)

	tran := TransactionIDGenerate(AccountID{Account: 3})

	resp, err := NewAccountCreateTransaction().
		SetKey(newKey).
		SetTransactionID(tran).
		SetInitialBalance(newBalance).
		SetMaxAutomaticTokenAssociations(100).
		Execute(client)
	require.NoError(t, err)

	_, err = resp.GetReceipt(client)
	require.NoError(t, err)
	//assert.Equal(t, balance.Hbars.tinybar, int64(10))
	if server != nil {
		server.Stop()
	}
}

func TestGenerateTransactionIDsPerExecution(t *testing.T) {
	count := 0
	transactionIds := make(map[string]bool)

	call := func(request *services.Transaction) *services.TransactionResponse {
		var response *services.TransactionResponse

		require.NotEmpty(t, request.SignedTransactionBytes)
		signedTransaction := services.SignedTransaction{}
		_ = protobuf.Unmarshal(request.SignedTransactionBytes, &signedTransaction)

		require.NotEmpty(t, signedTransaction.BodyBytes)
		transactionBody := services.TransactionBody{}
		_ = protobuf.Unmarshal(signedTransaction.BodyBytes, &transactionBody)

		require.NotNil(t, transactionBody.TransactionID)
		transactionId := transactionBody.TransactionID.String()
		require.NotEqual(t, "", transactionId)
		require.False(t, transactionIds[transactionId])
		transactionIds[transactionId] = true

		sigMap := signedTransaction.GetSigMap()
		require.NotNil(t, sigMap)
		require.NotEqual(t, 0, len(sigMap.SigPair))

		for _, sigPair := range sigMap.SigPair {
			verified := false

			switch k := sigPair.Signature.(type) {
			case *services.SignaturePair_Ed25519:
				pbTemp, _ := PublicKeyFromBytesEd25519(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.Ed25519)
			case *services.SignaturePair_ECDSASecp256K1:
				pbTemp, _ := PublicKeyFromBytesECDSA(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.ECDSASecp256K1)
			}
			require.True(t, verified)
		}

		if count < 2 {
			response = &services.TransactionResponse{
				NodeTransactionPrecheckCode: services.ResponseCodeEnum_TRANSACTION_EXPIRED,
			}
		} else {
			response = &services.TransactionResponse{
				NodeTransactionPrecheckCode: services.ResponseCodeEnum_OK,
			}
		}

		count += 1

		return response
	}
	responses := []interface{}{
		call, call, call,
	}

	client, server := NewMockClientAndServer(responses)

	_, err := NewFileCreateTransaction().
		SetContents([]byte("hello")).
		Execute(client)
	require.NoError(t, err)

	if server != nil {
		server.Stop()
	}
}

func DisabledTestSingleTransactionIDForExecutions(t *testing.T) {
	count := 0
	tran := TransactionIDGenerate(AccountID{Account: 1800})
	transactionIds := make(map[string]bool)
	transactionIds[tran._ToProtobuf().String()] = true

	call := func(request *services.Transaction) *services.TransactionResponse {
		var response *services.TransactionResponse

		require.NotEmpty(t, request.SignedTransactionBytes)
		signedTransaction := services.SignedTransaction{}
		_ = protobuf.Unmarshal(request.SignedTransactionBytes, &signedTransaction)

		require.NotEmpty(t, signedTransaction.BodyBytes)
		transactionBody := services.TransactionBody{}
		_ = protobuf.Unmarshal(signedTransaction.BodyBytes, &transactionBody)

		require.NotNil(t, transactionBody.TransactionID)
		transactionId := transactionBody.TransactionID.String()
		require.NotEqual(t, "", transactionId)
		require.True(t, transactionIds[transactionId])
		transactionIds[transactionId] = true

		sigMap := signedTransaction.GetSigMap()
		require.NotNil(t, sigMap)
		require.NotEqual(t, 0, len(sigMap.SigPair))

		for _, sigPair := range sigMap.SigPair {
			verified := false

			switch k := sigPair.Signature.(type) {
			case *services.SignaturePair_Ed25519:
				pbTemp, _ := PublicKeyFromBytesEd25519(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.Ed25519)
			case *services.SignaturePair_ECDSASecp256K1:
				pbTemp, _ := PublicKeyFromBytesECDSA(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.ECDSASecp256K1)
			}
			require.True(t, verified)
		}

		if count < 2 {
			response = &services.TransactionResponse{
				NodeTransactionPrecheckCode: services.ResponseCodeEnum_BUSY,
			}
		} else {
			response = &services.TransactionResponse{
				NodeTransactionPrecheckCode: services.ResponseCodeEnum_OK,
			}
		}

		count += 1

		return response
	}
	responses := []interface{}{
		call, call, call,
	}

	client, server := NewMockClientAndServer(responses)

	_, err := NewFileCreateTransaction().
		SetTransactionID(tran).
		SetContents([]byte("hello")).
		Execute(client)
	require.NoError(t, err)

	if server != nil {
		server.Stop()
	}
}

func DisabledTestSingleTransactionIDForExecutionsWithTimeout(t *testing.T) {
	count := 0
	tran := TransactionIDGenerate(AccountID{Account: 1800})
	transactionIds := make(map[string]bool)
	transactionIds[tran._ToProtobuf().String()] = true

	call := func(request *services.Transaction) *services.TransactionResponse {
		var response *services.TransactionResponse

		require.NotEmpty(t, request.SignedTransactionBytes)
		signedTransaction := services.SignedTransaction{}
		_ = protobuf.Unmarshal(request.SignedTransactionBytes, &signedTransaction)

		require.NotEmpty(t, signedTransaction.BodyBytes)
		transactionBody := services.TransactionBody{}
		_ = protobuf.Unmarshal(signedTransaction.BodyBytes, &transactionBody)

		require.NotNil(t, transactionBody.TransactionID)
		transactionId := transactionBody.TransactionID.String()
		require.NotEqual(t, "", transactionId)
		require.True(t, transactionIds[transactionId])
		transactionIds[transactionId] = true

		sigMap := signedTransaction.GetSigMap()
		require.NotNil(t, sigMap)
		require.NotEqual(t, 0, len(sigMap.SigPair))

		for _, sigPair := range sigMap.SigPair {
			verified := false

			switch k := sigPair.Signature.(type) {
			case *services.SignaturePair_Ed25519:
				pbTemp, _ := PublicKeyFromBytesEd25519(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.Ed25519)
			case *services.SignaturePair_ECDSASecp256K1:
				pbTemp, _ := PublicKeyFromBytesECDSA(sigPair.PubKeyPrefix)
				verified = pbTemp.Verify(signedTransaction.BodyBytes, k.ECDSASecp256K1)
			}
			require.True(t, verified)
		}

		if count < 2 {
			response = &services.TransactionResponse{
				NodeTransactionPrecheckCode: services.ResponseCodeEnum_TRANSACTION_EXPIRED,
			}
		} else {
			response = &services.TransactionResponse{
				NodeTransactionPrecheckCode: services.ResponseCodeEnum_OK,
			}
		}

		count += 1

		return response
	}
	responses := []interface{}{
		call, call, call,
	}

	client, server := NewMockClientAndServer(responses)

	_, err := NewFileCreateTransaction().
		SetTransactionID(tran).
		SetContents([]byte("hello")).
		Execute(client)
	require.Error(t, err)

	if server != nil {
		server.Stop()
		wg.Wait()
	}
}

func NewMockClientAndServer(responses []interface{}) (*Client, *grpc.Server) {
	client := ClientForNetwork(map[string]AccountID{
		"0.localhost:50211": {Account: 3},
		"2.localhost:50211": {Account: 4},
		"3.localhost:50211": {Account: 5},
	})

	key, _ := PrivateKeyFromStringEd25519("302e020100300506032b657004220420d45e1557156908c967804615af59a000be88c7aa7058bfcbe0f46b16c28f887d")
	client.SetOperator(AccountID{Account: 1800}, key)

	var server *grpc.Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		server = NewServer(responses)
	}()

	return client, server
}

func NewMockHandler(responses []interface{}) func(interface{}, context.Context, func(interface{}) error, grpc.UnaryServerInterceptor) (interface{}, error) {
	index := 0
	return func(_srv interface{}, _ctx context.Context, dec func(interface{}) error, _interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
		response := responses[index]
		index = index + 1

		switch response := response.(type) {
		case error:
			return nil, response
		case *services.TransactionResponse:
			return response, nil
		case *services.Response:
			return response, nil
		case func(request *services.Transaction) *services.TransactionResponse:
			request := new(services.Transaction)
			if err := dec(request); err != nil {
				return nil, err
			}
			return response(request), nil
		case func(request *services.Query) *services.Response:
			request := new(services.Query)
			if err := dec(request); err != nil {
				return nil, err
			}
			return response(request), nil
		default:
			return response, nil
		}
	}
}

func NewServer(responses []interface{}) *grpc.Server {
	server := grpc.NewServer()
	handler := NewMockHandler(responses)

	server.RegisterService(NewServiceDescription(handler, &services.CryptoService_ServiceDesc), nil)
	server.RegisterService(NewServiceDescription(handler, &services.FileService_ServiceDesc), nil)
	server.RegisterService(NewServiceDescription(handler, &services.SmartContractService_ServiceDesc), nil)
	server.RegisterService(NewServiceDescription(handler, &services.ConsensusService_ServiceDesc), nil)
	server.RegisterService(NewServiceDescription(handler, &services.TokenService_ServiceDesc), nil)
	server.RegisterService(NewServiceDescription(handler, &services.ScheduleService_ServiceDesc), nil)
	server.RegisterService(NewServiceDescription(handler, &services.FreezeService_ServiceDesc), nil)

	lis, err := net.Listen("tcp", "localhost:50211")
	if err != nil {
		panic(err)
	}

	if err = server.Serve(lis); err != nil {
		panic(err)
	}

	return server
}

func NewServiceDescription(handler func(interface{}, context.Context, func(interface{}) error, grpc.UnaryServerInterceptor) (interface{}, error), service *grpc.ServiceDesc) *grpc.ServiceDesc {
	var methods []grpc.MethodDesc
	for _, desc := range service.Methods {
		methods = append(methods, grpc.MethodDesc{
			MethodName: desc.MethodName,
			Handler:    handler,
		})
	}

	return &grpc.ServiceDesc{
		ServiceName: service.ServiceName,
		HandlerType: service.HandlerType,
		Methods:     methods,
		Streams:     []grpc.StreamDesc{},
		Metadata:    service.Metadata,
	}
}
