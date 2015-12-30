# Sourcify

Sourcify is a command line that tool reads stdin and outputs it back to stdout.
If input line looks like `<some_file_name>:<line_number>` then Sourcify
would print a source file content at the specified line number +-5 lines.
It is intended to be used with `llvm-symbolizer`.
