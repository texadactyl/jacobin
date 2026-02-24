/*
 * Jacobin VM - A Java virtual machine
 * Copyright (c) 2026 by  the Jacobin Authors. All rights reserved.
 * Licensed under Mozilla Public License 2.0 (MPL 2.0)  Consult jacobin.org.
 */

package javaLang

import (
	"testing"
)

func TestParseDescriptorToClasses_Invalid(t *testing.T) {
	// Test invalid descriptors
	invalidDescriptors := []string{
		"",
		"()",                   // Missing return type
		"(I",                   // Missing closing paren
		"I)V",                  // Missing opening paren
		"(Ljava/lang/String)V", // Missing semicolon
	}

	for _, desc := range invalidDescriptors {
		_, _, err := parseDescriptorToClasses(desc)
		if err == nil {
			t.Errorf("Expected error for invalid descriptor: %s", desc)
		}
	}
}
