# Markdown to MediaWiki Plus

Convert Markdown documents to MediaWiki format with Tieto branding.

## What This Does

This tool transforms your Markdown files into MediaWiki wiki markup, automatically applying:
- Tieto brand colors (Hero Blue #021e57 for headings)
- Accessible color contrast (WCAG AA compliant)
- Smart changelog formatting (newest entries first)
- Styled code blocks with yellow and blue accents

## Quick Start

```bash
# Build the tool
git clone https://github.com/olgasafonova/md-to-mediawiki-plus.git
cd md-to-mediawiki-plus
make build

# Convert a file
./md-to-mediawiki-plus -i input.md -o output.txt --with-css
```

Copy the contents of `output.txt` and paste into your MediaWiki editor.

## Important Requirements

**File Output:**
- Always save output as `.txt` files (not `.wiki`)
- Save output files outside your Obsidian vault directory

**System Requirements:**
- Go 1.21 or later

## Usage

### Basic Conversion

```bash
./md-to-mediawiki-plus -i input.md -o ~/Documents/output.txt
```

### With CSS Styling

```bash
./md-to-mediawiki-plus -i input.md -o ~/Documents/output.txt --with-css
```

### Command Options

| Option | Description |
|--------|-------------|
| `-i, --input` | Input Markdown file (required) |
| `-o, --output` | Output file path (default: prints to screen) |
| `--with-css` | Include CSS styling for colors and formatting |
| `-v, --version` | Show version |
| `-h, --help` | Show help |

## What Gets Converted

### Headings
Markdown headings become MediaWiki headings with Hero Blue color applied.

### Code Blocks
Inline `code` gets yellow background with Hero Blue text. Code blocks use syntax highlighting.

### Changelogs
If your Markdown contains a changelog, entries are automatically reversed to show newest first.

### Lists and Formatting
Standard Markdown lists, bold, italic, and links convert to their MediaWiki equivalents.

## Examples

See the `examples/` directory for sample input and output files.

## Troubleshooting

**Problem:** Output file won't save
- **Solution:** Make sure you're saving outside your Obsidian vault and using `.txt` extension

**Problem:** Colors don't appear in wiki
- **Solution:** Use the `--with-css` flag and ensure your MediaWiki instance allows inline styles

**Problem:** Build fails
- **Solution:** Verify you have Go 1.21+ installed: `go version`

## For Developers

### Running Tests
```bash
make test
```

### Code Quality
```bash
make lint
```

### CI/CD
The project includes GitHub Actions for automated testing and builds.

## License

MIT

---

**Version History:** See `CHANGELOG.md` for release notes and version history.
