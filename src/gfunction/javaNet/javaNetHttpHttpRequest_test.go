/*
 * Jacobin VM - A Java virtual machine
 * Copyright (c) 2026 by the Jacobin authors. Consult jacobin.org.
 * Licensed under Mozilla Public License 2.0 (MPL 2.0) All rights reserved.
 */

package javaNet

import (
	"jacobin/src/excNames"
	"jacobin/src/gfunction/ghelpers"
	"testing"
)

func TestLoad_Net_Http_HttpRequest_RegistersSignatures(t *testing.T) {
	saved := ghelpers.MethodSignatures
	defer func() { ghelpers.MethodSignatures = saved }()

	ghelpers.MethodSignatures = make(map[string]ghelpers.GMeth)
	Load_Net_Http_HttpRequest()

	expectedMethods := []string{
		// java.net.http.HttpRequest
		"java/net/http/HttpRequest.<clinit>()V",
		"java/net/http/HttpRequest.<init>()V",
		"java/net/http/HttpRequest.bodyPublisher()Ljava/util/Optional;",
		"java/net/http/HttpRequest.equals(Ljava/lang/Object;)Z",
		"java/net/http/HttpRequest.expectContinue()Z",
		"java/net/http/HttpRequest.hashCode()I",
		"java/net/http/HttpRequest.headers()Ljava/net/http/HttpHeaders;",
		"java/net/http/HttpRequest.method()Ljava/lang/String;",
		"java/net/http/HttpRequest.newBuilder()Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest.newBuilder(Ljava/net/URI;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest.newBuilder(Ljava/net/http/HttpRequest;Ljava/util/function/BiPredicate;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest.timeout()Ljava/util/Optional;",
		"java/net/http/HttpRequest.uri()Ljava/net/URI;",
		"java/net/http/HttpRequest.version()Ljava/util/Optional;",

		// java.net.http.HttpRequest$Builder
		"java/net/http/HttpRequest$Builder.<clinit>()V",
		"java/net/http/HttpRequest$Builder.DELETE()Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.GET()Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.HEAD()Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.POST(Ljava/net/http/HttpRequest$BodyPublisher;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.PUT(Ljava/net/http/HttpRequest$BodyPublisher;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.build()Ljava/net/http/HttpRequest;",
		"java/net/http/HttpRequest$Builder.copy()Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.expectContinue(Z)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.header(Ljava/lang/String;Ljava/lang/String;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.headers([Ljava/lang/String;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.method(Ljava/lang/String;Ljava/net/http/HttpRequest$BodyPublisher;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.setHeader(Ljava/lang/String;Ljava/lang/String;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.timeout(Ljava/time/Duration;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.uri(Ljava/net/URI;)Ljava/net/http/HttpRequest$Builder;",
		"java/net/http/HttpRequest$Builder.version(Ljava/net/http/HttpClient$Version;)Ljava/net/http/HttpRequest$Builder;",

		// java.net.http.HttpRequest$BodyPublisher
		"java/net/http/HttpRequest$BodyPublisher.<clinit>()V",
		"java/net/http/HttpRequest$BodyPublisher.contentLength()J",

		// java.net.http.HttpRequest$BodyPublishers
		"java/net/http/HttpRequest$BodyPublishers.<clinit>()V",
		"java/net/http/HttpRequest$BodyPublishers.<init>()V",
		"java/net/http/HttpRequest$BodyPublishers.concat([Ljava/net/http/HttpRequest$BodyPublisher;)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.fromPublisher(Ljava/util/concurrent/Flow$Publisher;)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.fromPublisher(Ljava/util/concurrent/Flow$Publisher;J)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.noBody()Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.ofByteArray([B)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.ofByteArray([BII)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.ofByteArrays(Ljava/lang/Iterable;)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.ofFile(Ljava/nio/file/Path;)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.ofInputStream(Ljava/util/function/Supplier;)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.ofString(Ljava/lang/String;)Ljava/net/http/HttpRequest$BodyPublisher;",
		"java/net/http/HttpRequest$BodyPublishers.ofString(Ljava/lang/String;Ljava/nio/charset/Charset;)Ljava/net/http/HttpRequest$BodyPublisher;",
	}

	for _, sig := range expectedMethods {
		meth, ok := ghelpers.MethodSignatures[sig]
		if !ok {
			t.Errorf("missing signature: %s", sig)
			continue
		}
		if meth.GFunction == nil {
			t.Errorf("nil GFunction for signature: %s", sig)
		}
	}

	if len(ghelpers.MethodSignatures) != len(expectedMethods) {
		t.Errorf("expected %d signatures, got %d", len(expectedMethods), len(ghelpers.MethodSignatures))
	}
}

func TestLoad_Net_Http_HttpRequest_TrapExecution(t *testing.T) {
	res := ghelpers.TrapFunction(nil)
	gerr, ok := res.(*ghelpers.GErrBlk)
	if !ok {
		t.Fatalf("expected *ghelpers.GErrBlk, got %T", res)
	}
	if gerr.ExceptionType != excNames.UnsupportedOperationException {
		t.Errorf("expected UnsupportedOperationException (%d), got %d", excNames.UnsupportedOperationException, gerr.ExceptionType)
	}
}
