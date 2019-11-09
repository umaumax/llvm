package ir

import (
	"github.com/umaumax/llvm/ir/constant"
	"github.com/umaumax/llvm/ir/value"
)

// --- [ Terminators ] ---------------------------------------------------------

// ~~~ [ ret ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewRet sets the terminator of the basic block to a new ret terminator based
// on the given return value. A nil return value indicates a void return.
func (block *Block) NewRet(x value.Value) *TermRet {
	term := NewRet(x)
	block.Term = term
	return term
}

// ~~~ [ br ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewBr sets the terminator of the basic block to a new unconditional br
// terminator based on the given target basic block.
func (block *Block) NewBr(target *Block) *TermBr {
	term := NewBr(target)
	block.Term = term
	return term
}

// ~~~ [ conditional br ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCondBr sets the terminator of the basic block to a new conditional br
// terminator based on the given branching condition and conditional target
// basic blocks.
func (block *Block) NewCondBr(cond value.Value, targetTrue, targetFalse *Block) *TermCondBr {
	term := NewCondBr(cond, targetTrue, targetFalse)
	block.Term = term
	return term
}

// ~~~ [ switch ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSwitch sets the terminator of the basic block to a new switch terminator
// based on the given control variable, default target basic block and switch
// cases.
func (block *Block) NewSwitch(x value.Value, targetDefault *Block, cases ...*Case) *TermSwitch {
	term := NewSwitch(x, targetDefault, cases...)
	block.Term = term
	return term
}

// ~~~ [ indirectbr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewIndirectBr sets the terminator of the basic block to a new indirectbr
// terminator based on the given target address (derived from a blockaddress
// constant) and set of valid target basic blocks.
func (block *Block) NewIndirectBr(addr constant.Constant, validTargets ...*Block) *TermIndirectBr {
	term := NewIndirectBr(addr, validTargets...)
	block.Term = term
	return term
}

// ~~~ [ invoke ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewInvoke sets the terminator of the basic block to a new invoke terminator
// based on the given invokee, function arguments and control flow return points
// for normal and exceptional execution.
//
// TODO: specify the set of underlying types of invokee.
func (block *Block) NewInvoke(invokee value.Value, args []value.Value, normal, exception *Block) *TermInvoke {
	term := NewInvoke(invokee, args, normal, exception)
	block.Term = term
	return term
}

// ~~~ [ resume ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewResume sets the terminator of the basic block to a new resume terminator
// based on the given exception argument to propagate.
func (block *Block) NewResume(x value.Value) *TermResume {
	term := NewResume(x)
	block.Term = term
	return term
}

// ~~~ [ catchswitch ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCatchSwitch sets the terminator of the basic block to a new catchswitch
// terminator based on the given exception scope, exception handlers and unwind
// target.
func (block *Block) NewCatchSwitch(scope ExceptionScope, handlers []*Block, unwindTarget UnwindTarget) *TermCatchSwitch {
	term := NewCatchSwitch(scope, handlers, unwindTarget)
	block.Term = term
	return term
}

// ~~~ [ catchret ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCatchRet sets the terminator of the basic block to a new catchret
// terminator based on the given exit catchpad and target basic block.
func (block *Block) NewCatchRet(from *InstCatchPad, to *Block) *TermCatchRet {
	term := NewCatchRet(from, to)
	block.Term = term
	return term
}

// ~~~ [ cleanupret ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCleanupRet sets the terminator of the basic block to a new cleanupret
// terminator based on the given exit cleanuppad and unwind target.
func (block *Block) NewCleanupRet(from *InstCleanupPad, to UnwindTarget) *TermCleanupRet {
	term := NewCleanupRet(from, to)
	block.Term = term
	return term
}

// ~~~ [ unreachable ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewUnreachable sets the terminator of the basic block to a new unreachable
// terminator.
func (block *Block) NewUnreachable() *TermUnreachable {
	term := NewUnreachable()
	block.Term = term
	return term
}
