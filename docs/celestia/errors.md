# Celestia Node/Core/App Errors

Exhaustive list of errors returned by celestia node/core/app when trying to submit a PFB transaction. There might be a few duplicates because some errors like `sdkerrors` are wrapped in celestia-app.

## celestia-node

- `SubmitPayForBlob` - `state: no blobs provided`
- `SubmitPayForBlob` - `parsing insufficient min gas price error`
- `SubmitPayForBlob` - `sdkErrors.ABCIError`
- `SubmitPayForBlob` - `failed to submit blobs after %d attempts: %w`

## celestia-app

- `AnteHandle` -  `errors.Wrap(sdkerrors.ErrInsufficientFee, "not enough gas to pay for blobs")`
- `checkTxFeeWithValidatorMinGasPrices` - `errors.Wrap(sdkerror.ErrTxDecode, "Tx must be a FeeTx")`
- `ValidateBasic` - `ErrNoNamespaces`
- `ValidateBasic` - `ErrNoShareVersions`
- `ValidateBasic` - `ErrNoBlobSizes`
- `ValidateBasic` - `ErrNoShareCommitments`
- `ValidateBasic` - `ErrInvalidNamespace`
- `ValidateBasic` - `ErrUnsupportedShareVersion`
- `ValidateBasic` - `ErrInvalidShareCommitment`
- `ValidateBasic` - `ErrMismatchedNumberOfPFBComponent.Wrapf("namespaces %d blob sizes %d share versions %d share commitments %d)`
- `ValidateBlobNamespace` - `ErrReservedNamespace`
- `ValidateBlobNamespace` - `ErrInvalidNamespaceVersion`
- `ValidateBlobs` - `namespace version %d is too large`
- `ValidateBlobs` - `ErrNoBlobs`
- `ValidateBlobs` - `ErrZeroBlobSize`
- `ValidateBlobs` - `ErrUnsupportedShareVersion`
- `ParseInsufficientMinGasPrice` - `parsing insufficient min gas price error`
- `ParseInsufficientMinGasPrice` - `unexpected case: required gas price is zero`

## celestia-core

- `BroadcastTxCommit` - `ErrTimedOutWaitingForTx` - `timed out waiting for tx to be included in a block`
- `BroadcastTxCommit` - `max_subscription_clients %d reached`
- `BroadcastTxCommit` - `max_subscriptions_per_client %d reached`
- `BroadcastTxCommit` - `failed to subscribe to tx: %w`
- `BroadcastTxCommit` - `error on broadcastTxCommit: %v`
- `BroadcastTxCommit` - `deliverTxSub was cancelled (reason: %s)`
- `PreCheck` - `tx size is too big`
- `CheckTx` - `ErrTxTooLarge` - `Tx too large. Max size is %d, but got %d`
- `CheckTx` - `ErrPreCheck` - cosmos-sdk errors, see below
- `CheckTx` - `ErrTxInCache` - `tx already exists in cache`
- `CheckTx` - `ErrMempoolIsFull` - `mempool is full: number of txs %d (max: %d), total txs bytes %d (max: %d)`
- `CheckTx` - `mem.proxyAppConn.Error()` - `proxyAppConn may error if tx buffer is full`
- `PostCheck` - `gas wanted %d is negative`
- `PostCheck` - `gas wanted %d is greater than max gas`
- `TryAddNewTx` - `ErrTxInMempool` - `tx already exists in mempool`
- `TryAddNewTx` - `ErrTxAlreadyRejected` - `tx was previously rejected`
- `TryAddNewTx` - `ErrPreCheck`
- `TryAddNewTx` - `mem.proxyAppConn.Error()`
- `TryAddNewTx` - `application rejected transaction with code %d (Log: %s)`
- `TryAddNewTx` - `rejected bad transaction after post check: %w`

## cosmos-sdk

`ErrTxDecode` - `tx parse error`
`ErrInvalidSequence` - `invalid sequence`
`ErrUnauthorized` - `unauthorized`
`ErrInsufficientFunds` - `insufficient funds`
`ErrUnknownRequest` - `unknown request`
`ErrInvalidAddress` - `invalid address`
`ErrInvalidPubKey` - `invalid pubkey`
`ErrUnknownAddress` - `unknown address`
`ErrInvalidCoins` - `invalid coins`
`ErrOutOfGas` - `out of gas`
`ErrMemoTooLarge` - `memo too large`
`ErrInsufficientFee` - `insufficient fee`
`ErrTooManySignatures` - `maximum number of signatures exceeded`
`ErrNoSignatures` - `no signatures supplied`
`ErrJSONMarshal` - `failed to marshal JSON bytes`
`ErrJSONUnmarshal` - `failed to unmarshal JSON bytes`
`ErrInvalidRequest` - `invalid request`
`ErrTxInMempoolCache` - `tx already in mempool`
`ErrMempoolIsFull` - `mempool is full`
`ErrTxTooLarge` - `tx too large`
`ErrKeyNotFound` - `key not found`
`ErrWrongPassword` - `invalid account password`
`ErrorInvalidSigner` - `tx intended signer does not match the given signer`
`ErrorInvalidGasAdjustment` - `invalid gas adjustment`
`ErrInvalidHeight` - `invalid height`
`ErrInvalidVersion` - `invalid version`
`ErrInvalidChainID` - `invalid chain-id`
`ErrInvalidType` - `invalid type`
`ErrTxTimeoutHeight` - `tx timeout height`
`ErrUnknownExtensionOptions` - `unknown extension options`
`ErrWrongSequence` - `incorrect account sequence`
`ErrPackAny` - `failed packing protobuf message to Any`
`ErrUnpackAny` - `failed unpacking protobuf message from Any`
`ErrLogic` - `internal logic error`
`ErrConflict` - `conflict`
`ErrNotSupported` - `feature not supported`
`ErrNotFound` - `not found`
`ErrIO` - `Internal IO error`
`ErrAppConfig` - `error in app.toml`
`ErrInvalidGasLimit` - `invalid gas limit`
`ErrPanic` - errorsmod.ErrPanic
