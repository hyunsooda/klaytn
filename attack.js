peer = admin.peers[0].id.substring(0,16)
for (i=0; i<100; i++) admin.syncStakingInfo(peer, 0, 0)
exit
