# Go V1 to V2

- `TopicUpdateTransaction` missing `[get|set]ExpirationTime()`
- `AccountUpdateTransaction` missing `SetInitialBalance` and `GetInitialBalance`

v2.0.0

- Removed `Build(*Client)` from all transactions
- Renamed `Ed25519PublicKey` → `PublicKey`
    - Added `Verify([]byte, []byte) bool`
        - Verifies a message was signe by the respective private key.
    - Added `VerifyTransaction(Transaction) bool`
        - Verifies the transaction was signed by the respective private key.
- Renamed `Ed25519PrivateKey` → `PrivateKey`
    - Added `SignTransaction(Transaction) []byte`
        - Signs the `Transaction` and returns the signature.
- Removed `ThresholdKey`
    - Use `KeyListWithThreshold(uint) *KeyList`
- Added `Key`
    - Added by `PublicKey`
    - Added by `PrivateKey`
    - Added by `KeyList`
    - Added by `ContractID`
- `KeyList`
    - Added `AddAllPublicKeys([]PublicKey) *KeyList`
- `Mnemonic`
    - Added `ToLegacyPrivateKey() (PrivateKey, error)`
    - Added `GenerateMnemonic24() (Mnemonic, error)`
    - Added `GenerateMnemonic12() (Mnemonic, error)`
    - Removed `GenerateMnemonic() (Mnemonic, error)`
        - Use `GenerateMnemonic12()` or `GenerateMnemonic24()` instead.
- Removed `MirrorClient`
    - Use `Client` instead, and set the mirror network using `SetMirrorNetwork()`
- Renamed `MirrorSubscriptionHandle` → `SubscriptionHandle`
- `QueryBuilder` → `Query`
    - Removed `SetPaymentTransaction()`
    - Added `GetNodeAccountIDs() []AccountID`
    - Added `SetNodeAccountIDs([]AccountID) *Query`
    - Added `GetMaxRetryCount() int`
    - Added `SetMaxRetry(int) *Query`
- `TransactionBuilder` and `Transaction`
    - Added `TransactionFromBytes([]byte) (interface{}, error)`
    - Added `ToBytes() []byte`
    - Renamed `ID()` →`GetTransactionId()`
    - Added `GetMaxTransactionFee() Hbar`
    - Added `GetTransactionMemo() String`
    - Added `GetTransactionHashPerNode() (map[AccountID][]byte, error)`
    - Added `GetTransactionValidDuration() time.Duration`
    - Added `AddSignature(PublicKey, byte[]) Transaction`
    - Added `GetSignatures() (map[AccountID]map[*PublicKey][]byte, error)`
    - Added `FreezeWith(Client) (Transaction, error)`
    - Added `Freeze() (Transaction, error)`
    - Removed `UnmarshalBinary([]byte) error`
    - Removed `MarshalBinary() ([]byte, error)`
    - Renamed `SetNodeAccountID(AccountID)` → `SetNodeAccountIDs([]AccountID)`
    - Added `GetTransactionHash() ([]byte, error)`
    - Added `SetMaxRetry(int) *Transaction`
    - Added `GetMaxRetry() int`
    - Added `IsFrozen() bool`
    - Changed `Execute(Client) (TransactionID, error)` → `Execute(Client) (TransactionResponse, error)`
    - Changed `Sign(Ed25519PrivateKey)` → `Sign(PrivateKey)`
    - Changed `SignWith(Ed25519PublicKey)` → `SignWith(PublicKey, TransactionSigner)`
- `AccountBalanceQuery` extends (Query)
    - Renamed `BalanceQuery` → `Query`
    - Changed `Execute(client *Client) (Hbar, error)` → `Execute(client *Client) (AccountBalance, error)`
    - Added `GetAccountID() AccountID`
    - Added `GetContractID() ContactID`
- Added `AccountBalance`
    - Added `Hbars Hbar`
    - Added `Token map[TokenID]uint64`
