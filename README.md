## Hyperledger Fabric Environment Sample

This contains the basic environment setting scripts and files that is required to accomplish.

## License <a name="license"></a>

Hyperledger Project source code files are made available under the Apache License, Version 2.0 (Apache-2.0), located in the [LICENSE](LICENSE) file. Hyperledger Project documentation files are made available under the Creative Commons Attribution 4.0 International License (CC-BY-4), available at http://creativecommons.org/licenses/by/4.0/.

## How to Use?

### Build network with simple script

- If you didn't have network by far, you can use below bash script to building network.

```sh
  # start the network from the bottom
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger$ cd fabcctv/
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger/fabcctv$ source ./startFabric.sh
```

- If you do have working network already, please kill(not a stop) that network using below script.

```sh
  # remove the existing images
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger$ cd basic-network/
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger/basic-network$ source ./teardown.sh
```

- If you do need to stop the docker containers for some reason, use below command.

```sh
  # stop the running container
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger$ cd basic-network/
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger/basic-network$ source ./stop.sh
```

### Add/Remove Chaincode

- How to add chaincode (per one chaincode)

```sh
  # add files that required to be deployed into the network
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger/chaincode$ touch fabcctv/go/newchaincode.go
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger/chaincode$ touch fabcctv/node/newchaincode.js

  # install required node module
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger/chaincode$ cd fabcctv/node/
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger/chaincode/fabcctv/node$ npm install

  # remove the existing images
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger$ source basic-network/.teardown.sh

  # start the network from the bottom
  # this make the changes being adapted into the network.
  ssel@SSEL:~/kr.ac.jbnu.ssel.env.hyperledger$ source fabcctv/.startFabric.sh
```

- Chaincode will be deployed into the network(e.g.,channel) automatically once the network raised up.

- chaincode/ directory will have all the chaincode that supposed to be deployed into the network
