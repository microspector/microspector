# Contributing

Thanks for your interest in contributing to Microspector!

## Code of Conduct

Help us keep Microspector open and inclusive. Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md).

## Getting Started

* submit a ticket for your issue, assuming one does not already exist
  * clearly describe the issue including steps to reproduce when it is a bug
  * identify specific versions of the binaries and client libraries
* fork the repository on GitHub

## Making Changes

* create a branch from where you want to base your work
  * we typically name branches according to the following format: `fix/<issue_number>`
* make commits of logical units
* make sure your commit messages are in a clear and readable format, example:
  
```
fixed bug in http command
  
* fixed ssl handshake
* cleanup variable names
* ...
```

* if you're fixing a bug or adding functionality you have to write a test that fails without your fix
* make sure to run `make test` in the root of the repo to ensure that your code is
  properly formatted and that tests pass.
    * test must prove the feature presents and fail when you revert your changes

## Submitting Changes

* push your changes to your branch in your fork of the repository
* submit a pull request against microspector's repository