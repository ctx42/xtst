# Introduction to Ctx42 Testing Libraries

<!-- TOC -->
* [Introduction to Ctx42 Testing Libraries](#introduction-to-ctx42-testing-libraries)
  * [Core Principles: Simplicity and Usability](#core-principles-simplicity-and-usability)
  * [Modular and Extensible Design](#modular-and-extensible-design)
  * [Current Status](#current-status)
<!-- TOC -->

This repository marks the beginning of Ctx42 Testing Libraries, a collection of 
testing and assertion packages poised to help developers approach testing. As 
it develops, it will offer a comprehensive suite of tools designed to make 
testing more efficient, enjoyable, and integral to the development process.

Whether you're a seasoned tester or just starting out, tools in this module is 
being crafted to meet your needs, providing a solid foundation for ensuring 
code reliability in projects of all sizes.

## Core Principles: Simplicity and Usability

At the heart of the module lies a commitment to minimalism and an exceptional
developer experience (DX). By maintaining zero external dependencies, the
framework stays lightweight and fast, free from the potential complexities and
conflicts that third-party libraries can introduce. Our documentation is being
carefully designed to be thorough, clear, and packed with practical examples,
empowering you to master the module with ease. Expect a great DX with features 
like a fluent, chainable API, descriptive error messages for quick debugging, 
all tailored to streamline your testing workflow.

## Modular and Extensible Design

`Xtst` is built as a collection of modular, laser-focused packages, each 
targeting a specific aspect of testing. For instance, you might leverage the 
`assert` package for assertions, the `mock` and `mocker` packages for test 
doubles, or the `tstkit` package to keep your tests readable and minimalistic. 
The modularity lets you customize your testing setup to fit your project’s 
exact needs, avoiding unnecessary overhead. Beyond customization, the 
extensible architecture invites you to create your own test helpers ensuring.

## Current Status

Currently, we work on adding core building blocks of the testing library.

So far we've completed:

- Package [must](pkg/must) with basic test helpers which panic on error.
- Package [notice](pkg/notice) helping to create and format assertion messages.
- Package [tester](pkg/tester) providing facilities to test `Test Helpers`.

Each of the packages have its own README.md file with documentation, and
`examples_test.go` file with usage examples.
