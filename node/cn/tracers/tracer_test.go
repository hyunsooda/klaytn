// Copyright 2018 The klaytn Authors
// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
//
// This file is derived from eth/tracers/tracer_test.go (2018/06/04).
// Modified and improved for the klaytn development.

package tracers

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/klaytn/klaytn/blockchain/state"
	"github.com/klaytn/klaytn/blockchain/vm"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/params"
)

type account struct{}

func (account) SubBalance(amount *big.Int)                          {}
func (account) AddBalance(amount *big.Int)                          {}
func (account) SetAddress(common.Address)                           {}
func (account) Value() *big.Int                                     { return nil }
func (account) SetBalance(*big.Int)                                 {}
func (account) SetNonce(uint64)                                     {}
func (account) Balance() *big.Int                                   { return nil }
func (account) Address() common.Address                             { return common.Address{} }
func (account) FeePayer() common.Address                            { return common.Address{} }
func (account) ReturnGas(*big.Int)                                  {}
func (account) SetCode(common.Hash, []byte)                         {}
func (account) ForEachStorage(cb func(key, value common.Hash) bool) {}

type dummyStatedb struct {
	state.StateDB
}

func (*dummyStatedb) GetRefund() uint64 { return 1337 }

func runTrace(tracer *Tracer) (json.RawMessage, error) {
	env := vm.NewEVM(vm.Context{BlockNumber: big.NewInt(1)}, &dummyStatedb{}, params.TestChainConfig, &vm.Config{Debug: true, Tracer: tracer})

	contract := vm.NewContract(account{}, account{}, big.NewInt(0), 10000)
	contract.Code = []byte{byte(vm.PUSH1), 0x1, byte(vm.PUSH1), 0x1, 0x0}

	_, err := env.Interpreter().Run(contract, []byte{})
	if err != nil {
		return nil, err
	}
	return tracer.GetResult()
}

// TestRegressionPanicSlice tests that we don't panic on bad arguments to memory access
func TestRegressionPanicSlice(t *testing.T) {
	tracer, err := New("{depths: [], step: function(log) { this.depths.push(log.memory.slice(-1,-2)); }, fault: function() {}, result: function() { return this.depths; }}", true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = runTrace(tracer); err != nil {
		t.Fatal(err)
	}
}

// TestRegressionPanicSlice tests that we don't panic on bad arguments to stack peeks
func TestRegressionPanicPeek(t *testing.T) {
	tracer, err := New("{depths: [], step: function(log) { this.depths.push(log.stack.peek(-1)); }, fault: function() {}, result: function() { return this.depths; }}", true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = runTrace(tracer); err != nil {
		t.Fatal(err)
	}
}

// TestRegressionPanicSlice tests that we don't panic on bad arguments to memory getUint
func TestRegressionPanicGetUint(t *testing.T) {
	tracer, err := New("{ depths: [], step: function(log, db) { this.depths.push(log.memory.getUint(-64));}, fault: function() {}, result: function() { return this.depths; }}", true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = runTrace(tracer); err != nil {
		t.Fatal(err)
	}
}

// TestTracingDeepObject tests if it returns an expected error when the json object has too many recursive children
func TestTracingDeepObject(t *testing.T) {
	tracer, err := New("{step: function() {}, fault: function() {}, result: function() { var o={}; var x=o; for (var i=0; i<1000; i++){ o.foo={}; o=o.foo; } return x; }}", true)
	if err != nil {
		t.Fatal(err)
	}

	_, err = runTrace(tracer)
	expectedErr := `RangeError: json encode recursion limit    in server-side tracer function 'result'`
	if !strings.Contains(err.Error(), expectedErr) {
		t.Errorf("Expected return error to be %s, got %v", expectedErr, err)
	}
}

func TestTracing(t *testing.T) {
	tracer, err := New("{count: 0, step: function() { this.count += 1; }, fault: function() {}, result: function() { return this.count; }}", true)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := runTrace(tracer)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ret, []byte("3")) {
		t.Errorf("Expected return value to be 3, got %s", string(ret))
	}
}

func TestUnsafeTracingDisabled(t *testing.T) {
	_, err := New("{count: 0, step: function() { this.count += 1; }, fault: function() {}, result: function() { return this.count; }}", false)
	if err == nil || err.Error() != "Only predefined tracers are supported" {
		t.Fatal("Must disable JS code based tracers if unsafe")
	}
}

func TestStack(t *testing.T) {
	tracer, err := New("{depths: [], step: function(log) { this.depths.push(log.stack.length()); }, fault: function() {}, result: function() { return this.depths; }}", true)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := runTrace(tracer)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ret, []byte("[0,1,2]")) {
		t.Errorf("Expected return value to be [0,1,2], got %s", string(ret))
	}
}

func TestOpcodes(t *testing.T) {
	tracer, err := New("{opcodes: [], step: function(log) { this.opcodes.push(log.op.toString()); }, fault: function() {}, result: function() { return this.opcodes; }}", true)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := runTrace(tracer)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ret, []byte("[\"PUSH1\",\"PUSH1\",\"STOP\"]")) {
		t.Errorf("Expected return value to be [\"PUSH1\",\"PUSH1\",\"STOP\"], got %s", string(ret))
	}
}

func TestHalt(t *testing.T) {
	t.Skip("duktape doesn't support abortion")

	timeout := errors.New("stahp")
	tracer, err := New("{step: function() { while(1); }, result: function() { return null; }}", true)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		tracer.Stop(timeout)
	}()

	if _, err = runTrace(tracer); err.Error() != "stahp    in server-side tracer function 'step'" {
		t.Errorf("Expected timeout error, got %v", err)
	}
}

func TestHaltBetweenSteps(t *testing.T) {
	tracer, err := New("{step: function() {}, fault: function() {}, result: function() { return null; }}", true)
	if err != nil {
		t.Fatal(err)
	}

	env := vm.NewEVM(vm.Context{BlockNumber: big.NewInt(1)}, &dummyStatedb{}, params.TestChainConfig, &vm.Config{Debug: true, Tracer: tracer})
	contract := vm.NewContract(&account{}, &account{}, big.NewInt(0), 0)

	tracer.CaptureState(env, 0, 0, 0, 0, nil, nil, contract, 0, nil)
	timeout := errors.New("stahp")
	tracer.Stop(timeout)
	tracer.CaptureState(env, 0, 0, 0, 0, nil, nil, contract, 0, nil)

	if _, err := tracer.GetResult(); err.Error() != timeout.Error() {
		t.Errorf("Expected timeout error, got %v", err)
	}
}
