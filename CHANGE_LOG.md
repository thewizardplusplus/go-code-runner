# Change Log

## [v1.4](https://github.com/thewizardplusplus/go-code-runner/tree/v1.4) (2021-05-15)

## [v1.3](https://github.com/thewizardplusplus/go-code-runner/tree/v1.3) (2021-04-06)

## [v1.2](https://github.com/thewizardplusplus/go-code-runner/tree/v1.2) (2021-03-27)

Checking of package imports used in the code written in the Go programming language.

- checking of package imports used in the code written in the Go programming language:
  - checking based on the list of allowed imports;
- compiling of a code written in the Go programming language:
  - checking of package imports used in the code (optionally):
    - checking based on the list of allowed imports.

### Features

- saving of a code to a temporary file:
  - storing of the temporary file with the code to an individual temporary directory;
- checking of package imports used in the code written in the Go programming language:
  - checking based on the list of allowed imports;
- compiling of a code written in the Go programming language:
  - automatic importing of the packages used in the code;
  - checking of package imports used in the code (optionally):
    - checking based on the list of allowed imports;
  - enriching of an error of the external command running by an output from the stderr stream;
- running of the compiled code (i.e. the executable file):
  - passing of a custom input as the stdin stream;
  - returning of an output from the stdout stream;
  - enriching of an error of the external command running by an output from the stderr stream;
- running of a test case set for the compiled code (i.e. the executable file):
  - representation of a test case:
    - input;
    - expected output;
  - checking of an actual output in each test case:
    - returning of the [sentinel errors](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully#sentinel%20errors):
      - failed running &mdash; it returns on a running error;
      - unexpected output &mdash; it returns when the expected and actual outputs do not match.

## [v1.1](https://github.com/thewizardplusplus/go-code-runner/tree/v1.1) (2021-03-27)

Running of a test case set for the compiled code (i.e. the executable file).

- running of a test case set for the compiled code (i.e. the executable file):
  - representation of a test case:
    - input;
    - expected output;
  - checking of an actual output in each test case:
    - returning of the [sentinel errors](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully#sentinel%20errors):
      - failed running &mdash; it returns on a running error;
      - unexpected output &mdash; it returns when the expected and actual outputs do not match.

### Features

- saving of a code to a temporary file:
  - storing of the temporary file with the code to an individual temporary directory;
- compiling of a code written in the Go programming language:
  - automatic importing of the packages used in the code;
  - enriching of an error of the external command running by an output from the stderr stream;
- running of the compiled code (i.e. the executable file):
  - passing of a custom input as the stdin stream;
  - returning of an output from the stdout stream;
  - enriching of an error of the external command running by an output from the stderr stream;
- running of a test case set for the compiled code (i.e. the executable file):
  - representation of a test case:
    - input;
    - expected output;
  - checking of an actual output in each test case:
    - returning of the [sentinel errors](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully#sentinel%20errors):
      - failed running &mdash; it returns on a running error;
      - unexpected output &mdash; it returns when the expected and actual outputs do not match.

## [v1.0](https://github.com/thewizardplusplus/go-code-runner/tree/v1.0) (2021-03-22)

Major version. Implementing of the compiling and running of a code written in the Go programming language.

### Features

- saving of a code to a temporary file:
  - storing of the temporary file with the code to an individual temporary directory;
- compiling of a code written in the Go programming language:
  - automatic importing of the packages used in the code;
  - enriching of an error of the external command running by an output from the stderr stream;
- running of the compiled code (i.e. the executable file):
  - passing of a custom input as the stdin stream;
  - returning of an output from the stdout stream;
  - enriching of an error of the external command running by an output from the stderr stream.
