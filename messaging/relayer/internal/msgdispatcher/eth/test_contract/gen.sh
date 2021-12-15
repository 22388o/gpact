#!/usr/bin/env bash

# This script is used to generate the msg_dispatcher_test_sol_wrapper.go from MsgDispatcherTest.sol

set -e
rm -rf build

solc MsgDispatcherTest.sol --bin --abi --optimize -o . --overwrite
~/go/bin/abigen --bin=MsgDispatcherTest.bin --abi=MsgDispatcherTest.abi --pkg=testmsgdispatcher --out=test_msg_dispatcher_sol_wrapper.go

solc --base-path ../../../../../../   ../../../../../attestor-sign/src/main/solidity/AttestorSignRegistrar.sol  --bin --abi --optimize -o . --overwrite
~/go/bin/abigen --bin=AttestorSignRegistrar.bin --abi=AttestorSignRegistrar.abi --pkg=test_registrar --out=test_registrar_sol_wrapper.go

solc --base-path ../../../../../../  ../../../../../event-relay/src/main/solidity/SignedEventStore.sol  --bin --abi --optimize -o . --overwrite
~/go/bin/abigen --bin=SignedEventStore.bin --abi=SignedEventStore.abi --pkg=testsignedeventstore --out=test_signed_event_store_sol_wrapper.go

solc --base-path ../../../../../../  ../../../../../../functioncall/sfc/src/main/solidity/SimpleCrosschainControl.sol  --bin --abi --optimize -o . --overwrite
~/go/bin/abigen --bin=SimpleCrosschainControl.bin --abi=SimpleCrosschainControl.abi --pkg=testsfc --out=test_sfc_sol_wrapper.go
