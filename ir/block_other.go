package ir

import (
	"github.com/umaumax/llvm/ir/enum"
	"github.com/umaumax/llvm/ir/types"
	"github.com/umaumax/llvm/ir/value"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewICmp appends a new icmp instruction to the basic block based on the given
// integer comparison predicate and integer scalar or vector operands.
func (block *Block) NewICmp(pred enum.IPred, x, y value.Value) *InstICmp {
	inst := NewICmp(pred, x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFCmp appends a new fcmp instruction to the basic block based on the given
// floating-point comparison predicate and floating-point scalar or vector
// operands.
func (block *Block) NewFCmp(pred enum.FPred, x, y value.Value) *InstFCmp {
	inst := NewFCmp(pred, x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewPhi appends a new phi instruction to the basic block based on the given
// incoming values.
func (block *Block) NewPhi(incs ...*Incoming) *InstPhi {
	inst := NewPhi(incs...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSelect appends a new select instruction to the basic block based on the
// given selection condition and operands.
func (block *Block) NewSelect(cond, x, y value.Value) *InstSelect {
	inst := NewSelect(cond, x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCall appends a new call instruction to the basic block based on the given
// callee and function arguments.
//
// TODO: specify the set of underlying types of callee.
func (block *Block) NewCall(callee value.Value, args ...value.Value) *InstCall {
	inst := NewCall(callee, args...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewVAArg appends a new va_arg instruction to the basic block based on the
// given variable argument list and argument type.
func (block *Block) NewVAArg(vaList value.Value, argType types.Type) *InstVAArg {
	inst := NewVAArg(vaList, argType)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewLandingPad appends a new landingpad instruction to the basic block based
// on the given result type and filter/catch clauses.
func (block *Block) NewLandingPad(resultType types.Type, clauses ...*Clause) *InstLandingPad {
	inst := NewLandingPad(resultType, clauses...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCatchPad appends a new catchpad instruction to the basic block based on
// the given exception scope and exception arguments.
func (block *Block) NewCatchPad(scope *TermCatchSwitch, args ...value.Value) *InstCatchPad {
	inst := NewCatchPad(scope, args...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCleanupPad appends a new cleanuppad instruction to the basic block based
// on the given exception scope and exception arguments.
func (block *Block) NewCleanupPad(scope ExceptionScope, args ...value.Value) *InstCleanupPad {
	inst := NewCleanupPad(scope, args...)
	block.Insts = append(block.Insts, inst)
	return inst
}
