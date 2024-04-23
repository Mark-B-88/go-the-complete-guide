# Go The Complete Guide

This is a [Udemy course](https://www.udemy.com/course/go-the-complete-guide/) created by [Maximilian Schwarzmüller](https://www.udemy.com/user/maximilian-schwarzmuller/)

---

# Using Nodemon to Watch for File Changes

```json
{
	"scripts": {
		"watch": "nodemon --ext go --exec go run ."
	},
	"devDependencies": {
		"nodemon": "^3.1.0"
	}
}
```

---

# What is Go?

Go is an open-source programming language developed by Google.

- It was developed in 2007, and made public in 2009
- Focuses on simplicity, clarity & scalability

  - Inspired by languages like Python
    - Aims to provide a clean, understandable syntax

- High performance & focus on concurrency

  - Similar to C or C++
    - Popular for tasks that benefit from multi-threading

- Batteries included

  - Go comes with a standard library
    - Many core features are built-in

- Static typing

  - Go is a type-safe language
    - Allows you to catch many errors early

- Popular Use Cases
  - Networking & API's
    - Microservices
    - CLI Tools

---

# Go Essentials

Understanding the **_key_** components of a Go Program

- Values
- Basic Types
- Core Language Features

---

# Packages - A Closer Look

Working with Go Packages

- Splitting Code Across Multiple Files
- Splitting Code Across Multiple Packages
- Importing & Using Custom Packages

---

# Understanding Pointers

Working With Addresses Instead of Values

- What are pointers?
- Why does this feature exist?
- How can you work with pointers?

---

# Structs

Grouping Data & Functions Into Collections

- What are Structs?
- Creating & Using Structs
- Adding Methods To Structs

---

# Arrays, Slices & Maps

Storing Data in Collections

- Understanding Arrays & Array limitations
- Understanding & using slices
- Slices Deep Dive
- Understanding & Using Maps

---

# Functions Deep Dive

Beyond the basics

- Using Functions as Values
- Anonymous Functions
- Recursion
- Variadic Functions

---

# Concurrency

Running Tasks In Parallel

- Understanding Goroutines
- Sending Data with Channels
- Controlling Code Flow & Simultaneous Tasks

---

# REST API

Building a REST API with Go

- Planning The Project
- Building the REST API Step By Step

---
