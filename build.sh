#!/usr/bin/env bash
set -e

./gradlew
cd messaging/event-relay/build
abigen --abi SignedEventStore.abi --bin SignedEventStore.bin --pkg eventstore --type EventStore --out eventstore.go
mv eventstore.go ../../../messaging/relayer/internal/soliditywrappers/eventstore
cd ../../attestor-sign/build
abigen --abi AttestorSignRegistrar.abi --bin AttestorSignRegistrar.bin --pkg registrar --type Registrar --out registrar.go
mv registrar.go ../../../messaging/relayer/internal/soliditywrappers/registrar
cd ../../../
cd functioncall/sfc/build
abigen --abi SimpleCrosschainControl.abi --bin SimpleCrosschainControl.bin --pkg sfc --type Sfc --out simple_crosschain_control.go
mv simple_crosschain_control.go ../../../messaging/relayer/internal/soliditywrappers/sfc
cd ../../../