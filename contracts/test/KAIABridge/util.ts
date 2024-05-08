export const DAY = 24 * 60 * 60;
export const WEEK = 7 * DAY;
export const YEAR = 365 * DAY;

export async function setBlockTimestamp(time: number, mine = true) : void {
  await ethers.provider.send("evm_increaseTime", [time]);
  if (mine) {
    await ethers.provider.send("evm_mine", []);
  }
}

export async function getBalance(addr: string) : bigInt {
  return await ethers.provider.getBalance(addr);
}

export async function getCurrentBlockTimestamp() : bigInt {
  const blockNumber = await ethers.provider.getBlockNumber();
  const block = await ethers.provider.getBlock(blockNumber);
  return block.timestamp;
}