- `AccountCreateTransaction` extends (Transaction)
    - Added `GetKey() (Key, error)`
    - Added `GetInitialBalance() Hbar`
    - Added `GetReceiverSignatureRequired() boolean`
    - Added `GetProxyAccountID() AccountID`
    - Added `GetAutoRenewPeriod() time.Duration`
    - Removed `SetSendRecordThreshold(Hbar)`
    - Removed `SetReceiveRecordThreshold(Hbar)`
- `AccountDeleteTransaction` extends (Transaction)
    - Renamed `SetDeleteAccountID()` → `setAccountID()`
    - Added `GetAccountID() AccountID`
    - Added `GetTransferAccountID() AccountID`
- `AccountID`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) AccountID`
- `AccountInfo`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) AccountInfo`
    - Added `TokenRelationships: []*TokenRelationship`
- `AccountInfoQuery` extends (Query)
    - Added `GetAccountID() AccountID`
    - Renamed `Cost(Client)` → `GetCost(Client)`
- `AccountRecordsQuery` extends (Query)
    - Added `GetAccountID() AccountID`
- `AccountStakersQuery` extends (Query)
    - Added `GetAccountID() AccountID`
- `AccountUpdateTransaction` extends (Transaction)
    - Added `GetAccountID() AccountID`
    - Added `GetKey() (Key, error)`
    - Added `GetInitialBalance() Hbar`
    - Added `GetReceiverSignatureRequired() boolean`
    - Added `GetProxyAccountID() AccountID`
    - Added `GetAutoRenewPeriod() time.Duration`
    - Added `GetExpirationTime() time.Time`
    - Removed `SetSendRecordThreshold(Hbar)`
    - Removed  `SetReceiveRecordThreshold(Hbar)`
- Removed `CryptoTransferTranscation`
    - Use `TransferTransaction` instead.
- `TransferTransaction` extends [`Transaction`]
    - Added `AddTokenTransfer(TokenID, AccountID, int64) TransferTransaction`
    - Added `GetTokenTransfers() map[TokenID][]TokenTransfer`
    - Added `AddHbarTransfer(AccountID, Hbar) TransferTransaction`
    - Added `GetHbarTransfers() map[AccountID]Hbar`
- Renamed `ContractBytecodeQuery` → `ContractByteCodeQuery` extends (Query)
    - Added `GetContractID() ContractID`
- `ContractCallQuery` extends (Query)
    - Added `GetContractID() ContractID`
    - Added `GetGas() uint64`
    - Added `GetFunctionParameters() []byte`
    - Added `SetFunctionParameters([]byte) *ContractCallQuery`
- `ContractCreateTransaction` extends (Transaction)
    - Added `GetBytecodeFileID() FileID`
    - Added `GetAdminKey() (Key, error)`
    - Added `GetGas() uint64`
    - Added `GetInitialBalance() Hbar`
    - Added `GetAutoRenewDuration() time.Duration`
    - Added `GetProxyAccountID() AccountID`
    - Added `GetContractMemo() String`
    - Added `GetConstructorParameters() []byte`
    - Added `SetConstructorParametersRaw([]byte) *ContractCreateTransaction`
    - Removed `SetInitialBalance()`
- `ContractDeleteTransaction` extends (Transaction)
    - Added `GetContractID() ContractID`
    - Added `GetTransferAccountID() AccountID`
    - Added `GetTransferContractID() ContractID`
- `ContractExecuteTransaction` extends (Transaction)
    - Added `GetContractID() ContractID`
    - Added `GetGas() uint64`
    - Added `GetPayableAmount() Hbar`
    - Added `byte[] getFunctionParameters()`
- `ContractID`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) ContractID`
- `ContractInfo`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) ContractInfo`
- `ContractInfoQuery` extends (Query)
    - Added `GetContractID() ContractID`
- Removed `ContractRecordsQuery`
- `ContractUpdateTransaction` extends (Transaction)
    - Added `GetContractID() ContractID`
    - Added `GetBytecodeFileID() FileID`
    - Added `GetAdminKey() (Key, error)`
    - Added `GetAutoRenewDuration() time.Duration`
    - Added `GetProxyAccountID() AccountID`
    - Added `GetContractMemo() String`
    - Added `GetExpirationTime() time.Time`
