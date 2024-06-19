 You are discussing the process of testing software components, specifically in PHP, using DepTrac tool. Here's a summary of your steps:

1. Define layers (components) in a configuration file, specifying their directories.
2. Use DepTrac to identify changed files not present in the origin branch.
3. Run tests selectively for the affected component(s), using PHPUnit and passing the test suite for that specific component as an argument.
4. Optionally, create a script (test_dosh) to run the entire test suite, compare the runtime with and without DepTrac, and provide statistics on the time saved.

Your approach helps in improving the efficiency of testing by running only necessary tests, saving both time and resources. It also allows for more thorough testing through integration tests for dependent components if needed, ensuring that any potential issues not covered by contract tests are addressed.
