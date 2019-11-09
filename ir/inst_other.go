package ir

import (
	"fmt"
	"strings"

	"github.com/umaumax/llvm/ir/enum"
	"github.com/umaumax/llvm/ir/types"
	"github.com/umaumax/llvm/ir/value"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstICmp is an LLVM IR icmp instruction.
type InstICmp struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Integer comparison predicate.
	Pred enum.IPred
	// Integer scalar or vector operands.
	X, Y value.Value // integer scalar, pointer, integer vector or pointer vector.

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type // boolean or boolean vector
	// (optional) Metadata.
	Metadata
}

// NewICmp returns a new icmp instruction based on the given integer comparison
// predicate and integer scalar or vector operands.
func NewICmp(pred enum.IPred, x, y value.Value) *InstICmp {
	inst := &InstICmp{Pred: pred, X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstICmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstICmp) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *types.IntType, *types.PointerType:
			inst.Typ = types.I1
		case *types.VectorType:
			inst.Typ = types.NewVector(xType.Len, types.I1)
		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *types.IntType, *types.PointerType or *types.VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstICmp) LLString() string {
	// 'icmp' Pred=IPred X=TypeValue ',' Y=Value Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "icmp %s %s, %s", inst.Pred, inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFCmp is an LLVM IR fcmp instruction.
type InstFCmp struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Floating-point comparison predicate.
	Pred enum.FPred
	// Floating-point scalar or vector operands.
	X, Y value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type // boolean or boolean vector
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// comparison predicate and floating-point scalar or vector operands.
func NewFCmp(pred enum.FPred, x, y value.Value) *InstFCmp {
	inst := &InstFCmp{Pred: pred, X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFCmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFCmp) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *types.FloatType:
			inst.Typ = types.I1
		case *types.VectorType:
			inst.Typ = types.NewVector(xType.Len, types.I1)
		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *types.FloatType or *types.VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstFCmp) LLString() string {
	// 'fcmp' FastMathFlags=FastMathFlag* Pred=FPred X=TypeValue ',' Y=Value
	// Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fcmp")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " %s %s, %s", inst.Pred, inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstPhi is an LLVM IR phi instruction.
type InstPhi struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Incoming values.
	Incs []*Incoming

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type // type of incoming value
	// (optional) Metadata.
	Metadata
}

// NewPhi returns a new phi instruction based on the given incoming values.
func NewPhi(incs ...*Incoming) *InstPhi {
	inst := &InstPhi{Incs: incs}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstPhi) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstPhi) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.Incs[0].X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstPhi) LLString() string {
	// 'phi' Typ=Type Incs=(Inc separator ',')+ Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "phi %s ", inst.Typ)
	for i, inc := range inst.Incs {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(inc.String())
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ___ [ Incoming value ] ______________________________________________________

// Incoming is an incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	X value.Value
	// Predecessor basic block of the incoming value.
	Pred *Block
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x value.Value, pred *Block) *Incoming {
	return &Incoming{X: x, Pred: pred}
}

// String returns the string representation of the incoming value.
func (inc *Incoming) String() string {
	// '[' X=Value ',' Pred=LocalIdent ']'
	return fmt.Sprintf("[ %s, %s ]", inc.X.Ident(), inc.Pred.Ident())
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSelect is an LLVM IR select instruction.
type InstSelect struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Selection condition.
	Cond value.Value // boolean or boolean vector
	// Operands.
	X, Y value.Value

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata
}

// NewSelect returns a new select instruction based on the given selection
// condition and operands.
func NewSelect(cond, x, y value.Value) *InstSelect {
	inst := &InstSelect{Cond: cond, X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSelect) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSelect) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstSelect) LLString() string {
	// 'select' FastMathFlags=FastMathFlag* Cond=TypeValue ',' X=TypeValue ','
	// Y=TypeValue Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("select")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " %s, %s, %s", inst.Cond, inst.X, inst.Y)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCall is an LLVM IR call instruction.
type InstCall struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Callee.
	// TODO: specify the set of underlying types of Callee.
	Callee value.Value
	// Function arguments.
	//
	// Arg has one of the following underlying types:
	//    value.Value
	//    *ir.Arg
	//    TODO: add metadata value?
	Args []value.Value

	// extra.

	// Type of result produced by the instruction, or function signature of the
	// callee (as used when callee is variadic).
	Typ types.Type
	// (optional) Tail; zero if not present.
	Tail enum.Tail
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Calling convention; zero if not present.
	CallingConv enum.CallingConv
	// (optional) Return attributes.
	ReturnAttrs []ReturnAttribute
	// (optional) Address space; zero if not present.
	AddrSpace types.AddrSpace
	// (optional) Function attributes.
	FuncAttrs []FuncAttribute
	// (optional) Operand bundles.
	OperandBundles []*OperandBundle
	// (optional) Metadata.
	Metadata
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
//
// TODO: specify the set of underlying types of callee.
func NewCall(callee value.Value, args ...value.Value) *InstCall {
	inst := &InstCall{Callee: callee, Args: args}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCall) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCall) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.Callee.Type().(*types.PointerType)
		if !ok {
			panic(fmt.Errorf("invalid callee type; expected *types.PointerType, got %T", inst.Callee.Type()))
		}
		sig, ok := t.ElemType.(*types.FuncType)
		if !ok {
			panic(fmt.Errorf("invalid callee type; expected *types.FuncType, got %T", t.ElemType))
		}
		if sig.Variadic {
			inst.Typ = sig
		} else {
			inst.Typ = sig.RetType
		}
	}
	if t, ok := inst.Typ.(*types.FuncType); ok {
		return t.RetType
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstCall) LLString() string {
	// Tailopt 'call' FastMathFlags=FastMathFlag* CallingConvopt
	// ReturnAttrs=ReturnAttribute* AddrSpaceopt Typ=Type Callee=Value '(' Args
	// ')' FuncAttrs=FuncAttribute* OperandBundles=('[' (OperandBundle separator
	// ',')+ ']')? Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	if !inst.Type().Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", inst.Ident())
	}
	if inst.Tail != enum.TailNone {
		fmt.Fprintf(buf, "%s ", inst.Tail)
	}
	buf.WriteString("call")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	if inst.CallingConv != enum.CallingConvNone {
		fmt.Fprintf(buf, " %s", callingConvString(inst.CallingConv))
	}
	for _, attr := range inst.ReturnAttrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	// (optional) Address space.
	if inst.AddrSpace != 0 {
		fmt.Fprintf(buf, " %s", inst.AddrSpace)
	}
	// Use function signature instead of return type for variadic functions.
	typ := inst.Type()
	if t, ok := inst.Typ.(*types.FuncType); ok {
		if t.Variadic {
			typ = t
		}
	}
	fmt.Fprintf(buf, " %s %s(", typ, inst.Callee.Ident())
	for i, arg := range inst.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString(")")
	for _, attr := range inst.FuncAttrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	if len(inst.OperandBundles) > 0 {
		buf.WriteString(" [ ")
		for i, operandBundle := range inst.OperandBundles {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(operandBundle.String())
		}
		buf.WriteString(" ]")
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstVAArg is an LLVM IR va_arg instruction.
type InstVAArg struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Variable argument list.
	ArgList value.Value
	// Argument type.
	ArgType types.Type

	// extra.

	// (optional) Metadata.
	Metadata
}

