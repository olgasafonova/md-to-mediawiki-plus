# Markdown to MediaWiki Plus

A robust, enterprise-grade Markdown to MediaWiki converter with Tietoevry branding and accessibility-first design.

This is an enhanced version of the original converter, featuring significant reliability improvements, bug fixes, and a comprehensive test suite.

## Features

- **🎨 Tietoevry Branding**:
    - Purple gradient headings (Hero Blue → Light Purple)
    - Peach + Hero Blue inline code styling
    - WCAG AA compliant colors
- **🔄 Smart Changelog**: Automatically reverses changelog order (Newest First) with robust parsing logic.
- **✅ Accessibility**: Green emoji checkmarks and accessible syntax highlighting.
- **🛡️ Reliable**: Fully tested codebase with unit tests and CI integration.
- **⚡ Fast**: High-performance Go implementation.

## Installation

### Prerequisites
- Go 1.21 or later

### Build from Source

```bash
git clone https://github.com/olgasafonova/md-to-mediawiki-plus.git
cd md-to-mediawiki-plus
make build
```

## Usage

```bash
./md-to-mediawiki-plus -i input.md -o output.wiki
```

### Options

- `-i, --input`: Input Markdown file (required)
- `-o, --output`: Output MediaWiki file (default: stdout)
- `--with-css`: Include CSS styling in output
- `-c, --concurrent`: Enable concurrent processing (currently runs sequentially for stability)
- `-v, --version`: Show version
- `-h, --help`: Show help

## Development

This project includes a full development environment setup.

### Running Tests
```bash
make test
```

### Linting
```bash
make lint
```

### CI/CD
Includes GitHub Actions workflow for automated testing and building.

## License
MIT