- `FileAppendTransaction`
    - Added `GetFileID() FileID`
    - Added `GetContents() []byte`
- `FileContentsQuery`
    - Added `GetFileID() FileID`
- `FileCreateTransaction`
    - Added `GetContents() []byte`
    - Added `GetKeys() KeyList`
    - Added `GetExpirationTime() time.Time`
    - Renamed `AddKey(PublicKey)` → `setKeys(Key...)`
- `FileDeleteTransaction`
    - Added `GetFileID() FileID`
- `FileID`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) FileID`
- `FileInfo`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) FileInfo`
    - Update `Keys []PublicKey` → `Keys KeyList`
- `FileInfoQuery`
    - Added `GetFileID() FileID`
- `FileUpdateTransaction`
    - Added `GetFileID() FileID`
    - Added `GetContents() []byte`
    - Added `GetKeys() KeyList`
    - Added `GetExpirationTime() time.Time`
    - Renamed `AddKey(PublicKey)` → `setKeys(Key...)`
- Removed `ConsensusTopicMessage`
- Renamed `MirrorConsensusTopicResponse` → `TopicMessage`
    - Added `TopicMessageChunk[] Chunks`
        - This will be non null for a topic message which is constructed from multiple transactions.
    - Renamed `message []byte` → `contents []byte`
    - Removed `GetMessage() []byte`
    - Removed `ConsensusTopicID TopicID`
- Renamed `MirrorConsensusTopicChunk` → `TopicMessageChunk`
- Renamed `MirrorTopicMessageQuery` → `TopicMessageQuery`
    - Change `Subscribe(MirrorClient, func(MirrorConsensusTopicResponse), func(error)) (MirrorSubscriptionHandle, error)`→ `Subscribe(*Client, func(TopicMessage)) (SubscriptionHandle, error)`
- Renamed `ConsensusTopicCreateTransaction` → `TopicCreateTransaction`
    - Added `GetTopicMemo() String`
    - Added `GetAdminKey() (Key, error)`
    - Added `GetSubmitKey() (Key, error)`
    - Added `GetAutoRenewDuration() time.Duration`
    - Added `GetAutoRenewAccountID() AccountID`
- Renamed `ConsensusTopicDeleteTransaction` → `TopicDeleteTransaction`
    - Added `GetTopicID() TopicID`
- Renamed `ConsensusMessageSubmitTransaction` → `TopicMessageSubmitTransaction`
    - Added `GetTopicID() TopicID`
    - Added `GetMessage() []byte`
    - Removed `SetChunkInfo(TransactionId, int, int)`
    - Added `GetMaxChunks() uint64`
- Renamed `ConsensusTopicID` → `TopicID`
- Renamed `ConsensusTopicInfo` → `TopicInfo`
    - Added `ToBytes() []byte`
    - Added `FromBytes() TopicInfo`
    - Change `AdminKey Ed25519PublicKey`  and  `SubmitKey Ed25519PublicKey`→ `AdminKey Key` and `SubmitKey Key`
- Renamed `ConsensusTopicInfoQuery` → `TopicInfoQuery`
    - Added `GetTopicID() TopicID`
- Renamed `ConsensusTopicUpdateTransaction` → `TopicUpdateTransaction`
    - Added `GetTopicID() TopicID`
    - Added `GetTopicMemo() String`
    - Added `GetAdminKey() (Key, error)`
    - Added `GetSubmitKey() (Key, error)`
    - Added `GetAutoRenewDuration() time.Duration`
    - Added `GetAutoRenewAccountID() AccountID`
- `TokenAssociateTransaction` extends (Transaction)
    - Added `GetAccountID() AccountID`
    - Added  `SetTokenIDs([]TokenID)`
    - Added `GetTokenIDs() []TokenID`
- Removed `TokenBalanceQuery`
    - Use `AccountBalanceQuery` to fetch token balances since `AccountBalance` contains `tokenBalances`.
