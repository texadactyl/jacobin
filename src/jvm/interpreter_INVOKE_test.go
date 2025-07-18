/*
 * Jacobin VM - A Java virtual machine
 * Copyright (c) 2024 by the Jacobin Authors. All rights reserved.
 * Licensed under Mozilla Public License 2.0 (MPL 2.0) Consult jacobin.org.
 */

package jvm

import (
	"io"
	"jacobin/classloader"
	"jacobin/frames"
	"jacobin/gfunction"
	"jacobin/globals"
	"jacobin/object"
	"jacobin/opcodes"
	"jacobin/stringPool"
	"jacobin/types"
	"os"
	"strings"
	"testing"
)

// This contains all the unit tests for the INVOKE family of bytecodes. They would normally
// appear in run_II-LD_test.go, but they would make that an enormous file. So, they're extracted here.

/* Restore next two tests when INVOKEINTERFACE is ported to interpreter
   ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// INVOKEINTERFACE: Invalid count field in the class file
func TestNewInvokeInterfaceInvalidCountField(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeSpecialJavaLangObject")
	}
	classloader.LoadBaseClasses() // must follow classloader.Init()

	f := newFrame(opcodes.INVOKEINTERFACE)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP
	f.Meth = append(f.Meth, 0x00) // the param count (which cannot be zero--this causes the error)
	f.Meth = append(f.Meth, 0x00)

	// create a dummy CP with 2 entries so that the CP slot index above does not cause an error.
	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0}
	f.CP = &CP

	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame
	interpret(fs)

	_ = w.Close()
	msg, _ := io.ReadAll(r)
	os.Stderr = normalStderr

	errMsg := string(msg)

	if errMsg == "" {
		t.Errorf("INVOKEINTERFACE: Should have returned an error for non-existent method, but didn't")
	} else {
		if !strings.Contains(errMsg, "Invalid values for INVOKEINTERFACE bytecode") {
			t.Errorf("INVOKEINTERFACE: Got unexpected error message: %s", errMsg)
		}
	}

}

// INVOKEINTERFACE: The CP entry does not point to an interface
func TestNewInvokeInterfaceNotPointingToInterface(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	_, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeSpecialJavaLangObject")
	}
	classloader.LoadBaseClasses() // must follow classloader.Init()

	f := newFrame(opcodes.INVOKEINTERFACE)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP
	f.Meth = append(f.Meth, 0x01)
	f.Meth = append(f.Meth, 0x00)

	// create a dummy CP with 2 entries so that the CP slot index above does not cause an error.
	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0} // expects classloader.Interface -- the error
	f.CP = &CP

	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame
	err = runFrame(fs)

	if err == nil {
		t.Errorf("INVOKEINTERFACE: Should have returned an error for non-existent method, but didn't")
	} else {
		if !strings.Contains(err.Error(), "did not point to an interface method type") {
			t.Errorf("INVOKEINTERFACE: Got unexpected error message: %s", err.Error())
		}
	}
	// restore stderr
	_ = w.Close()
	os.Stderr = normalStderr
}
*/
// INVOKESPECIAL should do nothing and report no errors
func TestNewInvokeSpecialJavaLangObject(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeSpecialJavaLangObject")
	}
	classloader.LoadBaseClasses() // must follow classloader.Init()

	f := newFrame(opcodes.INVOKESPECIAL)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP

	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0}

	CP.MethodRefs = make([]classloader.MethodRefEntry, 1)
	CP.MethodRefs[0] = classloader.MethodRefEntry{ClassIndex: 2, NameAndType: 3}

	CP.CpIndex[2] = classloader.CpEntry{Type: classloader.ClassRef, Slot: 0}
	CP.ClassRefs = make([]uint32, 4)
	CP.ClassRefs[0] = types.ObjectPoolStringIndex

	CP.CpIndex[3] = classloader.CpEntry{Type: classloader.NameAndType, Slot: 0}
	CP.NameAndTypes = make([]classloader.NameAndTypeEntry, 4)
	CP.NameAndTypes[0] = classloader.NameAndTypeEntry{
		NameIndex: 4,
		DescIndex: 5,
	}
	CP.CpIndex[4] = classloader.CpEntry{Type: classloader.UTF8, Slot: 0} // method name
	CP.Utf8Refs = make([]string, 4)
	CP.Utf8Refs[0] = "<init>"

	CP.CpIndex[5] = classloader.CpEntry{Type: classloader.UTF8, Slot: 1} // method name
	CP.Utf8Refs[1] = "()V"

	f.CP = &CP
	classloader.ResolveCPmethRefs(&CP)
	classname := "java/lang/Object"
	push(&f, object.MakeEmptyObjectWithClassName(&classname))
	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame

	interpret(fs)

	_ = w.Close()
	msg, _ := io.ReadAll(r)
	os.Stderr = normalStderr

	errMsg := string(msg)

	if errMsg != "" {
		t.Errorf("INVOKESPECIAL: Got unexpected error: %s", errMsg)
	}

	if f.TOS != 0 {
		t.Errorf("INVOKESPECIAL: Expected TOS after return to be 0, got %d", f.TOS)
	}
}

