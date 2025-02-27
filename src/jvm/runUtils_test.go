/*
 * Jacobin VM - A Java virtual machine
 * Copyright (c) 2024 by the Jacobin Authors. All rights reserved.
 * Licensed under Mozilla Public License 2.0 (MPL 2.0)  Consult jacobin.org.
 */

package jvm

import (
	"jacobin/classloader"
	"jacobin/globals"
	"jacobin/object"
	"jacobin/stringPool"
	"jacobin/trace"
	"testing"
)

// tests for runUtils.go. Note that most functions are tested inside the tests for run.go,
// but several benefit from standalone testing. Those are tested here

func TestConvertBoolByteToInt64(t *testing.T) {
	var bite = byte(0x01)
	res := convertInterfaceToInt64(bite)
	if res != 1 {
		t.Errorf("convertBoolByteToInt64(byte), expected = 1, got %d", res)
	}

	yesNo := true
	if convertInterfaceToInt64(yesNo) != 1 {
		t.Errorf("convertBoolByteToInt64(bool) != 1 (true), got %d", res)
	}
}

// test conversion of valid interfaces representing numeric values to int64
func TestConvertRemainingUntestedTypesToInt64(t *testing.T) {
	globals.InitGlobals("test")

	i8 := int8(42)
	val := convertInterfaceToInt64(i8)
	if val != 42 {
		t.Errorf("convertInterfaceToInt64(int8), expected = 42, got %d", val)
	}

	u8 := int8(42)
	val = convertInterfaceToInt64(u8)
	if val != 42 {
		t.Errorf("convertInterfaceToInt64(int8), expected = 42, got %d", val)
	}

	i16 := int16(-42)
	val = convertInterfaceToInt64(i16)
	if val != -42 {
		t.Errorf("convertInterfaceToInt64(int16), expected = -42, got %d", val)
	}

	u16 := uint16(142)
	val = convertInterfaceToInt64(u16)
	if val != 142 {
		t.Errorf("convertInterfaceToInt64(uint16), expected = 142, got %d", val)
	}

	i := int(-1042)
	val = convertInterfaceToInt64(i)
	if val != -1042 {
		t.Errorf("convertInterfaceToInt6(int), expected = -1042, got %d", val)
	}

	i32 := int32(-104232)
	val = convertInterfaceToInt64(i32)
	if val != -104232 {
		t.Errorf("convertInterfaceToInt64(int32), expected = -104232, got %d", val)
	}

	u32 := uint32(104232)
	val = convertInterfaceToInt64(u32)
	if val != 104232 {
		t.Errorf("convertInterfaceToInt64(uint32), expected = 104232, got %d", val)
	}
}

// test conversion of invalid type to int64
func TestConvertInvalidTypeToInt64(t *testing.T) {
	globals.InitGlobals("test")
	val := convertInterfaceToInt64(nil)
	if val != 0 {
		t.Errorf("convertInterfaceToInt64, expected = 0, got %d", val)
	}
}

// convert to uint64
func TestConvertFloatToInt64RoundDown(t *testing.T) {
	f := float64(5432.10)
	val := convertInterfaceToUint64(f)
	if val != 5432 {
		t.Errorf("convertFloatToInt64(float64), expected = 5432, got %d", val)
	}

}

func TestConvertFloatToInt64RoundUp(t *testing.T) {
	f := float64(5432.501)
	val := convertInterfaceToUint64(f)
	if val != 5433 {
		t.Errorf("convertFloatToInt64(float64), expected = 5433, got %d", val)
	}
}

// golang bytes are unsigned 8-bit fields. However, when a byte is part of a
// larger number (i.e., a 32-bit field) the most significant bit can indeed
// represent a sign. This test makes sure we convert such a data byte to a
// negative number.
func TestDataByteToInt64(t *testing.T) {
	b := byte(0xA0)
	val := byteToInt64(b)
	if !(val < 0) {
		t.Errorf("dataByteToInt64(byte), expected value < 0, got %d", val)
	}
}

func TestIfClassAisAsubclassOfBool(t *testing.T) {
	globals.InitGlobals("test")
	trace.Init()

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestInvokeSpecialJavaLangObject")
	}
	classloader.LoadBaseClasses() // must follow classloader.Init()
	classAname := "java/lang/ClassNotFoundException"
	classA := stringPool.GetStringIndex(&classAname)

	classBname := "java/lang/Throwable"
	classB := stringPool.GetStringIndex(&classBname)

	isIt := isClassAaSublclassOfB(classA, classB)
	if !isIt {
		t.Errorf("%s is a subclass of %s, but result said not",
			classAname, classBname)
	}
}

// check that a class is not a subclass of itself
func TestIfClassAisAsubclassOfItaelf(t *testing.T) {

}
func TestIfClassAisAsubclassOfBoolInvalid(t *testing.T) {
	globals.InitGlobals("test")

	isIt := isClassAaSublclassOfB(127, 127)
	if !isIt {
		t.Errorf("Expecting identical classes to return true, but returned false")
	}
}

// check that if an array is cast to an object, only java/lang/Object works.
func TestCheckCastArray1(t *testing.T) {
	globals.InitGlobals("test")
	trace.Init()

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestCheckCastArray1")
	}
	classloader.LoadBaseClasses()

	array := object.Make1DimArray(object.INT, 10)

	ret := checkcastArray(array, "java/lang/Object")
	if !ret {
		t.Errorf("checkcastArray(array, \"java/lang/Object\") shoud return true, got false")
	}

	ret = checkcastArray(array, "java/lang/Array")
	if ret {
		t.Errorf("checkcastArray(array, \"java/lang/Object\") shoud return false, got true")
	}
}

// check that two identical arrays come back as castable to each other
func TestCheckCastArray2(t *testing.T) {
	globals.InitGlobals("test")
	trace.Init()

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestCheckCastArray2")
	}
	classloader.LoadBaseClasses()

	array1 := object.Make1DimArray(object.INT, 10)
	array2 := object.Make1DimArray(object.INT, 10)

	ret := checkcastArray(array1, *(stringPool.GetStringPointer(array2.KlassName)))
	if !ret {
		t.Errorf("checkcastArray of two identical arrays should return true, got false")
	}
}

// check that two reference arrays are castable if one is a subclass of the other
func TestCheckCastArray3(t *testing.T) {
	globals.InitGlobals("test")
	trace.Init()

	// Initialize classloaders and method area
	err := classloader.Init()
	if err != nil {
		t.Errorf("Failure to load classes in TestCheckCastArray3")
	}
	classloader.LoadBaseClasses()

	object := object.MakeEmptyObject()
	objectKlassName := "[Ljava/lang/NullPointerException;"
	object.KlassName = stringPool.GetStringIndex(&objectKlassName)

	ret := checkcastArray(object, "[java/lang/Throwable")
	if !ret {
		t.Errorf("checkcastArray of a subclass array should return true, got false")
	}
}
