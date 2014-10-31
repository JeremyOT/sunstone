Benchmarks
==========

Control
-------
```
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00       142.94   -2934.86   -5440.08   -401.008  -743.313
root@sunstone3:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
^[[A
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00       142.22   -3014.82   -5509.88   -414.029  -756.678
root@sunstone3:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00       140.06   -3429.75   -5459.89   -478.264  -761.357
root@sunstone3:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
^[[A
^[[A
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       13132      0       82.0     -4744.91   -1129.659
212992           10.00       13132              82.0     -12393.64   -2950.651

root@sunstone3:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       13341      0       83.3     -4324.69   -1013.536
212992           10.00       13341              83.3     -12338.67   -2891.700

root@sunstone3:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       13077      0       81.7     -4399.87   -1051.929
212992           10.00       13077              81.7     -12322.96   -2946.197
```

Sunstone Host
---------
```
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00       157.72   -1289.93   -5299.89   -159.743  -656.329
root@sunstone1:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00       159.13   -1564.92   -5254.90   -192.071  -644.963
root@sunstone1:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00       156.59   -2039.90   -5309.94   -254.429  -662.290
root@sunstone1:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00       156.99   -1359.94   -5234.94   -169.194  -651.293
root@sunstone1:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
^[[A
^[[A
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       13856      0       86.6     -3994.73   -901.406
212992           10.00       13856              86.6     -12249.53   -2764.095

root@sunstone1:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       13482      0       84.2     -4394.67   -1019.244
212992           10.00       13481              84.2     -12319.14   -2857.146

root@sunstone1:/# netperf -H 172.16.19.130 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.16.19.130 (172.16.19.130) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       13194      0       82.4     -4469.89   -1059.188
212992           10.00       13194              82.4     -12319.19   -2919.165
```

Sunstone
----
```
root@cd29fdb3d141:/# netperf -H 172.17.130.3 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.3 (172.17.130.3) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00        63.59   -4149.72   -5874.82   -1274.631  -1804.513
root@cd29fdb3d141:/# netperf -H 172.17.130.3 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.3 (172.17.130.3) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00        67.89   -4194.84   -5654.88   -1206.763  -1626.785
root@cd29fdb3d141:/# netperf -H 172.17.130.3 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.3 (172.17.130.3) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00        61.60   -4169.71   -5924.80   -1322.162  -1878.681
root@cd29fdb3d141:/# netperf -H 172.17.130.3 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.3 (172.17.130.3) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       11960      0       74.7     -2879.89   -756.827
212992           10.00       11897              74.3     -11808.94   -3103.360

root@cd29fdb3d141:/# netperf -H 172.17.130.3 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.3 (172.17.130.3) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       11811      0       73.8     -2764.87   -735.009
212992           10.00       11761              73.5     -11728.13   -3117.786

root@cd29fdb3d141:/# netperf -H 172.17.130.3 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.3 (172.17.130.3) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       11894      0       74.3     -2824.79   -743.185
212992           10.00       11884              74.2     -11763.28   -3094.844
```

Sunstone MTU
--------
```
root@sunstone1:/# netperf -H 172.17.130.4 -f M -c 1 -C 1 -p 12865
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00        67.69   -3234.84   -5584.96   -933.351  -1611.433
root@sunstone1:/# netperf -H 172.17.130.4 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.4 (172.17.130.4) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00        61.12   -3414.86   -5595.07   -1091.193  -1787.862
root@sunstone1:/# netperf -H 172.17.130.4 -f M -c 1 -C 1 -p 12865
TCP STREAM TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.4 (172.17.130.4) port 0 AF_INET : demo
Recv   Send    Send                          Utilization       Service Demand
Socket Socket  Message  Elapsed              Send     Recv     Send    Recv
Size   Size    Size     Time     Throughput  local    remote   local   remote
bytes  bytes   bytes    secs.    MBytes  /s  % S      % S      us/KB   us/KB

 87380  16384  16384    10.00        73.12   -1904.89   -5265.15   -508.845  -1406.460
root@sunstone1:/# netperf -H 172.17.130.4 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.4 (172.17.130.4) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       15510      0       96.9     -3219.77   -649.103
212992           10.00       15509              96.9     -10973.64   -2212.276

root@sunstone1:/# netperf -H 172.17.130.4 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.4 (172.17.130.4) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       15560      0       97.2     -3009.86   -604.781
212992           10.00       15560              97.2     -10988.48   -2207.953

root@sunstone1:/# netperf -H 172.17.130.4 -f M -c 1 -C 1 -p 12865 -t UDP_STREAM
UDP UNIDIRECTIONAL SEND TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 172.17.130.4 (172.17.130.4) port 0 AF_INET : demo
Socket  Message  Elapsed      Messages                   CPU      Service
Size    Size     Time         Okay Errors   Throughput   Util     Demand
bytes   bytes    secs            #      #   MBytes/sec % SS     us/KB

212992   65507   10.00       15721      0       98.2     -2839.83   -564.816
212992           10.00       15720              98.2     -10989.72   -2185.755
```
