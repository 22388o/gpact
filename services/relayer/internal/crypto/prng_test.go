package crypto

/*
 * Copyright 2021 ConsenSys Software Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Check that the function executes without causing a panic.
func TestSystemRandom(t *testing.T) {
	b := make([]byte, 32)
	zeroBytes := make([]byte, 32)
	systemRandomBytes(b)

	assert.NotEqual(t, b, zeroBytes, "Random bytes all zero")
}

// This commented out code can be used to check what accuracy the nano second
// timer actually gives on a given platform.
// On a MacBookPro 2016, the time appears to only update each micro-second.
func TestNanoTime(t *testing.T) {
	t1 := time.Now().UnixNano()
	t2 := time.Now().UnixNano()
	interval := t2 - t1
	fmt.Println(interval)
}

// Make sure the get MAC access isn't just returning zeros.
func TestMacAddress(t *testing.T) {
	b := getMACAddr()
	zeroBytes := make([]byte, len(b))
	assert.NotEqual(t, b, zeroBytes, "Random bytes all zero")
}

func TestNewPRNG(t *testing.T) {
	b1 := make([]byte, 32)
	b2 := make([]byte, 32)
	zeroBytes := make([]byte, 32)

	securityDomain := []byte("test")
	r := NewPRNG(securityDomain)
	r.ReadBytes(b1)
	r.ReadBytes(b2)

	// The probability of the two random numbers being the same or zero is one in
	// 2**256 and two in 2**256
	assert.NotEqual(t, b1, zeroBytes, "b1 bytes all zero")
	assert.NotEqual(t, b2, zeroBytes, "b2 bytes all zero")
	assert.NotEqual(t, b1, b2, "successive calls to generate random data generated the same bytes")
}

// Check that the various lengths of random numbers can be generated,
// with a focus on edge cases.
func TestReadDifferentLengths(t *testing.T) {
	b0 := make([]byte, 0)
	b1 := make([]byte, 1)
	b31 := make([]byte, 31)
	b33 := make([]byte, 33)
	b63 := make([]byte, 63)
	b64 := make([]byte, 64)
	b65 := make([]byte, 65)
	b10000 := make([]byte, 10000)

	securityDomain := []byte("test")
	r := NewPRNG(securityDomain)
	r.ReadBytes(b0)
	r.ReadBytes(b1)
	r.ReadBytes(b31)
	r.ReadBytes(b33)
	r.ReadBytes(b63)
	r.ReadBytes(b64)
	r.ReadBytes(b65)
	r.ReadBytes(b10000)
}

func TestTwoSecurityDomains(t *testing.T) {
	b1 := make([]byte, 32)
	b2 := make([]byte, 32)
	zeroBytes := make([]byte, 32)

	securityDomain1 := []byte("1")
	r1 := NewPRNG(securityDomain1)
	r1.ReadBytes(b1)

	securityDomain2 := []byte("2")
	r2 := NewPRNG(securityDomain2)
	r2.ReadBytes(b2)

	// The probability of the two random numbers being the same or zero is one in
	// 2**256 and two in 2**256
	assert.NotEqual(t, b1, zeroBytes, "b1 bytes all zero")
	assert.NotEqual(t, b2, zeroBytes, "b2 bytes all zero")
	assert.NotEqual(t, b1, b2, "random values generated by separate PRNGs were the same")
}

// Make sure no panics are called during seed.
func TestSeed(t *testing.T) {
	zeroBytes := make([]byte, 32)
	securityDomain := []byte("test")
	r := NewPRNG(securityDomain)
	r.Reseed(zeroBytes)
}

// Make sure no panics are called during quick reseed.
func TestQuickSeed(t *testing.T) {
	securityDomain := []byte("test")
	r := NewPRNG(securityDomain)
	r.QuickReseedKick()
}

func TestGetReader(t *testing.T) {
	securityDomain := []byte("test")
	r := NewPRNG(securityDomain)
	rand := r.GetReader()

	b := make([]byte, 32)
	l, err := rand.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, l, len(b))
}