// NewVAArg returns a new va_arg instruction based on the given variable
// argument list and argument type.
func NewVAArg(argList value.Value, argType types.Type) *InstVAArg {
	return &InstVAArg{ArgList: argList, ArgType: argType}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstVAArg) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstVAArg) Type() types.Type {
	return inst.ArgType
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstVAArg) LLString() string {
	// 'va_arg' ArgList=TypeValue ',' ArgType=Type Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "va_arg %s, %s", inst.ArgList, inst.ArgType)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLandingPad is an LLVM IR landingpad instruction.
type InstLandingPad struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Result type.
	ResultType types.Type
	// (optional) Cleanup landing pad.
	Cleanup bool
	// Filter and catch clauses; zero or more if Cleanup is true, otherwise one
	// or more.
	Clauses []*Clause

	// extra.

	// (optional) Metadata.
	Metadata
}

// NewLandingPad returns a new landingpad instruction based on the given result
// type and filter/catch clauses.
func NewLandingPad(resultType types.Type, clauses ...*Clause) *InstLandingPad {
	return &InstLandingPad{ResultType: resultType, Clauses: clauses}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLandingPad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstLandingPad) Type() types.Type {
	return inst.ResultType
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstLandingPad) LLString() string {
	// 'landingpad' ResultType=Type Cleanupopt Clauses=Clause* Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "landingpad %s", inst.ResultType)
	if inst.Cleanup {
		buf.WriteString("\n\t\tcleanup")
	}
	for _, clause := range inst.Clauses {
		fmt.Fprintf(buf, "\n\t\t%s", clause)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ___ [ Landingpad clause ] ___________________________________________________

// Clause is a landingpad catch or filter clause.
type Clause struct {
	// Clause type (catch or filter).
	Type enum.ClauseType
	// Operand.
	X value.Value
}

// NewClause returns a new landingpad clause based on the given clause type and
// operand.
func NewClause(clauseType enum.ClauseType, x value.Value) *Clause {
	return &Clause{Type: clauseType, X: x}
}

// String returns the string representation of the landingpad clause.
func (clause *Clause) String() string {
	return fmt.Sprintf("%s %s", clause.Type, clause.X)
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCatchPad is an LLVM IR catchpad instruction.
type InstCatchPad struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Exception scope.
	Scope *TermCatchSwitch // TODO: rename to From? rename to Within?
	// Exception arguments.
	//
	// Arg has one of the following underlying types:
	//    value.Value
	//    TODO: add metadata value?
	Args []value.Value

	// extra.

	// (optional) Metadata.
	Metadata
}

// NewCatchPad returns a new catchpad instruction based on the given exception
// scope and exception arguments.
func NewCatchPad(scope *TermCatchSwitch, args ...value.Value) *InstCatchPad {
	return &InstCatchPad{Scope: scope, Args: args}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCatchPad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCatchPad) Type() types.Type {
	return types.Token
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstCatchPad) LLString() string {
	// 'catchpad' 'within' Scope=LocalIdent '[' Args=(ExceptionArg separator
	// ',')* ']' Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "catchpad within %s [", inst.Scope.Ident())
	for i, arg := range inst.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString("]")
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCleanupPad is an LLVM IR cleanuppad instruction.
type InstCleanupPad struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Exception scope.
	Scope ExceptionScope // TODO: rename to Parent? rename to From?
	// Exception arguments.
	//
	// Arg has one of the following underlying types:
	//    value.Value
	//    TODO: add metadata value?
	Args []value.Value

	// extra.

	// (optional) Metadata.
	Metadata
}

// NewCleanupPad returns a new cleanuppad instruction based on the given
// exception scope and exception arguments.
func NewCleanupPad(scope ExceptionScope, args ...value.Value) *InstCleanupPad {
	return &InstCleanupPad{Scope: scope, Args: args}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCleanupPad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCleanupPad) Type() types.Type {
	return types.Token
}

// LLString returns the LLVM syntax representation of the instruction.
func (inst *InstCleanupPad) LLString() string {
	// 'cleanuppad' 'within' Scope=ExceptionScope '[' Args=(ExceptionArg
	// separator ',')* ']' Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "cleanuppad within %s [", inst.Scope.Ident())
	for i, arg := range inst.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString("]")
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
