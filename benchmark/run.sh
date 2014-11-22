#!/bin/bash

case $1 in
  server)
    netserver "${@:2}"
    ip addr show dev eth0
    echo "Enter to exit"
    read DONE
    echo "Done"
    ;;
  *)
    PORT=12865
    if [[ -n "$3" ]]; then
      PORT=$3
    fi
    netperf -H $1 -f m -p $PORT -v 2 -t TCP_STREAM
    echo
    netperf -H $1 -f m -p $PORT -v 2 -t TCP_RR
    echo
    netperf -H $1 -f m -p $PORT -v 2 -t TCP_CRR
    echo
    netperf -H $1 -f m -p $PORT -v 2 -t UDP_STREAM
    echo
    netperf -H $1 -f m -p $PORT -v 2 -t UDP_RR
    ;;
esac
