# raft_kv_go

Key value storage system server with raft algorithm and grpc.

This project will hold 2 binaries, 1 for the kv store and 1 for a scheduler.


Code structure:

kv - holds grpc server
proto - all the nice proto stuff ;)
scheduler - this will schedul stuff, not sure if I'll keep it separate
shared? - kv and scheduler will most likely share some stuff :?


resources:

https://grpc.io/
