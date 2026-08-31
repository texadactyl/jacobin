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

func TestLoad_Net_Http_HttpClient_RegistersSignatures(t *testing.T) {
	saved := ghelpers.MethodSignatures
	defer func() { ghelpers.MethodSignatures = saved }()

	ghelpers.MethodSignatures = make(map[string]ghelpers.GMeth)
	Load_Net_Http_HttpClient()

	expectedMethods := []string{
		// java.net.http.HttpClient
		"java/net/http/HttpClient.<clinit>()V",
		"java/net/http/HttpClient.<init>()V",
		"java/net/http/HttpClient.authenticator()Ljava/util/Optional;",
		"java/net/http/HttpClient.awaitTermination(Ljava/time/Duration;)Z",
		"java/net/http/HttpClient.close()V",
		"java/net/http/HttpClient.connectTimeout()Ljava/util/Optional;",
		"java/net/http/HttpClient.cookieHandler()Ljava/util/Optional;",
		"java/net/http/HttpClient.executor()Ljava/util/Optional;",
		"java/net/http/HttpClient.followRedirects()Ljava/net/http/HttpClient$Redirect;",
		"java/net/http/HttpClient.isTerminated()Z",
		"java/net/http/HttpClient.newBuilder()Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient.newHttpClient()Ljava/net/http/HttpClient;",
		"java/net/http/HttpClient.newWebSocketBuilder()Ljava/net/http/WebSocket$Builder;",
		"java/net/http/HttpClient.proxy()Ljava/util/Optional;",
		"java/net/http/HttpClient.send(Ljava/net/http/HttpRequest;Ljava/net/http/HttpResponse$BodyHandler;)Ljava/net/http/HttpResponse;",
		"java/net/http/HttpClient.sendAsync(Ljava/net/http/HttpRequest;Ljava/net/http/HttpResponse$BodyHandler;)Ljava/util/concurrent/CompletableFuture;",
		"java/net/http/HttpClient.sendAsync(Ljava/net/http/HttpRequest;Ljava/net/http/HttpResponse$BodyHandler;Ljava/net/http/HttpResponse$PushPromiseHandler;)Ljava/util/concurrent/CompletableFuture;",
		"java/net/http/HttpClient.shutdown()V",
		"java/net/http/HttpClient.shutdownNow()V",
		"java/net/http/HttpClient.sslContext()Ljavax/net/ssl/SSLContext;",
		"java/net/http/HttpClient.sslParameters()Ljavax/net/ssl/SSLParameters;",
		"java/net/http/HttpClient.version()Ljava/net/http/HttpClient$Version;",

		// java.net.http.HttpClient$Builder
		"java/net/http/HttpClient$Builder.<clinit>()V",
		"java/net/http/HttpClient$Builder.authenticator(Ljava/net/Authenticator;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.build()Ljava/net/http/HttpClient;",
		"java/net/http/HttpClient$Builder.connectTimeout(Ljava/time/Duration;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.cookieHandler(Ljava/net/CookieHandler;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.executor(Ljava/util/concurrent/Executor;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.followRedirects(Ljava/net/http/HttpClient$Redirect;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.localAddress(Ljava/net/InetAddress;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.priority(I)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.proxy(Ljava/net/ProxySelector;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.sslContext(Ljavax/net/ssl/SSLContext;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.sslParameters(Ljavax/net/ssl/SSLParameters;)Ljava/net/http/HttpClient$Builder;",
		"java/net/http/HttpClient$Builder.version(Ljava/net/http/HttpClient$Version;)Ljava/net/http/HttpClient$Builder;",

		// java.net.http.HttpClient$Redirect
		"java/net/http/HttpClient$Redirect.<clinit>()V",
		"java/net/http/HttpClient$Redirect.<init>(Ljava/lang/String;I)V",
		"java/net/http/HttpClient$Redirect.$values()[Ljava/net/http/HttpClient$Redirect;",
		"java/net/http/HttpClient$Redirect.valueOf(Ljava/lang/String;)Ljava/net/http/HttpClient$Redirect;",
		"java/net/http/HttpClient$Redirect.values()[Ljava/net/http/HttpClient$Redirect;",

		// java.net.http.HttpClient$Version
		"java/net/http/HttpClient$Version.<clinit>()V",
		"java/net/http/HttpClient$Version.<init>(Ljava/lang/String;I)V",
		"java/net/http/HttpClient$Version.$values()[Ljava/net/http/HttpClient$Version;",
		"java/net/http/HttpClient$Version.valueOf(Ljava/lang/String;)Ljava/net/http/HttpClient$Version;",
		"java/net/http/HttpClient$Version.values()[Ljava/net/http/HttpClient$Version;",
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

func TestLoad_Net_Http_HttpClient_TrapExecution(t *testing.T) {
	res := ghelpers.TrapFunction(nil)
	gerr, ok := res.(*ghelpers.GErrBlk)
	if !ok {
		t.Fatalf("expected *ghelpers.GErrBlk, got %T", res)
	}
	if gerr.ExceptionType != excNames.UnsupportedOperationException {
		t.Errorf("expected UnsupportedOperationException (%d), got %d", excNames.UnsupportedOperationException, gerr.ExceptionType)
	}
}