// INVOKESPECIAL: verify that a call to a gmethod works correctly (passing in nothing, getting a link back)
func TestNewInvokeSpecialGmethodNoParams(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeSpecialGmethodNoParams")
	}

	gfunction.CheckTestGfunctionsLoaded()

	f := newFrame(opcodes.INVOKESPECIAL)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP

	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0}

	CP.MethodRefs = make([]classloader.MethodRefEntry, 1)
	CP.MethodRefs[0] = classloader.MethodRefEntry{ClassIndex: 2, NameAndType: 3}

	CP.CpIndex[2] = classloader.CpEntry{Type: classloader.ClassRef, Slot: 0}
	CP.ClassRefs = make([]uint32, 4)
	classname := "jacobin/test/Object"
	CP.ClassRefs[0] = stringPool.GetStringIndex(&classname)

	CP.CpIndex[3] = classloader.CpEntry{Type: classloader.NameAndType, Slot: 0}
	CP.NameAndTypes = make([]classloader.NameAndTypeEntry, 4)
	CP.NameAndTypes[0] = classloader.NameAndTypeEntry{
		NameIndex: 4,
		DescIndex: 5,
	}
	CP.CpIndex[4] = classloader.CpEntry{Type: classloader.UTF8, Slot: 0} // method name
	CP.Utf8Refs = make([]string, 4)
	CP.Utf8Refs[0] = "test"

	CP.CpIndex[5] = classloader.CpEntry{Type: classloader.UTF8, Slot: 1} // method name
	CP.Utf8Refs[1] = "()Ljava/lang/Object;"

	f.CP = &CP
	classloader.ResolveCPmethRefs(&CP)
	obj := object.MakeEmptyObject()
	push(&f, obj) // INVOKESPECIAL expects a pointer to an object on the op stack

	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame
	interpret(fs)

	_ = w.Close()
	msg, _ := io.ReadAll(r)
	os.Stderr = normalStderr

	errMsg := string(msg)

	if errMsg != "" {
		t.Errorf("INVOKESPECIAL: Got unexpected error: %s", errMsg)
	}

	if f.TOS != 0 { // it's 0 b/c the gfunction returns a value, that is pushed onto the op stack
		t.Errorf("Expecting TOS to be 0, got %d", f.TOS)
	}
}

// INVOKESPECIAL: verify call to a gmethod works correctly and pushes the returned D twice
func TestNewInvokeSpecialGmethodNoParamsReturnsD(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeSpecialGmethodReturnsD")
	}

	gfunction.CheckTestGfunctionsLoaded()

	f := newFrame(opcodes.INVOKESPECIAL)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP

	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0}

	CP.MethodRefs = make([]classloader.MethodRefEntry, 1)
	CP.MethodRefs[0] = classloader.MethodRefEntry{ClassIndex: 2, NameAndType: 3}

	CP.CpIndex[2] = classloader.CpEntry{Type: classloader.ClassRef, Slot: 0}
	CP.ClassRefs = make([]uint32, 4)
	classname := "jacobin/test/Object"
	CP.ClassRefs[0] = stringPool.GetStringIndex(&classname)

	CP.CpIndex[3] = classloader.CpEntry{Type: classloader.NameAndType, Slot: 0}
	CP.NameAndTypes = make([]classloader.NameAndTypeEntry, 4)
	CP.NameAndTypes[0] = classloader.NameAndTypeEntry{
		NameIndex: 4,
		DescIndex: 5,
	}
	CP.CpIndex[4] = classloader.CpEntry{Type: classloader.UTF8, Slot: 0} // method name
	CP.Utf8Refs = make([]string, 4)
	CP.Utf8Refs[0] = "test"

	CP.CpIndex[5] = classloader.CpEntry{Type: classloader.UTF8, Slot: 1} // method name
	CP.Utf8Refs[1] = "()D"

	f.CP = &CP
	classloader.ResolveCPmethRefs(&CP)
	obj := object.MakeEmptyObject()
	push(&f, obj) // INVOKESPECIAL expects a pointer to an object on the op stack

	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame
	interpret(fs)

	_ = w.Close()
	msg, _ := io.ReadAll(r)
	os.Stderr = normalStderr

	errMsg := string(msg)

	if errMsg != "" {
		t.Errorf("INVOKESPECIAL: Got unexpected error: %s", errMsg)
	}

	if f.TOS != 0 {
		t.Errorf("Expecting TOS to be 0, got %d", f.TOS)
	}
}