- `TokenBurnTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
    - Added `GetAmount() uint64`
- `TokenCreateTransaction` extends (Transaction)
    - Renamed `SetName(String)` →`setTokenName(String)`
    - Added `GetTokenName() string`
    - Renamed `SetSymbol(String)` →`setTokenSymbol(String)`
    - Added `GetTokenSymbol() string`
    - Renamed `SetTreasury(AccountID)` →`setTreasuryAccountID(AccountID)`
    - Renamed `SetAutoRenewAccount(AccountID)` →`setAutoRenewAccountID(AccountID)`
    - Added `GetAutoRenewAccountID() AccountID`
    - Added `GetTreasuryAccountID() AccountID`
    - Added `GetAdminKey() (Key, error)`
    - Added `GetKycKey() (Key, error)`
    - Added `GetSupplyKey() (Key, error)`
    - Added `GetWipeKey() (Key, error)`
    - Added `GetFreezeKey() (Key, error)`
    - Added `GetFreezeDefault() boolean`
    - Added `GetExpirationTime() time.Time`
    - Added `GetAutoRenewAccountID() AccountID`
    - Added `GetAutoRenewPeriod() time.Duration`
    - Added `GetDecimals() int`
- `TokenDeleteTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
- `TokenDisassociateTransaction` extends (Transaction)
    - Added `GetAccountID() AccountID`
    - Added `GetTokenIDs() []TokenID`
    - Added `SetTokenIDs([]TokenID)`
- `TokenFreezeTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
    - Added `GetAccointId() AccountID`
- `TokenGrantKycTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
    - Added `GetAccointId() AccountID`
- `TokenID`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) TokenID`
- `TokenInfo`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) TokenInfo`
- `TokenInfoQuery` extends (Query)
    - Added `GetTokenID() TokenID`
- `TokenMintTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
    - Added `GetAmount() uint64`
- `TokenRelationship`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) TokenRelationship`
- `TokenRevokeKycTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
    - Added `GetAccointId() AccountID`
- Removed `TokenTransferTransaction`
    - Use `TransferTransaction` instead.
- `TokenUnfreezeTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
    - Added `GetAccountId() AccountID`
- `TokenUpdateTransaction` extends (Transaction)
    - Renamed `SetName(String)` →`setTokenName(String)`
    - Added `GetTokenName() string`
    - Renamed `SetSymbol(String)` →`setTokenSymbol(String)`
    - Added `GetTokenSymbol() string`
    - Renamed `SetTreasury(AccountID)` →`setTreasuryAccountID(AccountID)`
    - Renamed `SetAutoRenewAccount(AccountID)` →`setAutoRenewAccountID(AccountID)`
    - Added `GetAutoRenewAccountID() AccountID`
    - Added `GetTreasuryAccountID() AccountID`
    - Added `GetAdminKey() (Key, error)`
    - Added `GetKycKey() (Key, error)`
    - Added `GetSupplyKey() (Key, error)`
    - Added `GetWipeKey() (Key, error)`
    - Added `GetFreezeKey() (Key, error)`
    - Added `GetFreezeDefault() boolean`
    - Added `GetExpirationTime() time.Time`
    - Added `GetAutoRenewAccountID() AccountID`
    - Added `GetAutoRenewPeriod() time.Duration`
    - Added `GetDecimals() int`
- `TokenWipeTransaction` extends (Transaction)
    - Added `GetTokenID() TokenID`
    - Added `GetAccountID() AccountID`
- `FreezeTransaction`
    - Added `GetStartTime() time.Time`
    - Added `GetEndTime() time.Time`
- Removed `HbarRangeException`
    - If `Hbar` is out of range `Hedera` will error instead.
- Removed `HederaConstants`
    - No replacement.
- Removed `HederaNetworkException`
- Renamed `HederaPrecheckStatusException` → `PrecheckStatusException`
- Renamed `HederaReceiptStatusException` → `ReceiptStatusException`
- Removed `HederaRecordStatusException`
    - `ReceiptStatusException` will be thrown instead.
- Removed `HederaStatusException`
    - A `PrecheckStatusException` or `ReceiptStatusException` will be thrown instead.
