/*
 * Jacobin VM - A Java virtual machine
 * Copyright (c) 2024 by the Jacobin Authors. All rights reserved.
 * Licensed under Mozilla Public License 2.0 (MPL 2.0)  Consult jacobin.org.
 */

package jvm

import (
	"io"
	"jacobin/classloader"
	"jacobin/gfunction"
	"jacobin/globals"
	"jacobin/thread"
	"jacobin/trace"
	"os"
	"strings"
	"testing"
)

// This test uses the byte array corresponding to the class file from the following Java code,
// which tests the IDIV bytecode exception for division by zero (see TestHexIDIVexception.java in
// the \testdata directory):
//
//      // division by zero to throw exception
//      public final class ThrowIDIVexception {
//          public static void main(String[] args) {
//              int n = 6;
//	            int x = 0;
//	            int y = n/x;
//          }
//      }

var ThrowIDIVexceptionBytes = []byte{
	0xCA, 0xFE, 0xBA, 0xBE, 0x00, 0x00, 0x00, 0x3D, 0x00, 0x0F, 0x0A, 0x00, 0x02, 0x00, 0x03, 0x07,
	0x00, 0x04, 0x0C, 0x00, 0x05, 0x00, 0x06, 0x01, 0x00, 0x10, 0x6A, 0x61, 0x76, 0x61, 0x2F, 0x6C,
	0x61, 0x6E, 0x67, 0x2F, 0x4F, 0x62, 0x6A, 0x65, 0x63, 0x74, 0x01, 0x00, 0x06, 0x3C, 0x69, 0x6E,
	0x69, 0x74, 0x3E, 0x01, 0x00, 0x03, 0x28, 0x29, 0x56, 0x07, 0x00, 0x08, 0x01, 0x00, 0x12, 0x54,
	0x68, 0x72, 0x6F, 0x77, 0x49, 0x44, 0x49, 0x56, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6F,
	0x6E, 0x01, 0x00, 0x04, 0x43, 0x6F, 0x64, 0x65, 0x01, 0x00, 0x0F, 0x4C, 0x69, 0x6E, 0x65, 0x4E,
	0x75, 0x6D, 0x62, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6C, 0x65, 0x01, 0x00, 0x04, 0x6D, 0x61, 0x69,
	0x6E, 0x01, 0x00, 0x16, 0x28, 0x5B, 0x4C, 0x6A, 0x61, 0x76, 0x61, 0x2F, 0x6C, 0x61, 0x6E, 0x67,
	0x2F, 0x53, 0x74, 0x72, 0x69, 0x6E, 0x67, 0x3B, 0x29, 0x56, 0x01, 0x00, 0x0A, 0x53, 0x6F, 0x75,
	0x72, 0x63, 0x65, 0x46, 0x69, 0x6C, 0x65, 0x01, 0x00, 0x17, 0x54, 0x68, 0x72, 0x6F, 0x77, 0x49,
	0x44, 0x49, 0x56, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6F, 0x6E, 0x2E, 0x6A, 0x61, 0x76,
	0x61, 0x00, 0x31, 0x00, 0x07, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00,
	0x05, 0x00, 0x06, 0x00, 0x01, 0x00, 0x09, 0x00, 0x00, 0x00, 0x1D, 0x00, 0x01, 0x00, 0x01, 0x00,
	0x00, 0x00, 0x05, 0x2A, 0xB7, 0x00, 0x01, 0xB1, 0x00, 0x00, 0x00, 0x01, 0x00, 0x0A, 0x00, 0x00,
	0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x09, 0x00, 0x0B, 0x00, 0x0C, 0x00, 0x01,
	0x00, 0x09, 0x00, 0x00, 0x00, 0x2E, 0x00, 0x02, 0x00, 0x04, 0x00, 0x00, 0x00, 0x0A, 0x10, 0x06,
	0x3C, 0x03, 0x3D, 0x1B, 0x1C, 0x6C, 0x3E, 0xB1, 0x00, 0x00, 0x00, 0x01, 0x00, 0x0A, 0x00, 0x00,
	0x00, 0x12, 0x00, 0x04, 0x00, 0x00, 0x00, 0x04, 0x00, 0x03, 0x00, 0x05, 0x00, 0x05, 0x00, 0x06,
	0x00, 0x09, 0x00, 0x07, 0x00, 0x01, 0x00, 0x0D, 0x00, 0x00, 0x00, 0x02, 0x00, 0x0E,
}

func TestHexIDIVException(t *testing.T) {
	if testing.Short() { // don't run if running quick tests only. (Used primarily so GitHub doesn't run and bork)
		t.Skip()
	}

	var normalStderr, rerr, werr *os.File
	var normalStdout, _, wout *os.File
	var err error

	// redirect stderr & stdout to capture results from stderr
	// stderr
	normalStderr = os.Stderr
	rerr, werr, err = os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe returned an error: %s", err.Error())
		return
	}
	os.Stderr = werr

	// stdout
	normalStdout = os.Stdout
	_, wout, _ = os.Pipe()
	os.Stdout = wout

	// Initialize global, logging, classloader
	// globals.InitGlobals("testWithoutShutdown") // let test run to completion, but don't shutdown
	globals.InitGlobals("test")
	trace.Init()
	globPtr = globals.GetGlobalRef()
	globPtr.FuncInstantiateClass = InstantiateClass

	// Initialise classloader
	err = classloader.Init()
	if err != nil {
		t.Errorf("classloader.Init returned an error: %s\n", err.Error())
		return
	}

	// Load the base classes
	classloader.LoadBaseClasses()

	// Show the map size and check it for java/lang/System
	mapSize := classloader.JmodMapSize()
	if mapSize < 1 {
		t.Errorf("map size < 1 (fatal error)")
		return
	}

	// Set up MethArea for ThrowIDIVexception
	eKI := classloader.Klass{
		Status: 'I', // I = initializing the load
		Loader: "",
		Data:   nil,
	}
	classloader.MethAreaInsert("ThrowIDIVexception", &eKI)

	// Load bytes for ThrowIDIVexception
	_, _, err = classloader.ParseAndPostClass(&classloader.BootstrapCL, "ThrowIDIVexception.class", ThrowIDIVexceptionBytes)
	if err != nil {
		t.Errorf("Got error from classloader.ParseAndPostCLass: %s", error.Error(err))
		return
	}

	// Run class ThrowIDIVexception
	classloader.MTable = make(map[string]classloader.MTentry)
	gfunction.MTableLoadGFunctions(&classloader.MTable)
	mainThread := thread.CreateThread()
	mainThread.AddThreadToTable(globPtr)
	StartExec("ThrowIDIVexception", &mainThread, globals.GetGlobalRef())

	_ = werr.Close()
	_ = wout.Close()
	msgStderr, _ := io.ReadAll(rerr)

	os.Stderr = normalStderr
	os.Stdout = normalStdout

	if string(msgStderr) == "" {
		t.Errorf("Should have received error message but got none\n")
	}

	msgExpected := "IDIV or LDIV: division by zero"
	if !strings.Contains(string(msgStderr), msgExpected) {
		t.Errorf("Error expected error message to contain \"%s\", got: \"%s\"\n",
			msgExpected, string(msgStderr))
	}
}