// INVOKESPECIAL: Test proper operation of a method that reports an error
func TestNewInvokeSpecialGmethodErrorReturn(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeSpecialGmethodErrorReturn")
	}

	gfunction.CheckTestGfunctionsLoaded()

	f := newFrame(opcodes.INVOKESPECIAL)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP

	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0}

	CP.MethodRefs = make([]classloader.MethodRefEntry, 1)
	CP.MethodRefs[0] = classloader.MethodRefEntry{ClassIndex: 2, NameAndType: 3}

	CP.CpIndex[2] = classloader.CpEntry{Type: classloader.ClassRef, Slot: 0}
	CP.ClassRefs = make([]uint32, 4)
	classname := "jacobin/test/Object"
	CP.ClassRefs[0] = stringPool.GetStringIndex(&classname)

	CP.CpIndex[3] = classloader.CpEntry{Type: classloader.NameAndType, Slot: 0}
	CP.NameAndTypes = make([]classloader.NameAndTypeEntry, 4)
	CP.NameAndTypes[0] = classloader.NameAndTypeEntry{
		NameIndex: 4,
		DescIndex: 5,
	}
	CP.CpIndex[4] = classloader.CpEntry{Type: classloader.UTF8, Slot: 0} // method name
	CP.Utf8Refs = make([]string, 4)
	CP.Utf8Refs[0] = "test"

	CP.CpIndex[5] = classloader.CpEntry{Type: classloader.UTF8, Slot: 1} // method name
	CP.Utf8Refs[1] = "(D)E"

	f.CP = &CP
	classloader.ResolveCPmethRefs(&CP)
	obj := object.MakeEmptyObject()
	push(&f, obj)        // INVOKESPECIAL expects a pointer to an object on the op stack
	push(&f, int64(999)) // push the one param

	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame
	interpret(fs)

	_ = w.Close()
	msg, _ := io.ReadAll(r)
	os.Stderr = normalStderr

	errMsg := string(msg)

	if errMsg == "" {
		t.Errorf("INVOKESPECIAL: Expected an error returned, got none")
	} else {
		if !strings.Contains(errMsg, "intended return of test error") {
			t.Errorf("INVOKESPECIAL: Expected error message re 'intended return of test error', got: %s", errMsg)
		}
	}
}

// INVOKESTATIC: verify that a call to a gmethod works correctly (passing in nothing, getting a link back)
func TestNewInvokeStaticGmethodNoParams(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeStaticGmethodNoParams")
	}

	gfunction.CheckTestGfunctionsLoaded()

	f := newFrame(opcodes.INVOKESTATIC)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP

	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0}

	CP.MethodRefs = make([]classloader.MethodRefEntry, 1)
	CP.MethodRefs[0] = classloader.MethodRefEntry{ClassIndex: 2, NameAndType: 3}

	CP.CpIndex[2] = classloader.CpEntry{Type: classloader.ClassRef, Slot: 0}
	CP.ClassRefs = make([]uint32, 4)
	classname := "jacobin/test/Object"
	CP.ClassRefs[0] = stringPool.GetStringIndex(&classname)

	CP.CpIndex[3] = classloader.CpEntry{Type: classloader.NameAndType, Slot: 0}
	CP.NameAndTypes = make([]classloader.NameAndTypeEntry, 4)
	CP.NameAndTypes[0] = classloader.NameAndTypeEntry{
		NameIndex: 4,
		DescIndex: 5,
	}
	CP.CpIndex[4] = classloader.CpEntry{Type: classloader.UTF8, Slot: 0} // method name
	CP.Utf8Refs = make([]string, 4)
	CP.Utf8Refs[0] = "test"

	CP.CpIndex[5] = classloader.CpEntry{Type: classloader.UTF8, Slot: 1} // method name
	CP.Utf8Refs[1] = "()Ljava/lang/Object;"

	f.CP = &CP
	classloader.ResolveCPmethRefs(&CP)

	// INVOKESTATIC needs a parsed/loaded object in the MethArea to function
	clData := classloader.ClData{
		Name:      "jacobin/test/Object",
		NameIndex: CP.ClassRefs[0],
		// Superclass:      "java/lang/Object",
		SuperclassIndex: types.ObjectPoolStringIndex,
		Module:          "",
		Pkg:             "",
		Interfaces:      nil,
		Fields:          nil,
		MethodTable:     nil,
		// Methods:         nil,
		Attributes: nil,
		SourceFile: "",
		Bootstraps: nil,
		CP:         classloader.CPool{},
		Access:     classloader.AccessFlags{},
		ClInit:     types.ClInitRun,
	}
	k := classloader.Klass{
		Status: 'X',
		Loader: "boostrap",
		Data:   &clData,
	}

	classloader.MethAreaInsert("jacobin/test/Object", &k)

	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame
	interpret(fs)

	_ = w.Close()
	msg, _ := io.ReadAll(r)
	os.Stderr = normalStderr

	errMsg := string(msg)

	if errMsg != "" {
		t.Errorf("INVOKESTATIC: Got unexpected error: %s", errMsg)
	}

	if f.TOS != 0 {
		t.Errorf("INVOKESTATIC: Expecting TOS to be 0, got %d", f.TOS)
	}
}