- Removed `HederaThrowable`
    - No replacement.
- Removed `LocalValidationException`
    - No replacement. Local validation is no longer done.
- `SystemDeleteTransaction`
    - Added `GetFileID() FileID`
    - Added `GetContractID() ContractID`
    - Added `GetExpirationTime() time.Time`
- `SystemUndeleteTransaction`
    - Added `GetFileID() FileID`
    - Added `GetContractID() ContractID`
- `TransactionId`
    - Added `ToBytes() []byte`
    - Added `FromBytes(byte[]) TransactionId`
    - Removed `TransactionId(TransactionIDOrBuilder)`
    - Removed `ToProto() TransactionID`
    - Removed `WithValidStart(AccountID, time.Time) TransactionId`
        - Use `TransactionId(AccountID, time.Time) new` instead.
    - Removed `TransactionId(AccountID)`
        - Use `Generate(AccountID) TransactionId` instead.
- Removed `TransactionList`
- `TransactionReciept`
    - Added `ToBytes() []byte`
    - Added `FromBytes() TransactionReceipt`
    - Expose `ExchangeRate *ExchangeRate`
    - Expose `AccountID *AccountID`
    - Expose `FileID *FileID`
    - Expose `ContractID *ContractID`
    - Expose `TopicID *TopicID`
    - Expose `TokenID *TokenID`
    - Expose `TopicSequenceNumber uint64`
    - Expose `topicRunningHash []byte`
    - Added `TotalSupply uint64`
    - Added `TopicRunningHashVersion uint64`
    - Removed `GetAccountID() AccountID`
        - Use `AccountID AccountID` directly instead.
    - Removed `GetContractID() ContractID`
        - Use `ContractID ContractID` directly instead.
    - Removed `GetFileID() FileID`
        - Use `FileID FileID` directly instead.
    - Removed `GetTokenID() TokenID`
        - Use `TokenID TokenID` directly instead.
    - Removed `GetConsensusTopicID() ConsensusTopicID`
        - Use `TopicID TopicID` directly instead.
    - Removed `GetConsensusTopicSequenceNumber()`
        - Use `sequenceNumber uint64` directly instead.
    - Removed `GetConsensusTopicRunningHash() []byte`
        - Use `topicRunningHash []byte` directly instead.
    - Removed `ToProto() TransactionReceipt`
- `TransactionReceiptQuery` extends (Query)
    - Added `GetTransactionId() TransactionID`
- `TransactionRecord`
    - Added `ToBytes() []byte`
    - Added `FromBytes() TransactionRecord`
    - Added `CallResult *ContractFunctionResult`
    - Added `CallResultIsCreate bool`
    - Removed `GetContratcExecuteResult() ContractFunctionResult`
        - Use `ContractFunctionResult contractFunctionResult` directly instead.
    - Removed `GetContratcCreateResult() ContractFunctionResult`
        - Use `ContractFunctionResult contractFunctionResult` directly instead.
    - Removed `ToProto() TransactionReceipt`
- `TransactionRecordQuery` extends (Query)
    - Added `GetTransactionId() TransactionId`
- `Hbar`
    - Added  `ToString(unit HbarUnit) string`
- `Client`
    - Revomed `NewClient(map[string]AccountID) *Client`
    - Added `SetMirrorNetwork([]String) void`
    - Added `GetMirrorNetwork() []String`
    - Added `ClientForNetwork(map[string]AccountID) *Client`
    - Added `Ping(AccountID) void`
    - Added `GetOperatorPublicKey() PublicKey`
    - Added `SetNetwork(map[string]AccountID) error`
    - Added `GetNetwork() map[string]AccountID`
    - Renamed `FromJson([]byte)` → `ClientFromConfig(byte[])`
    - Renamed `FromFile(String)` → `ClientFromConfigFile(String)`
    - Renamed `GetOperatorId()` → `getOperatorAccountID()`
    - Removed `ReplaceNodes(map[string]AccountID) *Client`
    - Removed `SetMaxTransactionFee() Client`
    - Removed `SetMaxQueryPayment() Client`