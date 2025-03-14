If you’ve spent any time writing Go tests, you’ve probably encountered the joy 
of `*testing.T`. It’s the backbone of Go’s testing framework—powerful, flexible, 
and ubiquitous. But as your test suite grows, you might find yourself repeating
the same chunks of test logic across multiple test cases. Enter _test helpers_:
reusable functions that streamline your tests, improve readability, and reduce
complexity. Libraries like assert are prime examples, turning verbose checks
into concise assertions.

But here’s the catch: how do you test the test helpers themselves? After all,
these are the tools you rely on to ensure your code works as expected. If they
fail, your tests might silently lie to you. This is where the `tester` package,
built by thoughtful Go developers, comes to the rescue. 

See the detailed documentation at [README.md](pkg/tester/README.md)
