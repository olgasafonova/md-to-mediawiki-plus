# Best Practices Example

This example demonstrates formatting patterns that work best with MediaWiki conversion.

## Handling Code with Underscores

When documenting code that contains underscores, always wrap them in backticks to prevent MediaWiki from interpreting them as italic markers.

### Document Types Configuration

The system supports several KOBO document types:

- `KOBO_MELDINGSDIALOG` - Main message dialog
- `KOBO_MELDINGSDIALOG_VEDLEGG` - Message attachments
- `KOBO_SVARUT` - Outgoing messages
- `KOBO_SVARINN` - Incoming messages

Without backticks, KOBO_MELDINGSDIALOG_VEDLEGG would render incorrectly with italics.

---

## Configuration Structure with Proper Spacing

Use horizontal rules between major sections to improve readability.

### System Components

**External System Integration**

Handles incoming documents from GEMI, SvarInn, and SIF API.

⬇️

**AutoSaver Properties**

Central configuration that enables the archiving process.

⬇️

**Archiving Process**

Defines which origins to process and how to handle them.

---

## Structured Lists Without Nesting Issues

Instead of nested bullets under numbered items, use bold headers with manual numbering.

### Field Mapping Configuration

Each field mapping has three components:

**1. External Schema Property** (required)

Field name from source metadata. Examples: `Subject`, `FirstName`, `DocumentDate`. Supports nested paths like `Person.FirstName` or `Document.Title`.

**2. Target Property** (required)

Public 360° field name loaded dynamically. Examples: Title, Category, Status, DocumentDate.

**3. Value** (optional)

Can be a literal value, placeholder, or expression. Leave empty to use source field value directly.

---

## Process Workflow Example

This example shows how to document step-by-step processes:

**Step 1: Load Configuration**

Read 360° AutoSaver settings from Registers and verify the process is enabled.

**Step 2: Read Unregistered Documents**

Filter documents by configured origin (examples: `GEMI_INVOICE`, `SVARINN_MESSAGE`).

**Step 3: Process Each Document**

Check for metadata file (XML or JSON attachment), read and convert if present, apply schema mapping rules or use default values.

**Step 4: Archive Document**

Translate to SIF objects, execute configured plugins, call SIF handler to complete archiving.

---

This example demonstrates:

- Proper use of backticks for codes with underscores
- Horizontal rules for section spacing
- Bold headers instead of nested lists
- Clear visual hierarchy without rendering issues