// INVOKESTATIC: verify that a call to a gmethod works correctly (passing in nothing, getting a link back)
func TestNewInvokeStaticGmethodErrorReturn(t *testing.T) {
	globals.InitGlobals("test")

	// redirect stderr so as not to pollute the test output with the expected error message
	normalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeStaticGmethodNoParams")
	}

	gfunction.CheckTestGfunctionsLoaded()

	f := newFrame(opcodes.INVOKESTATIC)
	f.Meth = append(f.Meth, 0x00)
	f.Meth = append(f.Meth, 0x01) // Go to slot 0x0001 in the CP

	CP := classloader.CPool{}
	CP.CpIndex = make([]classloader.CpEntry, 10)
	CP.CpIndex[0] = classloader.CpEntry{Type: 0, Slot: 0}
	CP.CpIndex[1] = classloader.CpEntry{Type: classloader.MethodRef, Slot: 0}

	CP.MethodRefs = make([]classloader.MethodRefEntry, 1)
	CP.MethodRefs[0] = classloader.MethodRefEntry{ClassIndex: 2, NameAndType: 3}

	CP.CpIndex[2] = classloader.CpEntry{Type: classloader.ClassRef, Slot: 0}
	CP.ClassRefs = make([]uint32, 4)
	classname := "jacobin/test/Object"
	CP.ClassRefs[0] = stringPool.GetStringIndex(&classname)

	CP.CpIndex[3] = classloader.CpEntry{Type: classloader.NameAndType, Slot: 0}
	CP.NameAndTypes = make([]classloader.NameAndTypeEntry, 4)
	CP.NameAndTypes[0] = classloader.NameAndTypeEntry{
		NameIndex: 4,
		DescIndex: 5,
	}
	CP.CpIndex[4] = classloader.CpEntry{Type: classloader.UTF8, Slot: 0} // method name
	CP.Utf8Refs = make([]string, 4)
	CP.Utf8Refs[0] = "test"

	CP.CpIndex[5] = classloader.CpEntry{Type: classloader.UTF8, Slot: 1} // method name
	CP.Utf8Refs[1] = "(D)E"

	f.CP = &CP
	classloader.ResolveCPmethRefs(&CP)
	
	push(&f, int64(999)) // push the one param

	// INVOKESTATIC needs a parsed/loaded object in the MethArea to function
	clData := classloader.ClData{
		Name:      "jacobin/test/Object",
		NameIndex: CP.ClassRefs[0],
		// Superclass:      "java/lang/Object",
		SuperclassIndex: types.ObjectPoolStringIndex,
		Module:          "",
		Pkg:             "",
		Interfaces:      nil,
		Fields:          nil,
		MethodTable:     nil,
		Attributes:      nil,
		SourceFile:      "",
		Bootstraps:      nil,
		CP:              classloader.CPool{},
		Access:          classloader.AccessFlags{},
		ClInit:          types.ClInitRun,
	}
	k := classloader.Klass{
		Status: 'X',
		Loader: "boostrap",
		Data:   &clData,
	}

	classloader.MethAreaInsert("jacobin/test/Object", &k)

	fs := frames.CreateFrameStack()
	fs.PushFront(&f) // push the new frame
	interpret(fs)

	_ = w.Close()
	msg, _ := io.ReadAll(r)
	os.Stderr = normalStderr

	errMsg := string(msg)

	if errMsg == "" {
		t.Errorf("INVOKESTATIC: Expected an error returned, got none")
	} else {
		if !strings.Contains(errMsg, "intended return of test error") {
			t.Errorf("INVOKESTATIC: Expected error message re 'intended return of test error', got: %s", errMsg)
		}
	}
}
