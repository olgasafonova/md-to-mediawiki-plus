# Comprehensive Example: All Markdown Features

This example demonstrates every feature supported by the md-to-mediawiki-plus converter.

## Heading Level 2

### Heading Level 3

#### Heading Level 4

##### Heading Level 5

###### Heading Level 6

---

## Inline Code Examples

Use `inline code` with special highlighting. API methods like `DocumentService/CreateDocument` and `FileService/GetFile` are highlighted with Tieto branding.

Variables like `API_TOKEN`, `fileId`, and `connectionString` get ==yellow background with navy text==.

---

## Code Blocks with Syntax Highlighting

### JSON Example

```json
{
  "apiVersion": "v1",
  "kind": "Configuration",
  "metadata": {
    "name": "autosaver-config",
    "namespace": "production",
    "labels": {
      "environment": "prod",
      "app": "autosaver"
    }
  },
  "data": {
    "endpoint": "https://api.example.com/v1",
    "timeout": 30000,
    "retryAttempts": 3,
    "enableLogging": true
  }
}
```

### XML Example

```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
  <!-- AutoSaver Schema Configuration -->
  <schema type="document" version="2.0">
    <mapping>
      <field source="title" target="CreateDocumentParameter.Title" required="true"/>
      <field source="description" target="CreateDocumentParameter.Description"/>
      <field source="category" target="DocumentCategory" default="INN"/>
    </mapping>
    <validation>
      <rule field="title" minLength="3" maxLength="255"/>
      <rule field="documentDate" format="yyyy-MM-dd"/>
    </validation>
    <defaults>
      <value key="DocumentArchive">SAKSDOK</value>
      <value key="Status">F</value>
    </defaults>
  </schema>
</configuration>
```

### C# Example

```csharp
using System;
using System.Threading.Tasks;
using Public360.Services;

namespace AutoSaver.Integration
{
    /// <summary>
    /// Handles document archiving from external systems
    /// </summary>
    public class DocumentArchiver : IArchiveService
    {
        private readonly ISifClient _client;
        private readonly ILogger<DocumentArchiver> _logger;

        public DocumentArchiver(ISifClient client, ILogger<DocumentArchiver> logger)
        {
            _client = client ?? throw new ArgumentNullException(nameof(client));
            _logger = logger;
        }

        public async Task<ArchiveResult> ArchiveDocumentAsync(DocumentRequest request)
        {
            try
            {
                // Validate request
                if (string.IsNullOrEmpty(request.Title))
                    throw new ValidationException("Title is required");

                // Create document
                var documentId = await _client.CreateDocumentAsync(new CreateDocumentParameter
                {
                    Title = request.Title,
                    Description = request.Description,
                    DocumentDate = DateTime.Now,
                    Category = request.Category ?? "INN",
                    Archive = "SAKSDOK"
                });

                _logger.LogInformation("Document created: {DocumentId}", documentId);

                return new ArchiveResult { Success = true, DocumentId = documentId };
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Failed to archive document");
                return new ArchiveResult { Success = false, Error = ex.Message };
            }
        }
    }
}
```

### Go Example

```go
package autosaver

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
)

// SchemaMapper handles mapping between external schemas and P360 parameters
type SchemaMapper struct {
    config *Configuration
    logger Logger
}

// DocumentMapping represents field mappings for document creation
type DocumentMapping struct {
    SourceField  string `json:"sourceField"`
    TargetField  string `json:"targetField"`
    Required     bool   `json:"required"`
    DefaultValue string `json:"defaultValue,omitempty"`
}

// MapDocument converts external document format to P360 parameters
func (m *SchemaMapper) MapDocument(ctx context.Context, input map[string]interface{}) (*CreateDocumentParams, error) {
    params := &CreateDocumentParams{
        CreatedDate: time.Now(),
        Archive:     "SAKSDOK",
    }

    // Map title (required field)
    if title, ok := input["title"].(string); ok && title != "" {
        params.Title = title
    } else {
        return nil, fmt.Errorf("title is required")
    }

    // Map optional fields
    if desc, ok := input["description"].(string); ok {
        params.Description = desc
    }

    if category, ok := input["category"].(string); ok {
        params.Category = category
    } else {
        params.Category = "INN" // default
    }

    m.logger.Info("Document mapped successfully", "title", params.Title)

    return params, nil
}
```

