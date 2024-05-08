import { expect } from "chai";

describe("[Upgradable Test]", function () {
  let bridge;
  let operator;
  let guardian;
  let judge;
  let minOperatorRequiredConfirm;
  let minGuardianRequiredConfirm;
  let operator1;
  let operator2;
  let operator3;
  let operator4;
  let guardian1;
  let guardian2;
  let guardian3;
  let guardian4;
  let judge1;

  let sender;
  let receiver;
  let amount;
  let seq;
  let maxTryTransfer;

  beforeEach(async function () {
    upgrades.silenceWarnings()
    const [
      _operator1, _operator2, _operator3, _operator4,
      _guardian1, _guardian2, _guardian3, _guardian4, _judge1
    ] = await ethers.getSigners();
    operator1 = _operator1;
    operator2 = _operator2;
    operator3 = _operator3;
    operator4 = _operator4;
    guardian1 = _guardian1;
    guardian2 = _guardian2;
    guardian3 = _guardian3;
    guardian4 = _guardian4;
    judge1 = _judge1;
    minOperatorRequiredConfirm = 3;
    minGuardianRequiredConfirm = 3;
    maxTryTransfer = 0;

    const accs = [operator1, operator2, operator3, operator4, guardian1, guardian2, guardian3, guardian4];
    for (let acc of accs) {
      await hre.network.provider.send("hardhat_setBalance", [
        acc.address,
        "0x1000000000000000000000000000000000000",
      ]);
    }

    const guardianFactory = await ethers.getContractFactory("Guardian");
    guardian = await upgrades.deployProxy(guardianFactory, [[
      guardian1.address,
      guardian2.address,
      guardian3.address,
      guardian4.address,
    ], minGuardianRequiredConfirm]);

    const operatorFactory = await ethers.getContractFactory("Operator");
    operator = await upgrades.deployProxy(operatorFactory, [[
      operator1.address,
      operator2.address,
      operator3.address,
      operator4.address,
    ], guardian.address, minOperatorRequiredConfirm]);

    const judgeFactory = await ethers.getContractFactory("Judge");
    judge = await upgrades.deployProxy(judgeFactory, [[
      judge1.address
    ], guardian.address, 1]);

    const bridgeFactory = await ethers.getContractFactory("KAIABridge");
    bridge = await upgrades.deployProxy(bridgeFactory, [
      operator.address, guardian.address, judge.address, maxTryTransfer
    ]);
 
    await hre.network.provider.send("hardhat_setBalance", [
      bridge.address,
      "0x1000000000000000000000000000000000000",
    ]);

    sender = "0x0000000000000000000000000000000000000123";
    receiver = "0x0000000000000000000000000000000000000456";
    amount = 1;
    seq = 1;

    let rawTxData = (await operator.populateTransaction.changeBridge(bridge.address)).data;
    await guardian.connect(guardian1).submitTransaction(operator.address, rawTxData, 0);
    await guardian.connect(guardian2).confirmTransaction(1);
    await expect(guardian.connect(guardian3).confirmTransaction(1))
      .to.emit(operator, "ChangeBridge");

    // make and finalize a provision before on each test
    const provision = [seq, sender, receiver, amount];
    rawTxData = (await bridge.populateTransaction.provision(provision)).data;
    expect((await operator.getConfirmations(1)).length).to.equal(0);
    await operator.connect(operator1).submitTransaction(bridge.address, rawTxData, 0);
    await operator.connect(operator2).confirmTransaction(1);
    await operator.connect(operator3).confirmTransaction(1);
    expect((await operator.getConfirmations(1)).length).to.equal(3);
    expect(await bridge.greatestConfirmedSeq()).to.equal(1);

  })

  it("#bridge contract upgrade", async function () {
    // contract upgrade flow:
    // 1. pause the bridge
    // 2. upgrade the bridge (contract deployment and initialization)
    // 3. resume the bridge

    // 1. pause
    expect(await bridge.pause()).to.equal(false);
    let rawTxData = (await bridge.populateTransaction.pauseBridge("Bridge paused temporally")).data;
    await guardian.connect(guardian1).submitTransaction(bridge.address, rawTxData, 0);
    await guardian.connect(guardian2).confirmTransaction(2);
    await guardian.connect(guardian3).confirmTransaction(2);
    expect(await bridge.pause()).to.equal(true);

    // 2. deploy new version of contract (bridge)
    const newBridgeFactory = await ethers.getContractFactory("NewKAIABridge");
    const newBridge = await upgrades.upgradeProxy(bridge.address, newBridgeFactory);
    expect(await bridge.greatestConfirmedSeq()).to.equal(1);
    expect(await newBridge.greatestConfirmedSeq()).to.equal(1);

    // 3. resume the bridge
    rawTxData = (await newBridge.populateTransaction.resumeBridge("Bridge resumed")).data;
    await guardian.connect(guardian1).submitTransaction(newBridge.address, rawTxData, 0);
    await guardian.connect(guardian2).confirmTransaction(3);
    await guardian.connect(guardian3).confirmTransaction(3);
    expect(await newBridge.pause()).to.equal(false);

    // 4. make and finalize a new provision on the newly deployed bridge contract
    const provision = [seq + 1, sender, receiver, amount];
    rawTxData = (await newBridge.populateTransaction.provision(provision)).data;
    await operator.connect(operator1).submitTransaction(newBridge.address, rawTxData, 0);
    await operator.connect(operator2).confirmTransaction(2);
    await operator.connect(operator3).confirmTransaction(2);
    expect(await newBridge.greatestConfirmedSeq()).to.equal(2);

    expect(await newBridge.newFunc()).to.be.equal(123);
  })

  it("#operator contract upgrade", async function () {
    // 1. pause
    expect(await bridge.pause()).to.equal(false);
    let rawTxData = (await bridge.populateTransaction.pauseBridge("Bridge paused temporally")).data;
    await guardian.connect(guardian1).submitTransaction(bridge.address, rawTxData, 0);
    await guardian.connect(guardian2).confirmTransaction(2);
    await guardian.connect(guardian3).confirmTransaction(2);
    expect(await bridge.pause()).to.equal(true);

    // 2. deploy new operator contract and remove one operator
    const existingOperators = await operator.getOperators();
    const newOperatorFactory = await ethers.getContractFactory("Operator");
    const newOperator = await upgrades.upgradeProxy(operator.address, newOperatorFactory);
    expect(await newOperator.getOperators()).to.deep.equal(existingOperators)

    rawTxData = (await newOperator.populateTransaction.removeOperator(operator2.address)).data;
    await guardian.connect(guardian1).submitTransaction(newOperator.address, rawTxData, 0);
    await guardian.connect(guardian2).confirmTransaction(3);
    await guardian.connect(guardian3).confirmTransaction(3);
    expect((await newOperator.getOperators()).length).to.equal(existingOperators.length - 1);

    // 3. resume the bridge
    rawTxData = (await bridge.populateTransaction.resumeBridge("Bridge resumed")).data;
    await guardian.connect(guardian1).submitTransaction(bridge.address, rawTxData, 0);
    await guardian.connect(guardian3).confirmTransaction(4);
    await guardian.connect(guardian4).confirmTransaction(4);
    expect(await bridge.pause()).to.equal(false);
  })

  it("#enumerableSet upgradable test", async function () {
    const enumSetFactory = await ethers.getContractFactory("EnumSet");
    const enumSet = await upgrades.deployProxy(enumSetFactory, []);

    await enumSet.add(3);
    await enumSet.add(4);
    await enumSet.add(5);
    expect(await enumSet.getAll()).to.deep.equal([3,4,5])
    await enumSet.remove(4);

    const newEnumSetFactory = await ethers.getContractFactory("NewEnumSet");
    const newEnumSet = await upgrades.upgradeProxy(enumSet.address, newEnumSetFactory);

    expect(await newEnumSet.getAll()).to.deep.equal([3,5])
    await enumSet.remove(5);
    expect(await newEnumSet.getAll()).to.deep.equal([3])
  });
})
