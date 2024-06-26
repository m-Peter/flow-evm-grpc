const { assert } = require('chai')
const conf = require('./config')
const helpers = require('./helpers')
const web3 = conf.web3

it('store revertReason field in transaction receipts', async () => {
    let deployed = await helpers.deployContract("storage")
    let contractAddress = deployed.receipt.contractAddress

    // make sure deploy was successful
    assert.equal(deployed.receipt.status, conf.successStatus)

    // assert that the receipt for the contract deployment transaction
    // does not have any revert reason.
    // Note: `revertReason` field is dropped from the result when
    // using `web3.eth.getTransactionReceipt`, that's why we need
    // to use `helpers.callRPCMethod`.
    let receipt = await helpers.callRPCMethod(
        'eth_getTransactionReceipt',
        [deployed.receipt.transactionHash]
    )
    assert.isUndefined(receipt.body['result'].revertReason)

    // we construct a transaction that reverts on purpose, with an assertion error
    let callAssertError = deployed.contract.methods.assertError().encodeABI()
    let assertErrorTx = {
        from: conf.eoa.address,
        to: contractAddress,
        data: callAssertError,
        gas: 1_000_000,
        gasPrice: 0
    }
    let signedTx = await conf.eoa.signTransaction(assertErrorTx)

    // we need to use `helpers.callRPCMethod` to test this out,
    // because `web3.eth.sendSignedTransaction` will not let
    // through transactions that revert, it will error out.
    let response = await helpers.callRPCMethod(
        'eth_sendRawTransaction',
        [signedTx.rawTransaction]
    )
    assert.equal(200, response.status)

    let latestHeight = await web3.eth.getBlockNumber()
    let block = await web3.eth.getBlock(latestHeight)
    assert.equal(block.number, 5n)

    let revertedTx = await web3.eth.getTransactionFromBlock(latestHeight, 0)
    // Give some time to the engine to ingest the latest transaction
    await new Promise(res => setTimeout(res, 1500))
    rcp = await helpers.callRPCMethod(
        'eth_getTransactionReceipt',
        [revertedTx.hash]
    )
    // make sure the `revertReason` field is included in the response
    assert.equal(
        rcp.body['result'].revertReason,
        '0x08c379a000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000014417373657274204572726f72204d657373616765000000000000000000000000'
    )

    // we construct a transaction that reverts on purpose, with a custom error
    let callCustomError = deployed.contract.methods.customError().encodeABI()
    let customErrorTx = {
        from: conf.eoa.address,
        to: contractAddress,
        data: callCustomError,
        gas: 1_000_000,
        gasPrice: 0
    }
    signedTx = await conf.eoa.signTransaction(customErrorTx)

    response = await helpers.callRPCMethod(
        'eth_sendRawTransaction',
        [signedTx.rawTransaction]
    )
    assert.equal(200, response.status)

    latestHeight = await web3.eth.getBlockNumber()
    block = await web3.eth.getBlock(latestHeight)
    assert.equal(block.number, 6n)

    revertedTx = await web3.eth.getTransactionFromBlock(latestHeight, 0)
    // Give some time to the engine to ingest the latest transaction
    await new Promise(res => setTimeout(res, 1500))
    rcp = await helpers.callRPCMethod(
        'eth_getTransactionReceipt',
        [revertedTx.hash]
    )
    // make sure the `revertReason` field is included in the response
    assert.equal(
        rcp.body['result'].revertReason,
        '0x9195785a00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001056616c756520697320746f6f206c6f7700000000000000000000000000000000'
    )

}).timeout(20 * 1000)