### TypeScript Example

```typescript
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import axios, { AxiosInstance } from 'axios';

/**
 * Service for interacting with P360 SIF API
 */
@Injectable()
export class SifApiService {
  private readonly client: AxiosInstance;
  private readonly baseUrl: string;

  constructor(private configService: ConfigService) {
    this.baseUrl = this.configService.get<string>('P360_API_URL');

    this.client = axios.create({
      baseURL: this.baseUrl,
      timeout: 30000,
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      }
    });

    // Add authentication interceptor
    this.client.interceptors.request.use(async (config) => {
      const token = await this.getAuthToken();
      config.headers.Authorization = `Bearer ${token}`;
      return config;
    });
  }

  async createDocument(params: CreateDocumentParams): Promise<DocumentResult> {
    try {
      const response = await this.client.post<DocumentResponse>(
        '/DocumentService/CreateDocument',
        {
          Title: params.title,
          Description: params.description,
          DocumentDate: params.documentDate.toISOString(),
          Category: params.category || 'INN',
          Archive: params.archive || 'SAKSDOK',
          Status: 'F'
        }
      );

      return {
        success: true,
        documentId: response.data.DocumentId,
        recno: response.data.Recno
      };
    } catch (error) {
      console.error('Failed to create document:', error);
      throw new Error(`Document creation failed: ${error.message}`);
    }
  }

  private async getAuthToken(): Promise<string> {
    // Token acquisition logic
    return this.configService.get<string>('P360_API_TOKEN');
  }
}

interface CreateDocumentParams {
  title: string;
  description?: string;
  documentDate: Date;
  category?: string;
  archive?: string;
}

interface DocumentResult {
  success: boolean;
  documentId?: number;
  recno?: number;
  error?: string;
}
```

---

## Tables

### File Format Support

| Format | Extension | MIME Type | Max Size | Supported |
|--------|-----------|-----------|----------|-----------|
| PDF | .pdf | application/pdf | 100 MB | ✅ Yes |
| Word | .docx | application/vnd.openxmlformats-officedocument.wordprocessingml.document | 50 MB | ✅ Yes |
| Excel | .xlsx | application/vnd.openxmlformats-officedocument.spreadsheetml.sheet | 25 MB | ✅ Yes |
| PNG | .png | image/png | 10 MB | ✅ Yes |
| JPEG | .jpg, .jpeg | image/jpeg | 10 MB | ✅ Yes |
| ZIP | .zip | application/zip | 200 MB | ✅ Yes |
| CSV | .csv | text/csv | 5 MB | ✅ Yes |

### API Endpoints Comparison

| Endpoint | Method | Authentication | Rate Limit |
|----------|--------|----------------|------------|
| `/DocumentService/CreateDocument` | POST | Required | 100/min |
| `/DocumentService/GetDocument` | GET | Required | 1000/min |
| `/FileService/UploadFile` | POST | Required | 50/min |
| `/FileService/GetFile` | GET | Required | 500/min |
| `/CaseService/CreateCase` | POST | Required | 100/min |

---

## Lists and Nesting

### Simple Bullet Points

- First item
- Second item
- Third item
- Fourth item

### Simple Numbered List

1. First step
2. Second step
3. Third step
4. Fourth step

### Nested Bullet Points

- Main category 1
  - Subcategory 1.1
  - Subcategory 1.2
    - Sub-subcategory 1.2.1
    - Sub-subcategory 1.2.2
  - Subcategory 1.3
- Main category 2
  - Subcategory 2.1
    - Sub-subcategory 2.1.1
      - Deep nesting 2.1.1.1
      - Deep nesting 2.1.1.2
    - Sub-subcategory 2.1.2
  - Subcategory 2.2
- Main category 3

### Nested Numbered Lists

1. First major step
   1. Sub-step 1.1
   2. Sub-step 1.2
      1. Detail 1.2.1
      2. Detail 1.2.2
   3. Sub-step 1.3
2. Second major step
   1. Sub-step 2.1
      1. Detail 2.1.1
         1. Fine detail 2.1.1.1
         2. Fine detail 2.1.1.2
      2. Detail 2.1.2
   2. Sub-step 2.2
3. Third major step

