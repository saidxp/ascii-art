# ASCII Art

A Go command-line program that turns a string into a banner and prints it to standard output. Each printable ASCII character has an eight-row drawing in the bundled [Standard font](acii-art/standard.txt). To print a word, the program joins the first row of every character, then the second, and so on.

## Run locally

Requires **Go 1.22.3 or later**. The project uses only the Go standard library.

**Build note:** The current [entry point](acii-art/main.go) ends with an extra `}` after `main` has closed. Remove that final brace before running the examples below; otherwise Go reports `syntax error: non-declaration statement outside function body`.

From the repository root:

```sh
cd acii-art
go run . 'Hello, Go!'
```

Keep the quotes around text containing spaces so it reaches the program as one argument. Run from `acii-art` because the program looks for `standard.txt` in its working directory.

## Examples

- `go run . 'A'` prints a single glyph. Its visible rows are shown below; the font also includes blank rows.
- `go run . 'Hello\nGo'` prints two banners with a blank line between them. Here `\n` is a literal backslash followed by `n`, not an actual newline in the argument.
- `go run . 'Hello' > banner.txt` saves the banner instead of displaying it in the terminal.

```text
    /\
   /  \
  / /\ \
 / ____ \
/_/    \_\
```