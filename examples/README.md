# Examples Directory

This directory contains comprehensive sample files demonstrating all converter features.

## Files

### comprehensive-example.md

**Purpose**: Complete reference showcasing ALL converter features

**Includes**:
- ✅ All heading levels (H1-H6) with Hero Blue styling
- ✅ Code blocks with syntax highlighting for:
  - JSON
  - XML
  - C#
  - Go
  - TypeScript
- ✅ Inline code with yellow background and navy text
- ✅ Tables (multiple columns, formatted data)
- ✅ Lists:
  - Simple bullet points
  - Simple numbered lists
  - Nested bullet points (3+ levels deep)
  - Nested numbered lists (3+ levels deep)
  - Mixed nesting (bullets in numbers, numbers in bullets)
- ✅ Highlighted text with ==yellow background==
- ✅ Horizontal dividers (---)
- ✅ Icons and callout boxes (info, warning, success, note, tip)
- ✅ Text formatting (bold, italic, strikethrough)
- ✅ Links (external, internal anchors, URLs, emails)
- ✅ Complex nested structures

**Use case**: Template for creating your own documentation with full feature support

---

### comprehensive-example.txt

**Purpose**: Pre-converted MediaWiki output of comprehensive-example.md

**Use case**:
- Reference for expected converter output
- Quick copy-paste to test MediaWiki rendering
- Verify your conversion matches expected format

---

## How to Use

### Convert the example:

```bash
# Basic conversion (no CSS)
./md-to-mediawiki-plus -i examples/comprehensive-example.md -o output.txt

# With CSS styling (recommended)
./md-to-mediawiki-plus -i examples/comprehensive-example.md -o output.txt --with-css
```

### View the output:

Open the generated `.txt` file and copy its contents into your MediaWiki editor.

The CSS styling will apply automatically when the page is saved, giving you:
- Hero Blue headings (#021e57)
- Yellow + Navy inline code (`#f5ff56` background, `#021e57` text)
- WCAG AA compliant syntax highlighting
- Proper table formatting
- Styled callout boxes

---

## Creating Your Own Documentation

Use `comprehensive-example.md` as a template:

1. Copy sections you need from the comprehensive example
2. Replace with your actual content
3. Run the converter with `--with-css` flag
4. Copy output to MediaWiki
5. Verify rendering and styling

---

## Color Scheme Reference

**Tieto Brand Colors**:
- Hero Blue: `#021e57` (headings, inline code text)
- Yellow: `#f5ff56` (inline code background)

**Syntax Highlighting** (WCAG AA compliant):
- Comments: `#A0A1A7` (gray, italic)
- Keywords: `#0000FF` (blue, bold)
- Strings: `#50A14F` (green)
- Numbers: `#A626A4` (purple)
- Functions: `#4078F2` (light blue, semi-bold)
- XML tags: `#E45649` (red, bold)
- XML attributes: `#986801` (orange)