### Mixed Nested Lists (Bullets in Numbers)

1. Configure the system
   - Open admin panel
   - Navigate to settings
     - Select "AutoSaver Configuration"
     - Choose schema type
   - Save changes
2. Map the fields
   - Define source fields
     - title
     - description
     - documentDate
   - Map to target parameters
     - CreateDocumentParameter.Title
     - CreateDocumentParameter.Description
3. Test the mapping
   - Use test data
   - Verify output
   - Check for errors

### Mixed Nested Lists (Numbers in Bullets)

- Development setup
  1. Install prerequisites
  2. Clone repository
  3. Configure environment
     1. Copy .env.example to .env
     2. Set API credentials
     3. Configure endpoints
  4. Run tests
- Production deployment
  1. Build application
  2. Run security scan
     1. Check dependencies
     2. Review code
  3. Deploy to server
  4. Monitor logs
- Maintenance tasks

---

## Highlighted Text and Inline Styling

Regular text with ==highlighted yellow background== and ==navy text color==.

Important: Use the `--with-css` flag to enable ==inline code styling==.

API tokens like `ABC123XYZ` should be stored securely.

Configuration values such as `timeout: 30000` and `maxRetries: 3` can be customized.

---

## Dividers

Use three hyphens for a divider:

---

And they create clear visual separation between sections.

---

## Icons and Callouts

> [!info] Information
> This is an informational callout with an info icon.

> [!warning] Warning
> This is a warning callout. Rate limits apply to all API endpoints.

> [!success] Success
> Document archived successfully with ID: 12345

> [!note] Note
> All dates should be in ISO 8601 format: `2024-01-15T10:30:00Z`

> [!tip] Pro Tip
> Use schema export functionality to reuse configurations across environments.

---

## Text Formatting

**Bold text** for emphasis.

*Italic text* for subtle emphasis.

***Bold and italic*** for strong emphasis.

~~Strikethrough~~ for deprecated features.

Regular text with `inline code` and more regular text.

Combine **bold with `code`** and *italic with `code`*.

---

## Links

External link: [Public 360 Documentation](https://www.tieto.com/p360)

Internal reference: See [Heading Level 2](#heading-level-2) section.

URL directly: https://www.example.com

Email: support@example.com

---

## Complex Nested Example

This shows all nesting patterns in one structure:

1. **AutoSaver Setup**
   - Prerequisites
     - Go 1.21+
     - P360 instance
     - SIF API access
   - Installation steps
     1. Download binary
     2. Configure settings
        - Set `endpoint` to your API URL
        - Add authentication token
        - Configure schema mapping
           1. Define source format (XML or JSON)
           2. Map fields to P360 parameters
              - title → `CreateDocumentParameter.Title`
              - date → `CreateDocumentParameter.DocumentDate`
           3. Set default values
        - Test connection
     3. Start service
   - Verification
     - Check logs for errors
     - Test with sample document
       1. Create test JSON file
       2. Submit via API
       3. Verify in P360
          - Check document created
          - Verify metadata
          - Confirm file attachment

2. **Production Deployment**
   1. Security configuration
      - Enable TLS
      - Configure firewalls
        - Allow inbound on port 443
        - Restrict source IPs
      - Set up monitoring
   2. Performance tuning
      - Increase timeout values
        - Connection timeout: `30s`
        - Read timeout: `60s`
      - Configure caching
      - Enable compression
   3. Backup procedures

---

## Summary

This comprehensive example includes:

- ✅ All heading levels (H1 through H6)
- ✅ Code blocks with syntax highlighting (JSON, XML, C#, Go, TypeScript)
- ✅ Inline code with styling
- ✅ Tables with multiple columns
- ✅ Simple bullet points
- ✅ Simple numbered lists
- ✅ Nested bullet points (multiple levels)
- ✅ Nested numbered lists (multiple levels)
- ✅ Mixed nesting (bullets in numbers, numbers in bullets)
- ✅ Highlighted text with ==yellow background==
- ✅ Divider lines (---)
- ✅ Icons and callout boxes
- ✅ Text formatting (bold, italic, strikethrough)
- ✅ Links (external, internal, URLs)
- ✅ Complex nested structures combining all patterns

---

**Note**: When using the converter, include the `--with-css` flag for full styling support.
