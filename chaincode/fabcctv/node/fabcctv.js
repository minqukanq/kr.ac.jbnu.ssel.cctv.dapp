/*
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
*/

'use strict';
const shim = require('fabric-shim');
const util = require('util');

let Chaincode = class {

  // The Init method is called when the Smart Contract 'fabcctv' is instantiated by the blockchain network
  // Best practice is to have any Ledger initialization in separate function -- see initLedger()
  async Init(stub) {
    console.info('=========== Instantiated fabcctv chaincode ===========');
    return shim.success();
  }

  // The Invoke method is called as a result of an application request to run the Smart Contract
  // 'fabcctv'. The calling application program has also specified the particular smart contract
  // function to be called, with arguments
  async Invoke(stub) {
    let ret = stub.getFunctionAndParameters();
    console.info(ret);

    let method = this[ret.fcn];
    if (!method) {
      console.error('no function of name:' + ret.fcn + ' found');
      throw new Error('Received unknown function ' + ret.fcn + ' invocation');
    }
    try {
      let payload = await method(stub, ret.params);
      return shim.success(payload);
    } catch (err) {
      console.log(err);
      return shim.error(err);
    }
  }

  async queryFrame(stub, args) {
    if (args.length != 1) {
      throw new Error('Incorrect number of arguments. Expecting frameKey ex:199502221420');
    }
    let frameKey = args[0];

    let frameAsBytes = await stub.getState(frameKey); // get the frame from chaincode state
    if (!frameAsBytes || frameAsBytes.toString().length <= 0) {
      throw new Error(frameKey + ' does not exist: ');
    }
    console.log(frameAsBytes.toString());
    return frameAsBytes;
  }

  async validateFrame(stub, args) {

    if (args.length != 2) {
      throw new Error('Incorrect number of arguments. Expecting 2');
    }
    let frameKey = args[0];

    let frameAsBytes = await stub.getState(frameKey); // get the frame from chaincode state
    if (!frameAsBytes || frameAsBytes.toString().length <= 0) {
      throw new Error(frameKey + ' does not exist: ');
    }
    console.log(frameAsBytes.toString());

    if (frameAsBytes.toString() != args[1]) {
      return frameAsBytes;
    }

    return shim.success();
  }

  async initLedger(stub, args) {
    console.info('============= START : Initialize Ledger ===========');
    let frames = [];
    frames.push({
      time: "199502221420", // TODO: check the format of date is good or not
      data: "Hash" // TODO: insert sample hash image hash value
    });

    for (let i = 0; i < frames.length; i++) {
      frames[i].docType = 'frame';
      await stub.putState('FRAME' + i, Buffer.from(JSON.stringify(frames[i])));
      console.info('Added <--> ', frames[i]);
    }
    console.info('============= END : Initialize Ledger ===========');
  }

  async createFrame(stub, args) {
    console.info('============= START : Create Frame ===========');
    if (args.length != 3) {
      throw new Error('Incorrect number of arguments. Expecting 3');
    }

    var frame = {
      docType: 'frame',
      key: args[1],
      value: args[2]
    };

    await stub.putState(args[0], Buffer.from(JSON.stringify(frame)));
    console.info('============= END : Create Frame ===========');
  }

  async queryFramesByRange(stub, args) {

    if (args.length != 1) {
      throw new Error('Incorrect number of arguments. Expecting 1');
    }

    let startKey = '000000000000'; // this value should not being changed
    let endKey = args[0];

    let iterator = await stub.getStateByRange(startKey, endKey);

    let allResults = [];
    while (true) {
      let res = await iterator.next();

      if (res.value && res.value.value.toString()) {
        let jsonRes = {};
        console.log(res.value.value.toString('utf8'));

        jsonRes.Key = res.value.key;
        try {
          jsonRes.Record = JSON.parse(res.value.value.toString('utf8'));
        } catch (err) {
          console.log(err);
          jsonRes.Record = res.value.value.toString('utf8');
        }
        allResults.push(jsonRes);
      }
      if (res.done) {
        console.log('end of data');
        await iterator.close();
        console.info(allResults);
        return Buffer.from(JSON.stringify(allResults));
      }
    }
  }
};

shim.start(new Chaincode());
