#!/usr/bin/env bash

# This script is used to generate the msg_dispatcher_test_sol_wrapper.go from MsgDispatcherTest.sol

set -e
rm -rf build

solc MsgDispatcherTest.sol --bin --abi --optimize -o . --overwrite
~/go/bin/abigen --bin=MsgDispatcherTest.bin --abi=MsgDispatcherTest.abi --pkg=msgdispatchertest --out=msg_dispatcher_test_sol_wrapper.go