# Go Blog Posts

- [Go’s Sweet 16](https://go.dev/blog/16years), 2025-11-14, Austin Clements, for the Go team
  > Happy Birthday, Go!

- [The Green Tea Garbage Collector](https://go.dev/blog/greenteagc), 2025-10-29, Michael Knyszek and Austin Clements
  > Go 1.25 includes a new experimental garbage collector, Green Tea.

- [Flight Recorder in Go 1.25](https://go.dev/blog/flight-recorder), 2025-09-26, Carlos Amedee and Michael Knyszek
  > Go 1.25 introduces a new tool in the diagnostic toolbox, flight recording.

- [It's survey time! How has Go has been working out for you?](https://go.dev/blog/survey2025-announce), 2025-09-16, Todd Kulesza, on behalf of the Go team
  > Help shape the future of Go

- [A new experimental Go API for JSON](https://go.dev/blog/jsonv2-exp), 2025-09-09, Joe Tsai, Daniel Martí, Johan Brandhorst-Satzkorn, Roger Peppe, Chris Hines, and Damien Neil
  > Go 1.25 introduces experimental support for encoding/json/jsontext and encoding/json/v2 packages.

- [Testing Time (and other asynchronicities)](https://go.dev/blog/testing-time), 2025-08-26, Damien Neil
  > A discussion of testing asynchronous code and an exploration of the `testing/synctest` package. Based on the GopherCon Europe 2025 talk with the same title.

- [Container-aware GOMAXPROCS](https://go.dev/blog/container-aware-gomaxprocs), 2025-08-20, Michael Pratt and Carlos Amedee
  > New GOMAXPROCS defaults in Go 1.25 improve behavior in containers.

- [Go 1.25 is released](https://go.dev/blog/go1.25), 2025-08-12, Dmitri Shuralyov, on behalf of the Go team
  > Go 1.25 adds container-aware GOMAXPROCS, testing/synctest package, experimental GC, experimental encoding/json/v2, and more.

- [The FIPS 140-3 Go Cryptographic Module](https://go.dev/blog/fips140), 2025-07-15, Filippo Valsorda (Geomys), Daniel McCarney (Geomys), and Roland Shoemaker (Google)
  > Go now has a built-in, native FIPS 140-3 compliant mode.

- [Generic interfaces](https://go.dev/blog/generic-interfaces), 2025-07-07, Axel Wagner
  > Adding type parameters to interface types is surprisingly powerful

- [[ On | No ] syntactic support for error handling](https://go.dev/blog/error-syntax), 2025-06-03, Robert Griesemer
  > Go team plans around error handling support

- [Go Cryptography Security Audit](https://go.dev/blog/tob-crypto-audit), 2025-05-19, Roland Shoemaker and Filippo Valsorda
  > Go's cryptography libraries underwent an audit by Trail of Bits.

- [More predictable benchmarking with testing.B.Loop](https://go.dev/blog/testing-b-loop), 2025-04-02, Junyang Shao
  > Better benchmark looping in Go 1.24.

- [Goodbye core types - Hello Go as we know and love it!](https://go.dev/blog/coretypes), 2025-03-26, Robert Griesemer
  > Go 1.25 simplifies the language spec by removing the notion of core types

- [Traversal-resistant file APIs](https://go.dev/blog/osroot), 2025-03-12, Damien Neil
  > New file access APIs in Go 1.24.

- [From unique to cleanups and weak: new low-level tools for efficiency](https://go.dev/blog/cleanups-and-weak), 2025-03-06, Michael Knyszek
  > Weak pointers and better finalization in Go 1.24.

- [Faster Go maps with Swiss Tables](https://go.dev/blog/swisstable), 2025-02-26, Michael Pratt
  > Go 1.24 improves map performance with a brand new map implementation

- [Testing concurrent code with testing/synctest](https://go.dev/blog/synctest), 2025-02-19, Damien Neil
  > Go 1.24 contains an experimental package to aid in testing concurrent code.

- [Extensible Wasm Applications with Go](https://go.dev/blog/wasmexport), 2025-02-13, Cherry Mui
  > Go 1.24 enhances WebAssembly capabilities with function export and reactor mode

- [Go 1.24 is released!](https://go.dev/blog/go1.24), 2025-02-11, Junyang Shao, on behalf of the Go team
  > Go 1.24 brings generic type aliases, map performance improvements, FIPS 140 compliance and more.

- [Go Developer Survey 2024 H2 Results](https://go.dev/blog/survey2024-h2-results), 2024-12-20, Alice Merrick
  > What we learned from our 2024 H2 developer survey

- [Go Protobuf: The new Opaque API](https://go.dev/blog/protobuf-opaque), 2024-12-16, Michael Stapelberg
  > We are adding a new generated code API to Go Protobuf.

- [Go Turns 15](https://go.dev/blog/15years), 2024-11-11, Austin Clements, for the Go team
  > Happy 15th birthday, Go!

- [What's in an (Alias) Name?](https://go.dev/blog/alias-names), 2024-09-17, Robert Griesemer
  > A description of generic alias types, a planned feature for Go 1.24

- [Building LLM-powered applications in Go](https://go.dev/blog/llmpowered), 2024-09-12, Eli Bendersky
  > LLM-powered applications in Go using Gemini, langchaingo and Genkit

- [Share your feedback about developing with Go](https://go.dev/blog/survey2024-h2), 2024-09-09, Alice Merrick, for the Go team
  > Help shape the future of Go by sharing your thoughts via the Go Developer Survey

- [Telemetry in Go 1.23 and beyond](https://go.dev/blog/gotelemetry), 2024-09-03, Robert Findley
  > Go 1.23 includes opt-in telemetry for the Go toolchain.

- [New unique package](https://go.dev/blog/unique), 2024-08-27, Michael Knyszek
  > New package for interning in Go 1.23.

- [Range Over Function Types](https://go.dev/blog/range-functions), 2024-08-20, Ian Lance Taylor
  > A description of range over function types, a new feature in Go 1.23.

- [Go 1.23 is released](https://go.dev/blog/go1.23), 2024-08-13, Dmitri Shuralyov, on behalf of the Go team
  > Go 1.23 adds iterators, continues loop enhancements, improves compatibility, and more.

- [Secure Randomness in Go 1.22](https://go.dev/blog/chacha8rand), 2024-05-02, Russ Cox and Filippo Valsorda
  > ChaCha8Rand is a new cryptographically secure pseudorandom number generator used in Go 1.22.

- [Evolving the Go Standard Library with math/rand/v2](https://go.dev/blog/randv2), 2024-05-01, Russ Cox
  > Go 1.22 adds math/rand/v2 and charts a course for the evolution of the Go standard library.

- [Go Developer Survey 2024 H1 Results](https://go.dev/blog/survey2024-h1-results), 2024-04-09, Alice Merrick and Todd Kulesza
  > What we learned from our 2024 H1 developer survey

- [More powerful Go execution traces](https://go.dev/blog/execution-traces-2024), 2024-03-14, Michael Knyszek
  > New features and improvements to execution traces from the last year.

- [Robust generic functions on slices](https://go.dev/blog/generic-slice-functions), 2024-02-22, Valentin Deleplace
  > Avoiding memory leaks in the slices package.

- [Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements), 2024-02-13, Jonathan Amsterdam, on behalf of the Go team
  > Go 1.22's additions to patterns for HTTP routes.

- [Go 1.22 is released!](https://go.dev/blog/go1.22), 2024-02-06, Eli Bendersky, on behalf of the Go team
  > Go 1.22 enhances for loops, brings new standard library functionality and improves performance.

- [Share your feedback about developing with Go](https://go.dev/blog/survey2024-h1), 2024-01-23, Alice Merrick, for the Go team
  > Help shape the future of Go by sharing your thoughts via the Go Developer Survey

- [Finding unreachable functions with deadcode](https://go.dev/blog/deadcode), 2023-12-12, Alan Donovan
  > deadcode is a new command to help identify functions that cannot be called.

- [Go Developer Survey 2023 H2 Results](https://go.dev/blog/survey2023-h2-results), 2023-12-05, Todd Kulesza
  > What we learned from our 2023 H2 developer survey

- [Fourteen Years of Go](https://go.dev/blog/14years), 2023-11-10, Russ Cox, for the Go team
  > Happy Birthday, Go!

- [Everything You Always Wanted to Know About Type Inference - And a Little Bit More](https://go.dev/blog/type-inference), 2023-10-09, Robert Griesemer
  > A description of how type inference for Go works. Based on the GopherCon 2023 talk with the same title.

- [Deconstructing Type Parameters](https://go.dev/blog/deconstructing-type-parameters), 2023-09-26, Ian Lance Taylor
  > Why the function signatures in the slices packages are so complicated.

- [Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview), 2023-09-19, David Chase and Russ Cox
  > Go 1.21 shipped a preview of a change in Go 1.22 to make for loops less error-prone.

- [WASI support in Go](https://go.dev/blog/wasi), 2023-09-13, Johan Brandhorst-Satzkorn, Julien Fabre, Damian Gryski, Evan Phoenix, and Achille Roussel
  > Go 1.21 adds a new port targeting the WASI preview 1 syscall API

- [Scaling gopls for the growing Go ecosystem](https://go.dev/blog/gopls-scalability), 2023-09-08, Robert Findley and Alan Donovan
  > As the Go ecosystem gets bigger, gopls must get smaller

- [Profile-guided optimization in Go 1.21](https://go.dev/blog/pgo), 2023-09-05, Michael Pratt
  > Introduction to profile-guided optimization, generally available in Go 1.21.

- [Perfectly Reproducible, Verified Go Toolchains](https://go.dev/blog/rebuild), 2023-08-28, Russ Cox
  > Go 1.21 is the first perfectly reproducible Go toolchain.

- [Structured Logging with slog](https://go.dev/blog/slog), 2023-08-22, Jonathan Amsterdam
  > The Go 1.21 standard library includes a new structured logging package, log/slog.

- [Forward Compatibility and Toolchain Management in Go 1.21](https://go.dev/blog/toolchain), 2023-08-14, Russ Cox
  > Go 1.21 manages Go toolchains like any other dependency; you will never need to manually download and install a Go toolchain again.

- [Backward Compatibility, Go 1.21, and Go 2](https://go.dev/blog/compat), 2023-08-14, Russ Cox
  > Go 1.21 expands Go's commitment to backward compatibility, so that every new Go toolchain is the best possible implementation of older toolchain semantics as well.

- [Go 1.21 is released!](https://go.dev/blog/go1.21), 2023-08-08, Eli Bendersky, on behalf of the Go team
  > Go 1.21 brings language improvements, new standard library packages, PGO GA, backward and forward compatibility in the toolchain and faster builds.

- [Experimenting with project templates](https://go.dev/blog/gonew), 2023-07-31, Cameron Balahan
  > Announcing golang.org/x/tools/cmd/gonew, an experimental tool for starting new Go projects from predefined templates

- [Share your feedback about developing with Go](https://go.dev/blog/survey2023-h2), 2023-07-25, Todd Kulesza, for the Go team
  > Help shape the future of Go by sharing your thoughts via the Go Developer Survey

- [Govulncheck v1.0.0 is released!](https://go.dev/blog/govulncheck), 2023-07-13, Julie Qiu, for the Go security team
  > Version v1.0.0 of golang.org/x/vuln has been released, introducing a new API and other improvements.

- [Go 1.21 Release Candidate](https://go.dev/blog/go1.21rc), 2023-06-21, Eli Bendersky, on behalf of the Go team
  > Go 1.21 RC brings language improvements, new standard library packages, PGO GA, backward and forward compatibility in the toolchain and faster builds.

- [Go Developer Survey 2023 Q1 Results](https://go.dev/blog/survey2023-q1-results), 2023-05-11, Alice Merrick
  > An analysis of the results from the 2023 Q1 Go Developer Survey.

- [Code coverage for Go integration tests](https://go.dev/blog/integration-test-coverage), 2023-03-08, Than McIntosh
  > Code coverage for integration tests, available in Go 1.20.

- [All your comparable types](https://go.dev/blog/comparable), 2023-02-17, Robert Griesemer
  > type parameters, type sets, comparable types, constraint satisfaction

- [Profile-guided optimization preview](https://go.dev/blog/pgo-preview), 2023-02-08, Michael Pratt
  > Introduction to profile-guided optimization, available as a preview in Go 1.20.

- [Go 1.20 is released!](https://go.dev/blog/go1.20), 2023-02-01, Robert Griesemer, on behalf of the Go team
  > Go 1.20 brings PGO, faster builds, and various tool, language, and library improvements.

- [Share your feedback about developing with Go](https://go.dev/blog/survey2023-q1), 2023-01-18, Alice Merrick, for the Go team
  > Help shape the future of Go by sharing your thoughts via the Go Developer Survey

- [Thirteen Years of Go](https://go.dev/blog/13years), 2022-11-10, Russ Cox, for the Go team
  > Happy Birthday, Go!

- [Go runtime: 4 years later](https://go.dev/blog/go119runtime), 2022-09-26, Michael Knyszek
  > A check-in on the status of Go runtime development

- [Go Developer Survey 2022 Q2 Results](https://go.dev/blog/survey2022-q2-results), 2022-09-08, Todd Kulesza
  > An analysis of the results from the 2022 Q2 Go Developer Survey.

- [Vulnerability Management for Go](https://go.dev/blog/vuln), 2022-09-06, Julie Qiu, for the Go security team
  > Announcing vulnerability management for Go, to help developers learn about known vulnerabilities in their dependencies.

- [Go 1.19 is released!](https://go.dev/blog/go1.19), 2022-08-02, The Go Team
  > Go 1.19 adds richer doc comments, performance improvements, and more.

- [Share your feedback about developing with Go](https://go.dev/blog/survey2022-q2), 2022-06-01, Todd Kulesza, for the Go team
  > Help shape the future of Go by sharing your thoughts via the Go Developer Survey

- [Go Developer Survey 2021 Results](https://go.dev/blog/survey2021-results), 2022-04-19, Alice Merrick
  > An analysis of the results from the 2021 Go Developer Survey.

- [When To Use Generics](https://go.dev/blog/when-generics), 2022-04-12, Ian Lance Taylor
  > When to use generics when writing Go code, and when not to use them.

- [Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces), 2022-04-05, Beth Brown, for the Go team
  > Learn about Go workspaces and some of the workflows they enable.

- [How Go Mitigates Supply Chain Attacks](https://go.dev/blog/supply-chain), 2022-03-31, Filippo Valsorda
  > Go tooling and design help mitigate supply chain attacks at various stages.

- [An Introduction To Generics](https://go.dev/blog/intro-generics), 2022-03-22, Robert Griesemer and Ian Lance Taylor
  > An introduction to generics in Go.

- [Go 1.18 is released!](https://go.dev/blog/go1.18), 2022-03-15, The Go Team
  > Go 1.18 adds generics, native fuzzing, workspace mode, performance improvements, and more.

- [Announcing Go 1.18 Beta 2](https://go.dev/blog/go1.18beta2), 2022-01-31, Jeremy Faller and Steve Francia, for the Go team
  > Go 1.18 Beta 2 is our second preview of Go 1.18. Please try it and let us know if you find problems.

- [Two New Tutorials for 1.18](https://go.dev/blog/tutorials-go1.18), 2022-01-14, Katie Hockman, for the Go team
  > Two new tutorials have been published in preparation for the release of Go 1.18.

- [Go 1.18 Beta 1 is available, with generics](https://go.dev/blog/go1.18beta1), 2021-12-14, Russ Cox, for the Go team
  > Go 1.18 Beta 1 is our first preview of Go 1.18. Please try it and let us know if you find problems.

- [Twelve Years of Go](https://go.dev/blog/12years), 2021-11-10, Russ Cox, for the Go team
  > Happy Birthday, Go!

- [A new search experience on pkg.go.dev](https://go.dev/blog/pkgsite-search-redesign), 2021-11-09, Julie Qiu
  > Package search on pkg.go.dev has been updated, and you can now search for symbols!

- [Announcing the 2021 Go Developer Survey](https://go.dev/blog/survey2021), 2021-10-26, Alice Merrick
  > Please take the 2021 Go Developer Survey. We want to hear from you!

- [Code of Conduct Updates](https://go.dev/blog/conduct-2021), 2021-09-16, Carmen Andoh, Russ Cox, and Steve Francia
  > A small update to, and an update on enforcement of, the Go Code of Conduct

- [Automatic cipher suite ordering in crypto/tls](https://go.dev/blog/tls-cipher-suites), 2021-09-15, Filippo Valsorda
  > Go 1.17 is making TLS configuration easier and safer by automating TLS cipher suite preference ordering.

- [Tidying up the Go web experience](https://go.dev/blog/tidy-web), 2021-08-18, Russ Cox
  > Consolidating our web sites onto go.dev.

- [Go 1.17 is released](https://go.dev/blog/go1.17), 2021-08-16, Matt Pearring and Alex Rakoczy
  > Go 1.17 adds performance improvements, module optimizations, arm64 on Windows, and more.

- [The Go Collective on Stack Overflow](https://go.dev/blog/stackoverflow), 2021-06-23, Steve Francia
  > Announcing the Go Collective, a new experience for Go on Stack Overflow.

- [Fuzzing is Beta Ready](https://go.dev/blog/fuzz-beta), 2021-06-03, Katie Hockman and Jay Conrod
  > Native Go fuzzing is now ready for beta testing on tip.

- [Go Developer Survey 2020 Results](https://go.dev/blog/survey2020-results), 2021-03-09, Alice Merrick
  > An analysis of the results from the 2020 Go Developer Survey.

- [Contexts and structs](https://go.dev/blog/context-and-structs), 2021-02-24, Jean Barkhuysen, Matt T. Proud
  > 

- [New module changes in Go 1.16](https://go.dev/blog/go116-module-changes), 2021-02-18, Jay Conrod
  > Go 1.16 enables modules by default, provides a new way to install executables, and lets module authors retract published versions.

- [Go 1.16 is released](https://go.dev/blog/go1.16), 2021-02-16, Matt Pearring and Dmitri Shuralyov
  > Go 1.16 adds embedded files, Apple Silicon support, and more.

- [Gopls on by default in the VS Code Go extension](https://go.dev/blog/gopls-vscode-go), 2021-02-01, Go tools team
  > Gopls, which provides IDE features for Go to many editors, is now used by default in VS Code Go.

- [Command PATH security in Go](https://go.dev/blog/path-security), 2021-01-19, Russ Cox
  > How to decide if your programs are vulnerable to PATH problems, and what to do about it.

- [A Proposal for Adding Generics to Go](https://go.dev/blog/generics-proposal), 2021-01-12, Ian Lance Taylor
  > Generics is entering the language change proposal process

- [Go on ARM and Beyond](https://go.dev/blog/ports), 2020-12-17, Russ Cox
  > Go's support for ARM64 and other architectures

- [Redirecting godoc.org requests to pkg.go.dev](https://go.dev/blog/godoc.org-redirect), 2020-12-15, Julie Qiu
  > The plan for moving from godoc.org to pkg.go.dev.

- [Eleven Years of Go](https://go.dev/blog/11years), 2020-11-10, Russ Cox, for the Go team
  > Happy Birthday, Go!

- [Pkg.go.dev has a new look!](https://go.dev/blog/pkgsite-redesign), 2020-11-10, Julie Qiu
  > Announcing a new user experience on pkg.go.dev.

- [Announcing the 2020 Go Developer Survey](https://go.dev/blog/survey2020), 2020-10-20, Alice Merrick
  > Please take the 2020 Go Developer Survey. We want to hear from you!

- [Go 1.15 is released](https://go.dev/blog/go1.15), 2020-08-11, Alex Rakoczy
  > Go 1.15 adds a new linker, X.509 changes, runtime improvements, compiler improvements, GOPROXY improvements, and more.

- [Keeping Your Modules Compatible](https://go.dev/blog/module-compatibility), 2020-07-07, Jean Barkhuysen and Jonathan Amsterdam
  > How to keep your modules compatible with prior minor/patch versions.

- [The Next Step for Generics](https://go.dev/blog/generics-next-step), 2020-06-16, Ian Lance Taylor and Robert Griesemer
  > An updated generics design draft, and a translation tool for experimentation

- [Pkg.go.dev is open source!](https://go.dev/blog/pkgsite), 2020-06-15, Julie Qiu
  > 

- [The VS Code Go extension joins the Go project](https://go.dev/blog/vscode-go), 2020-06-09, The Go team
  > Announcement of VS Code Go’s move to the Go project.

- [Go Developer Survey 2019 Results](https://go.dev/blog/survey2019-results), 2020-04-20, Todd Kulesza
  > An analysis of the results from the 2019 Go Developer Survey.

- [Go, the Go Community, and the Pandemic](https://go.dev/blog/pandemic), 2020-03-25, Carmen Andoh, Russ Cox, and Steve Francia
  > How the Go team is approaching the pandemic, what you can expect from us, and what you can do.

- [A new Go API for Protocol Buffers](https://go.dev/blog/protobuf-apiv2), 2020-03-02, Joe Tsai, Damien Neil, and Herbie Ong
  > Announcing a major revision of the Go API for protocol buffers.

- [Go 1.14 is released](https://go.dev/blog/go1.14), 2020-02-25, Alex Rakoczy
  > Go 1.14 adds production-ready module support, faster defers, better goroutine preemption, and more.

- [Next steps for pkg.go.dev](https://go.dev/blog/pkg.go.dev-2020), 2020-01-31, Julie Qiu
  > What the Go team is planning for pkg.go.dev in 2020.

- [Proposals for Go 1.15](https://go.dev/blog/go1.15-proposals), 2020-01-28, Robert Griesemer, for the Go team
  > For Go 1.15, we propose three minor language cleanup changes.

- [Announcing the 2019 Go Developer Survey](https://go.dev/blog/survey2019), 2019-11-20, Todd Kulesza
  > Please take the 2019 Go Developer Survey. We want to hear from you!

- [Go.dev: a new hub for Go developers](https://go.dev/blog/go.dev), 2019-11-13, Steve Francia and Julie Qiu
  > Announcing go.dev, which answers: who else is using Go, what do they use it for, and how can I find useful Go packages?

- [Go Turns 10](https://go.dev/blog/10years), 2019-11-08, Russ Cox, for the Go team
  > Happy 10th birthday, Go!

- [Go Modules: v2 and Beyond](https://go.dev/blog/v2-go-modules), 2019-11-07, Jean Barkhuysen and Tyler Bui-Palsulich
  > How to release major version 2 of your module.

- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors), 2019-10-17, Damien Neil and Jonathan Amsterdam
  > How to use the new Go 1.13 error interfaces and functions.

- [Publishing Go Modules](https://go.dev/blog/publishing-go-modules), 2019-09-26, Tyler Bui-Palsulich
  > How to write and publish modules for use as dependencies.

- [Go 1.13 is released](https://go.dev/blog/go1.13), 2019-09-03, Andrew Bonventre
  > Go 1.13 adds module authentication, new number literals, error wrapping, TLS 1.3 on by default, and more.

- [Module Mirror and Checksum Database Launched](https://go.dev/blog/module-mirror-launch), 2019-08-29, Katie Hockman
  > The Go module mirror and checksum database provide faster, verified downloads of your Go dependencies.

- [Migrating to Go Modules](https://go.dev/blog/migrating-to-go-modules), 2019-08-21, Jean Barkhuysen
  > How to use Go modules to manage your program's dependencies.

- [Contributors Summit 2019](https://go.dev/blog/contributors-summit-2019), 2019-08-15, Carmen Andoh and contributors
  > Reporting from the Go Contributor Summit at GopherCon 2019.

- [Experiment, Simplify, Ship](https://go.dev/blog/experiment), 2019-08-01, Russ Cox
  > How we develop Go, a talk from GopherCon 2019.

- [Why Generics?](https://go.dev/blog/why-generics), 2019-07-31, Ian Lance Taylor
  > Why should we add generics to Go, and what might they look like?

- [Announcing The New Go Store](https://go.dev/blog/store), 2019-07-18, Cassandra Salisbury
  > Unfortunately, the Go store is offline.

- [Next steps toward Go 2](https://go.dev/blog/go2-next-steps), 2019-06-26, Robert Griesemer, for the Go team
  > What Go 2 language changes should we include in Go 1.14?

- [Go 2018 Survey Results](https://go.dev/blog/survey2018-results), 2019-03-28, Todd Kulesza, Steve Francia
  > What we learned from the December 2018 Go User Survey.

- [Debugging what you deploy in Go 1.12](https://go.dev/blog/debug-opt), 2019-03-21, David Chase
  > Go 1.12 improves support for debugging optimized binaries.

- [Using Go Modules](https://go.dev/blog/using-go-modules), 2019-03-19, Tyler Bui-Palsulich and Eno Compton
  > An introduction to the basic operations needed to get started with Go modules.

- [The New Go Developer Network](https://go.dev/blog/go-developer-network), 2019-03-14, GoBridge Leadership Team
  > Announcing the Go Developer Network, a collection of Go user groups sharing best practices.

- [What's new in the Go Cloud Development Kit](https://go.dev/blog/go-cloud2019), 2019-03-04, The Go Cloud Development Kit team at Google
  > Recent changes to the Go Cloud Development Kit (Go CDK).

- [Go 1.12 is released](https://go.dev/blog/go1.12), 2019-02-25, Andrew Bonventre
  > Go 1.12 adds opt-in TLS 1.3, improved modules, and more.

- [Go Modules in 2019](https://go.dev/blog/modules2019), 2018-12-19, Russ Cox
  > What the Go team is planning for Go modules in 2019.

- [Go 2, here we come!](https://go.dev/blog/go2-here-we-come), 2018-11-29, Robert Griesemer
  > How Go 2 proposals will be evaluated, selected, and shipped.

- [Nine years of Go](https://go.dev/blog/9years), 2018-11-10, Steve Francia
  > Happy 9th birthday, Go!

- [Participate in the 2018 Go User Survey](https://go.dev/blog/survey2018), 2018-11-08, Ran Tao, Steve Francia
  > Please take the 2018 Go User Survey. We want to hear from you!

- [Announcing App Engine’s New Go 1.11 Runtime](https://go.dev/blog/appengine-go111), 2018-10-16, Eno Compton and Tyler Bui-Palsulich
  > Google Cloud is announcing a new Go 1.11 runtime for App Engine, with fewer limits on app structure.

- [Compile-time Dependency Injection With Go Cloud's Wire](https://go.dev/blog/wire), 2018-10-09, Robert van Gent
  > How to use Wire, a dependency injection tool for Go.

- [Participate in the 2018 Go Company Questionnaire](https://go.dev/blog/survey2018-company), 2018-10-04, Ran Tao, Steve Francia
  > Please take the 2018 Go Company Questionnaire.

- [Go 2 Draft Designs](https://go.dev/blog/go2draft), 2018-08-28, 
  > Announcing the draft designs for the major Go 2 changes.

- [Go 1.11 is released](https://go.dev/blog/go1.11), 2018-08-24, Andrew Bonventre
  > Go 1.11 adds preliminary support for Go modules, WebAssembly, and more.

- [Portable Cloud Programming with Go Cloud](https://go.dev/blog/go-cloud), 2018-07-24, Eno Compton and Cassandra Salisbury
  > Announcing Go Cloud, for portable cloud programming with Go.

- [Getting to Go: The Journey of Go's Garbage Collector](https://go.dev/blog/ismmkeynote), 2018-07-12, Rick Hudson
  > A technical talk about the structure and details of the new, low-latency Go garbage collector.

- [Updating the Go Code of Conduct](https://go.dev/blog/conduct-2018), 2018-05-23, Steve Francia
  > Revising the Go Code of Conduct.

- [Go's New Brand](https://go.dev/blog/go-brand), 2018-04-26, Steve Francia
  > Go’s new look and logo (don't worry, the mascot isn’t changing!).

- [A Proposal for Package Versioning in Go](https://go.dev/blog/versioning-proposal), 2018-03-26, Russ Cox
  > Proposing official support for package versioning in Go, using Go modules.

- [Go 2017 Survey Results](https://go.dev/blog/survey2017-results), 2018-02-26, Steve Francia
  > What we learned from the December 2017 Go User Survey.

- [Go 1.10 is released](https://go.dev/blog/go1.10), 2018-02-16, Brad Fitzpatrick
  > Go 1.10 adds automatic caching of build & test results, and more.

- [Hello, 中国!](https://go.dev/blog/hello-china), 2018-01-22, Andrew Bonventre
  > The Go home page and binary downloads is now available in China, at https://golang.google.cn.

- [Participate in the 2017 Go User Survey](https://go.dev/blog/survey2017), 2017-11-16, Steve Francia
  > Please take the 2017 Go User Survey. We want to hear from you!

- [Eight years of Go](https://go.dev/blog/8years), 2017-11-10, Steve Francia
  > Happy 8th birthday, Go!

- [Community Outreach Working Group](https://go.dev/blog/cwg), 2017-09-05, Steve Francia & Cassandra Salisbury
  > Announcing the Go Community Outreach Working Group (CWG).

- [Go 1.9 is released](https://go.dev/blog/go1.9), 2017-08-24, Francesc Campoy
  > Go 1.9 adds type aliases, bit intrinsics, optimizations, and more.

- [Contribution Workshop](https://go.dev/blog/contributor-workshop), 2017-08-09, Steve Francia, Cassandra Salisbury, Matt Broberg, and Dmitri Shuralyov
  > The Go contributor workshop trained new contributors at GopherCon.

- [Contributors Summit](https://go.dev/blog/contributors-summit), 2017-08-03, Sam Whited
  > Reporting from the Go Contributor Summit at GopherCon 2017.

- [Toward Go 2](https://go.dev/blog/toward-go2), 2017-07-13, Russ Cox
  > How we will all work together toward Go 2.

- [Introducing the Developer Experience Working Group](https://go.dev/blog/developer-experience), 2017-04-10, The Developer Experience Working Group
  > Announcing the Developer eXperience Working Group (DXWG).

- [HTTP/2 Server Push](https://go.dev/blog/h2push), 2017-03-24, Jaana Burcu Dogan and Tom Bergan
  > How to use HTTP/2 server push to reduce page load times.

- [Go 2016 Survey Results](https://go.dev/blog/survey2016-results), 2017-03-06, Steve Francia, for the Go team
  > What we learned from the December 2017 Go User Survey.

- [Go 1.8 is released](https://go.dev/blog/go1.8), 2017-02-16, Chris Broadfoot
  > Go 1.8 adds faster non-x86 compiled code, sub-millisecond garbage collection pauses, HTTP/2 push, and more.

- [Participate in the 2016 Go User Survey and Company Questionnaire](https://go.dev/blog/survey2016), 2016-12-13, Steve Francia
  > Please take the 2016 Go User Survey and Company Questionnaire. We want to hear from you!

- [Go fonts](https://go.dev/blog/go-fonts), 2016-11-16, Nigel Tao, Chuck Bigelow, and Rob Pike
  > Announcing the Go font family, by Bigelow & Holmes.

- [Seven years of Go](https://go.dev/blog/7years), 2016-11-10, The Go Team
  > Happy 7th birthday, Go!

- [Introducing HTTP Tracing](https://go.dev/blog/http-tracing), 2016-10-04, Jaana Burcu Dogan
  > How to use Go 1.7's HTTP tracing to understand your client requests.

- [Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests), 2016-10-03, Marcel van Lohuizen
  > How to use Go 1.7's new subtests and sub-benchmarks.

- [Smaller Go 1.7 binaries](https://go.dev/blog/go1.7-binary-size), 2016-08-18, David Crawshaw
  > Go 1.7 includes some binary size reductions important for small devices.

- [Go 1.7 is released](https://go.dev/blog/go1.7), 2016-08-15, Chris Broadfoot
  > Go 1.7 adds faster x86 compiled code, context in the standard library, and more.

- [Go 1.6 is released](https://go.dev/blog/go1.6), 2016-02-17, Andrew Gerrand
  > Go 1.6 adds HTTP/2, template blocks, and more.

- [Language and Locale Matching in Go](https://go.dev/blog/matchlang), 2016-02-09, Marcel van Lohuizen
  > How to internationalize your web site with Go's language and locale matching.

- [Six years of Go](https://go.dev/blog/6years), 2015-11-10, Andrew Gerrand
  > Happy 6th birthday, Go!

- [Golang UK 2015](https://go.dev/blog/gouk15), 2015-10-09, Francesc Campoy
  > Reporting from GolangUK 2015, the first London Go conference.

- [Go GC: Prioritizing low latency and simplicity](https://go.dev/blog/go15gc), 2015-08-31, Richard Hudson
  > Go 1.5 is the first step toward a new low-latency future for the Go garbage collector.

- [Go 1.5 is released](https://go.dev/blog/go1.5), 2015-08-19, Andrew Gerrand
  > Go 1.5 adds a new, much faster garbage collector, more parallelism by default, go tool trace, and more.

- [GopherCon 2015 Roundup](https://go.dev/blog/gophercon2015), 2015-07-28, Andrew Gerrand
  > Reporting from GopherCon 2015.

- [Go, Open Source, Community](https://go.dev/blog/open-source), 2015-07-08, Russ Cox
  > Why is Go open source, and how can we strengthen our open-source community?

- [Qihoo 360 and Go](https://go.dev/blog/qihoo), 2015-07-06, Yang Zhou
  > How Qihoo 360 uses Go.

- [GopherChina Trip Report](https://go.dev/blog/gopherchina), 2015-07-01, Robert Griesemer
  > Reporting from GopherChina 2015, the first Go conference in China.

- [Testable Examples in Go](https://go.dev/blog/examples), 2015-05-07, Andrew Gerrand
  > How to add examples, which double as tests, to your packages.

- [Package names](https://go.dev/blog/package-names), 2015-02-04, Sameer Ajmani
  > How to name your packages.

- [Errors are values](https://go.dev/blog/errors-are-values), 2015-01-12, Rob Pike
  > Idioms and patterns for handling errors in Go.

- [GothamGo: gophers in the big apple](https://go.dev/blog/gothamgo), 2015-01-09, Francesc Campoy
  > Reporting from GothamGo 2015, the first full-day Go conference in New York City.

- [The Gopher Gala is the first worldwide Go hackathon](https://go.dev/blog/gophergala), 2015-01-07, Francesc Campoy
  > The Gopher Gala, the first global Go hackathon, will take place January 23-25, 2015.

- [Generating code](https://go.dev/blog/generate), 2014-12-22, Rob Pike
  > How to use go generate.

- [Go 1.4 is released](https://go.dev/blog/go1.4), 2014-12-10, Andrew Gerrand
  > Go 1.4 adds support for Android, go generate, optimizations, and more.

- [Half a decade with Go](https://go.dev/blog/5years), 2014-11-10, Andrew Gerrand
  > Happy 5th birthday, Go!

- [Go at Google I/O and Gopher SummerFest](https://go.dev/blog/io2014), 2014-10-06, Francesc Campoy
  > Reporting from Google I/O 2014 and the GoSF Go SummerFest.

- [Deploying Go servers with Docker](https://go.dev/blog/docker), 2014-09-26, Andrew Gerrand
  > How to use Docker's new official base images for Go.

- [Constants](https://go.dev/blog/constants), 2014-08-25, Rob Pike
  > An introduction to constants in Go.

- [Go at OSCON](https://go.dev/blog/osconreport), 2014-08-20, Francesc Campoy
  > Reporting from OSCON 2014.

- [Go Concurrency Patterns: Context](https://go.dev/blog/context), 2014-07-29, Sameer Ajmani
  > An introduction to the Go context package.

- [Go will be at OSCON 2014](https://go.dev/blog/oscon), 2014-07-15, Francesc Campoy
  > If you will be at OSCON 2014, July 20-29 in Portland, Oregon, be sure to check out these Go talks.

- [Go 1.3 is released](https://go.dev/blog/go1.3), 2014-06-18, Andrew Gerrand
  > Go 1.3 adds better performance, static analysis in godoc, and more.

- [GopherCon 2014 Wrap Up](https://go.dev/blog/gophercon), 2014-05-28, Andrew Gerrand
  > Reporting from GopherCon 2014.

- [The Go Gopher](https://go.dev/blog/gopher), 2014-03-24, Rob Pike and Andrew Gerrand
  > The backstory of the Go gopher.

- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines), 2014-03-13, Sameer Ajmani
  > How to use Go's concurrency to build data-processing pipelines.

- [Go talks at FOSDEM 2014](https://go.dev/blog/fosdem14), 2014-02-24, Andrew Gerrand
  > Reporting from the Go Devroom at FOSDEM 2014.

- [Go on App Engine: tools, tests, and concurrency](https://go.dev/blog/appengine-dec2013), 2013-12-13, Andrew Gerrand and Johan Euphrosine
  > Announcing improvements to Go on App Engine.

- [Inside the Go Playground](https://go.dev/blog/playground), 2013-12-12, Andrew Gerrand
  > How the Go playground works.

- [The cover story](https://go.dev/blog/cover), 2013-12-02, Rob Pike
  > Introducing Go 1.12's code coverage tool.

- [Go 1.2 is released](https://go.dev/blog/go1.2), 2013-12-01, Andrew Gerrand
  > Go 1.2 adds test coverage results, goroutine preemption, and more.

- [Text normalization in Go](https://go.dev/blog/normalization), 2013-11-26, Marcel van Lohuizen
  > How and why to normalize UTF-8 text in Go.

- [Four years of Go](https://go.dev/blog/4years), 2013-11-10, Andrew Gerrand
  > Happy 4th birthday, Go!

- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings), 2013-10-23, Rob Pike
  > How strings work in Go, and how to use them.

- [Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices), 2013-09-26, Rob Pike
  > How Go arrays and slices work, and how to use copy and append.

- [The first Go program](https://go.dev/blog/first-go-program), 2013-07-18, Andrew Gerrand
  > Rob Pike dug up the first Go program ever written.

- [Introducing the Go Race Detector](https://go.dev/blog/race-detector), 2013-06-26, Dmitry Vyukov and Andrew Gerrand
  > How and why to use the Go race detector to improve your programs.

- [Go and the Google Cloud Platform](https://go.dev/blog/io2013-talks-cloud), 2013-06-12, Andrew Gerrand
  > Two talks about using Go with the Google Cloud Platform, from Google I/O 2013.

- [A conversation with the Go team](https://go.dev/blog/io2013-chat), 2013-06-06, 
  > At Google I/O 2013, several members of the Go team hosted a "Fireside chat."

- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency), 2013-05-23, Andrew Gerrand
  > Watch Sameer Ajmani's talk, “Advanced Go Concurrency Patterns,” from Google I/O 2013.

- [Go 1.1 is released](https://go.dev/blog/go1.1), 2013-05-13, Andrew Gerrand
  > Go 1.1 is faster, less picky about return statements, and adds method expressions.

- [The path to Go 1](https://go.dev/blog/go1-path), 2013-03-14, Andrew Gerrand
  > Watch Rob Pike and Andrew Gerrand's talk, The Path to Go 1.

- [Two recent Go articles](https://go.dev/blog/two-recent-go-articles), 2013-03-06, Andrew Gerrand
  > Two Go articles: “Go at Google: Language Design in the Service of Software Engineering” and “Getting Started with Go, App Engine and Google+ API”

- [Get thee to a Go meetup](https://go.dev/blog/meetups), 2013-02-27, Andrew Gerrand
  > How to find or start a local group of gophers.

- [Go maps in action](https://go.dev/blog/maps), 2013-02-06, Andrew Gerrand
  > How and when to use Go maps.

- [go fmt your code](https://go.dev/blog/gofmt), 2013-01-23, Andrew Gerrand
  > How and why to format your Go code using gofmt.

- [Concurrency is not parallelism](https://go.dev/blog/waza-talk), 2013-01-16, Andrew Gerrand
  > Watch Rob Pike's talk, _Concurrency is not parallelism._

- [The App Engine SDK and workspaces (GOPATH)](https://go.dev/blog/appengine-gopath), 2013-01-09, Andrew Gerrand
  > App Engine SDK 1.7.4 adds support for GOPATH-style workspaces.

- [Two recent Go talks](https://go.dev/blog/two-recent-go-talks), 2013-01-02, Andrew Gerrand
  > Two Go talks: “Go: A Simple Programming Environment” and “Go: Code That Grows With Grace”.

- [Go turns three](https://go.dev/blog/3years), 2012-11-10, Russ Cox
  > Happy 3rd birthday, Go!

- [Go updates in App Engine 1.7.1](https://go.dev/blog/appengine-171), 2012-08-22, Andrew Gerrand
  > App Engine SDK 1.7.1 adds memcache and other functionality for Go.

- [Organizing Go code](https://go.dev/blog/organizing-go-code), 2012-08-16, Andrew Gerrand
  > How to name and package the parts of your Go program to best serve your users.

- [Gccgo in GCC 4.7.1](https://go.dev/blog/gccgo-in-gcc-471), 2012-07-11, Ian Lance Taylor
  > GCC 4.7.1 adds support for Go 1.

- [Go videos from Google I/O 2012](https://go.dev/blog/io2012-videos), 2012-07-02, Andrew Gerrand
  > Talks about Go from Google I/O 2012.

- [Go version 1 is released](https://go.dev/blog/go1), 2012-03-28, Andrew Gerrand
  > A major milestone: announcing Go 1, the first stable version of Go.

- [Getting to know the Go community](https://go.dev/blog/survey2011), 2011-12-21, Andrew Gerrand
  > Please take a Gopher Survey. We want to hear from you!

- [Building StatHat with Go](https://go.dev/blog/stathat), 2011-12-19, Patrick Crosby
  > How StatHat uses Go, and why they chose it.

- [From zero to Go: launching on the Google homepage in 24 hours](https://go.dev/blog/turkey-doodle), 2011-12-13, Reinaldo Aguiar
  > How Go helped launch the Google Doodle for Thanksgiving 2011.

- [The Go Programming Language turns two](https://go.dev/blog/2years), 2011-11-10, Andrew Gerrand
  > Happy 2nd birthday, Go!

- [Writing scalable App Engine applications](https://go.dev/blog/appengine-scalable), 2011-11-01, David Symonds
  > How to build scalable web applications using Go with Google App Engine.

- [Debugging Go programs with the GNU Debugger](https://go.dev/blog/debug-gdb), 2011-10-30, Andrew Gerrand
  > Announcing a new article about debugging Go programs with GDB.

- [Go App Engine SDK 1.5.5 released](https://go.dev/blog/appengine-155), 2011-10-11, Andrew Gerrand
  > Go App Engine SDK 1.5.5 includes Go release.r60.2.

- [A preview of Go version 1](https://go.dev/blog/go1-preview), 2011-10-05, Russ Cox
  > What the Go team is planning for Go version 1.

- [Learn Go from your browser](https://go.dev/blog/tour), 2011-10-04, Andrew Gerrand
  > Announcing the Go tour, https://tour.golang.org/.

- [The Go image/draw package](https://go.dev/blog/image-draw), 2011-09-29, Nigel Tao
  > An introduction to image compositing in Go using the image/draw package.

- [The Go image package](https://go.dev/blog/image), 2011-09-21, Nigel Tao
  > An introduction to 2-D image processing with the Go image package.

- [The Laws of Reflection](https://go.dev/blog/laws-of-reflection), 2011-09-06, Rob Pike
  > How reflections works in Go, how to think about it, and how to use it.

- [Two Go Talks: "Lexical Scanning in Go" and "Cuddle: an App Engine Demo"](https://go.dev/blog/sydney-gtug), 2011-09-01, Andrew Gerrand
  > Two talks about Go from the Sydney GTUG: Rob Pike explains lexical scanning, and Andrew Gerrand builds a simple real-time chat using App Engine.

- [Go for App Engine is now generally available](https://go.dev/blog/appengine-ga), 2011-07-21, Andrew Gerrand
  > You can use Go on App Engine now!

- [Error handling and Go](https://go.dev/blog/error-handling-and-go), 2011-07-12, Andrew Gerrand
  > An introduction to Go errors.

- [First Class Functions in Go](https://go.dev/blog/functions-codewalk), 2011-06-30, Andrew Gerrand
  > Announcing a new Go codewalk, exploring first class functions.

- [Profiling Go Programs](https://go.dev/blog/pprof), 2011-06-24, Russ Cox, July 2011; updated by Shenghou Ma, May 2013
  > How to use Go's built-in profiler to understand and optimize your programs.

- [Spotlight on external Go libraries](https://go.dev/blog/external-libraries), 2011-06-03, Andrew Gerrand
  > Some popular Go libraries and how to use them.

- [A GIF decoder: an exercise in Go interfaces](https://go.dev/blog/gif-decoder), 2011-05-25, Rob Pike
  > How Go's interfaces work nicely in the Go GIF decoder.

- [Go at Google I/O 2011: videos](https://go.dev/blog/io2011), 2011-05-23, Andrew Gerrand
  > Two talks about Go from Google I/O 2011.

- [Go and Google App Engine](https://go.dev/blog/appengine), 2011-05-10, David Symonds, Nigel Tao, and Andrew Gerrand
  > Announcing support for Go in Google App Engine.

- [Go at Heroku](https://go.dev/blog/heroku), 2011-04-21, Keith Rarick and Blake Mizerany
  > Two Heroku system engineers discuss their experiences using Go.

- [Introducing Gofix](https://go.dev/blog/introducing-gofix), 2011-04-15, Russ Cox
  > How to use go fix to update your code with each new Go release.

- [Godoc: documenting Go code](https://go.dev/blog/godoc), 2011-03-31, Andrew Gerrand
  > How and why to document your Go packages.

- [Gobs of data](https://go.dev/blog/gob), 2011-03-24, Rob Pike
  > Introducing gob, a high-speed Go-to-Go wire encoding format.

- [C? Go? Cgo!](https://go.dev/blog/cgo), 2011-03-17, Andrew Gerrand
  > How to use cgo to let Go packages call C code.

- [Go becomes more stable](https://go.dev/blog/stable-releases), 2011-03-16, Andrew Gerrand
  > Moving from weekly unstable Go releases toward less frequent, more stable ones.

- [JSON and Go](https://go.dev/blog/json), 2011-01-25, Andrew Gerrand
  > How to generate and consume JSON-formatted data in Go.

- [Go Slices: usage and internals](https://go.dev/blog/slices-intro), 2011-01-05, Andrew Gerrand
  > How to use Go slices, and how they work.

- [Go: one year ago today](https://go.dev/blog/1year), 2010-11-10, Andrew Gerrand
  > Happy 1st birthday, Go!

- [Debugging Go code (a status report)](https://go.dev/blog/debug-status), 2010-11-02, Luuk van Dijk
  > What works and what doesn't when debugging Go programs with GDB.

- [Real Go Projects: SmartTwitter and web.go](https://go.dev/blog/smarttwitter), 2010-10-19, Michael Hoisie
  > How Michael Hoisie used Go to build SmartTwitter and web.go.

- [Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/concurrency-timeouts), 2010-09-23, Andrew Gerrand
  > How to implement timeouts using Go's concurrency support.

- [Introducing the Go Playground](https://go.dev/blog/playground-intro), 2010-09-15, Andrew Gerrand
  > Announcing the Go Playground, https://play.golang.org/.

- [Go Wins 2010 Bossie Award](https://go.dev/blog/bossie), 2010-09-06, Andrew Gerrand
  > Go wins a 2010 Bossie Award for “best open source application development software.”

- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover), 2010-08-04, Andrew Gerrand
  > An introduction to the Go's defer, panic, and recover control flow mechanisms.

- [Share Memory By Communicating](https://go.dev/blog/codelab-share), 2010-07-13, Andrew Gerrand
  > A preview of the new Go codelab, Share Memory by Communicating.

- [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax), 2010-07-07, Rob Pike
  > Why Go's declaration syntax doesn't look like, and is much simpler than, C's.

- [Go Programming session video from Google I/O](https://go.dev/blog/io2010), 2010-06-06, Andrew Gerrand
  > A talk by Rob Pike and Russ Cox about Go, from Google I/O 2010.

- [Go at I/O: Frequently Asked Questions](https://go.dev/blog/io2010-faq), 2010-05-27, Andrew Gerrand
  > Q&A about Go from Google I/O 2010.

- [Upcoming Google I/O Go Events](https://go.dev/blog/io2010-preview), 2010-05-12, Andrew Gerrand
  > If you will be at Google I/O 2010, be sure to catch up with the Go team at these events.

- [New Talk and Tutorials](https://go.dev/blog/new-talk-and-tutorials), 2010-05-05, Andrew Gerrand
  > More materials for learning about Go: one talk, one codelab, and one screencast.

- [JSON-RPC: a tale of interfaces](https://go.dev/blog/json-rpc), 2010-04-27, Andrew Gerrand
  > How to use the net/rpc package's interfaces to create a JSON-RPC system.

- [Third-party libraries: goprotobuf and beyond](https://go.dev/blog/protobuf), 2010-04-20, Andrew Gerrand
  > Announcing Go support for Protocol Buffers, Google's data interchange format.

- [Go: What's New in March 2010](https://go.dev/blog/hello-world), 2010-03-18, Andrew Gerrand
  > First post!


