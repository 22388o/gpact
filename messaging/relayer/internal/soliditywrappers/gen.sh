#!/usr/bin/env bash

# This script is used to generate Solidity wrappers used for the relayer. The script is currently
# not part of the build, and hence needs to be run separately.
# To execute the script:
# * Be in the ./gpact directory.
# * sh messaging/relayer/internal/soliditywrappers/gen.sh
set -e

# Run this script from ./gpact
HERE=messaging/relayer/internal/soliditywrappers
BUILDDIR=$HERE/build

rm -rf $BUILDDIR

ABIGEN=~/go/bin/abigen

PKG=registrar
solc messaging/attestor-sign/src/main/solidity/AttestorSignRegistrar.sol  --base-path .  --bin --abi --hashes --optimize -o $BUILDDIR --overwrite
mkdir -p $HERE/$PKG
$ABIGEN --bin=$BUILDDIR/AttestorSignRegistrar.bin --abi=$BUILDDIR/AttestorSignRegistrar.abi --pkg=$PKG --out=$HERE/$PKG/attestor_sign_registrar.go

PKG=eventstore
solc messaging/event-relay/src/main/solidity/SignedEventStore.sol  --base-path .  --bin --abi --hashes --optimize -o $BUILDDIR --overwrite
mkdir -p $HERE/$PKG
$ABIGEN --bin=$BUILDDIR/SignedEventStore.bin --abi=$BUILDDIR/SignedEventStore.abi --pkg=$PKG --out=$HERE/$PKG/signed_event_store.go

PKG=sfc
solc functioncall/sfc/src/main/solidity/SimpleCrosschainControl.sol  --base-path .  --bin --abi --hashes --optimize -o $BUILDDIR --overwrite
mkdir -p $HERE/$PKG
$ABIGEN --bin=$BUILDDIR/SimpleCrosschainControl.bin --abi=$BUILDDIR/SimpleCrosschainControl.abi --pkg=$PKG --out=$HERE/$PKG/simple_crosschain_control.go

PKG=test
solc messaging/event-relay/src/test/solidity/AppTest.sol  --base-path .  --bin --abi --hashes --optimize -o $BUILDDIR --overwrite
mkdir -p $HERE/$PKG
$ABIGEN --bin=$BUILDDIR/AppTest.bin --abi=$BUILDDIR/AppTest.abi --pkg=$PKG --out=$HERE/$PKG/app_test.go
