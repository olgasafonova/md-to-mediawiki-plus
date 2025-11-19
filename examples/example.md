# API Documentation Example

This example demonstrates all the features of the enhanced Markdown to MediaWiki converter.

## File Service API

### Overview

The File Service API provides methods for managing files and documents. Use the `FileService/GetFile` method to retrieve files and the `FileService/UploadFile` method to upload new documents.

> [!info] All API endpoints require authentication
> Make sure to include your API token in the Authorization header.

### Example Request

Here's how to call the file service using JavaScript:

```javascript
const API_URL = 'https://api.example.com';

async function getFile(fileId) {
  const response = await fetch(`${API_URL}/FileService/GetFile`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer YOUR_TOKEN'
    },
    body: JSON.stringify({
      FileRecno: fileId,
      IncludeMetadata: true
    })
  });

  return await response.json();
}
```

### Response Format

The API returns JSON data:

```json
{
  "fileId": 12345,
  "fileName": "document.pdf",
  "size": 2048576,
  "mimeType": "application/pdf",
  "metadata": {
    "author": "John Doe",
    "created": "2024-01-15T10:30:00Z",
    "modified": "2024-01-20T14:45:00Z"
  }
}
```

> [!success] Files up to 100MB are supported
> Larger files should use the chunked upload endpoint.

## Configuration

### XML Configuration

Configure the service using XML:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
  <!-- File service settings -->
  <fileService>
    <maxFileSize>104857600</maxFileSize>
    <allowedTypes>
      <type>pdf</type>
      <type>docx</type>
      <type>xlsx</type>
    </allowedTypes>
    <storage>
      <provider>AzureBlob</provider>
      <container>documents</container>
    </storage>
  </fileService>
</configuration>
```

### Python Example

Here's a Python implementation:

```python
import requests

def upload_file(file_path, metadata=None):
    """Upload a file to the service"""
    url = f"{API_URL}/FileService/UploadFile"

    # Prepare multipart form data
    with open(file_path, 'rb') as f:
        files = {'file': f}
        data = {'metadata': metadata} if metadata else {}

        response = requests.post(
            url,
            files=files,
            data=data,
            headers={'Authorization': f'Bearer {API_TOKEN}'}
        )

    return response.json()
```

> [!warning] Rate Limiting
> The API is limited to 100 requests per minute. Exceeding this will result in 429 errors.

## API Methods

Available methods include:

* `FileService/GetFile` - Retrieve a file by ID
* `FileService/UploadFile` - Upload a new file
* `FileService/DeleteFile` - Delete an existing file
* `FileService/ListFiles` - Get all files
  * Supports pagination
  * Supports filtering by type

### Ordered Process

1. Authenticate with the API
2. Upload your file
3. Store the returned file ID
4. Use the ID to retrieve or manage the file

## Code Formatting

Inline code like `const x = 42;` is also highlighted. API endpoints such as `DocumentService/CreateDocument` get special Tietoevry branding.

### Tables

Here's a comparison of file types:

| File Type | Max Size | MIME Type | Supported |
|-----------|----------|-----------|-----------|
| PDF | 100 MB | application/pdf | Yes |
| DOCX | 50 MB | application/vnd.openxmlformats-officedocument.wordprocessingml.document | Yes |
| XLSX | 25 MB | application/vnd.openxmlformats-officedocument.spreadsheetml.sheet | Yes |
| PNG | 10 MB | image/png | Yes |

> [!note] Additional formats can be enabled
> Contact support to enable additional file format support.

## Highlights and Emphasis

You can use ==highlighted text== to draw attention to important information. Regular **bold** and *italic* formatting is also supported.

## Links

For more information, see the [official documentation](https://docs.example.com) or visit our [[Internal Wiki]].

---

This example demonstrates:
- Code syntax highlighting with accessible colors
- Tietoevry-branded callout boxes
- API endpoint highlighting
- Auto-detection of code languages
- Table conversion
- All standard Markdown features
