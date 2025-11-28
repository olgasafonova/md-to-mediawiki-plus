
## SIF API Enhancements

### File Variants Download Support

**What changed:**

- Added new SIF API capability to download all variants of a specific file version in a single zipped package
- Introduced two new methods:
  - `GenerateFileVariantsDownloadToken/GetAllFileVariantsAsZip` - Returns a token for file variants
  - Use existing `FileService/GetFile` to download the actual zip using the token
- Filter by FileRecno, UID, and Version
- Optional IncludeMetadata parameter for flexibility
- Returns zip token from temp storage

**Why it matters:**

This capability was not previously available in the SIF API. The new feature makes it faster and more efficient to:
- Export complete document sets for archival or external submissions
- Share all available formats of a document with external stakeholders
- Integrate document retrieval into automated workflows and microservices

The API provides flexibility with optional metadata inclusion, so you only retrieve what you need.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/382606
- https://tieto-si.visualstudio.com/360/_workitems/edit/387156
- https://tieto-si.visualstudio.com/360/_workitems/edit/390420

**Note to us:** ==Candidate for external release webinar==

---

### Document Templates Access via SIF

**What changed:**

- New SIF API methods enable external applications and microservices to discover and download 360 document templates with complete metadata
- Two new methods introduced:
  - `SupportService/GetDocumentTemplatesMetaData` - Returns extensive template metadata (description, dispatch type, format, scripts, categories, etc.)
  - `SupportService/DownloadDocumentTemplate` - Downloads the actual template file as a stream
- Applications can query available templates and retrieve both template files and configuration details

**Why it matters:**

This opens up document template access beyond the 360 user interface, enabling:
- Microservices to generate documents using 360 templates (like eMeetings Live for meeting protocols)
- External applications to maintain consistency with your 360 document standards
- Automated document production workflows that leverage your existing template library

Integration partners and custom applications can now seamlessly use your organization's approved templates without manual intervention.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/382954
- https://tieto-si.visualstudio.com/360/_workitems/edit/386723

**Note to us:** ==Candidate for external release webinar==

---

### SIFEvents Message Handling Improvements

**What changed:**

- Technical upgrade to the SIFEvents message acknowledgement system to align with latest RabbitMQ version requirements
- Refactored to ensure messages are acknowledged on the correct communication channel with proper async handling
- All methods are now async
- Channels are properly disposed after sessions
- Acknowledgement happens on the channel that originally consumed the message

**Why it matters:**

This enhancement ensures:
- More reliable message delivery for event subscriptions
- Better stability and performance for applications consuming 360 events
- Future-proof compatibility with modern messaging infrastructure
- Reduced risk of message loss or processing errors

Applications relying on 360 events for real-time updates and integrations will benefit from improved robustness and reliability.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/390094

**Note to us:** ==Internal webinar only - technical improvement==

---

## Microsoft Teams Integration Enhancements

### Teams File Import Progress Visibility

**What changed:**

- Enhanced user experience when importing files from Public 360 into Microsoft Teams channels
- System now provides real-time progress updates during large file imports
- Eliminated confusing timeout errors that appeared even when imports succeeded
- Users can see what's happening behind the scenes during long-running imports

**Why it matters:**

Previously, importing large amounts of data would show timeout errors even though the import was completing successfully in the background. Users had no way to know their import was actually working, leading to confusion and duplicate import attempts.

With this improvement:
- Users see real-time progress during file imports
- No more misleading timeout errors for long-running imports
- Clear visibility into what's happening behind the scenes
- Better confidence that imports are completing successfully

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/371860

**Note to us:** ==Candidate for external release webinar==

---

### Teams Import Job Management Improvements

**What changed:**

- Backend enhancements to the Teams file import system
- Improved data models for temporary storage and dependency queue management
- Import jobs now appear correctly in the dependency queue for monitoring and troubleshooting
- Fixed faulty data that prevented dependency queue from showing import jobs

**Why it matters:**

System administrators and support teams can now:
- Track asynchronous import jobs through the dependency queue
- Monitor and troubleshoot import operations more effectively
- Better understand system workload and job status
- Resolve issues faster with improved visibility into background processes

This creates a more robust and manageable Teams integration infrastructure.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/386175
- https://tieto-si.visualstudio.com/360/_workitems/edit/390277

**Note to us:** ==Internal webinar only - technical improvement==

---

### Improved Meeting Archiving Error Handling

**What changed:**

- Better error handling and user messaging for two common meeting archiving scenarios:
  1. Attempting to archive old meetings (typically older than 90 days) that are no longer accessible via Microsoft Graph API
  2. External users (from other organizations) attempting to access meeting archiving features they don't have permissions for
- Graceful error handling with clear explanations
- User-friendly messaging instead of cryptic error codes

**Why it matters:**

Instead of cryptic error codes and technical messages, users now receive:
- Clear explanations of why a meeting can't be archived
- Helpful guidance on what actions they can take
- Better understanding of system limitations with external users
- Reduced confusion and support tickets

The system handles these edge cases gracefully, improving the overall user experience for meeting archiving.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/386185 (Meeting not found error)
- https://tieto-si.visualstudio.com/360/_workitems/edit/386188 (External users handling)

**Note to us:** ==Consider for external release webinar - improves user experience==

---

### Automatic Case Binding for Teams Archiving

**What changed:**

- When archiving documents through Teams import archive with "Force archiving" method enabled, system now automatically uses the case that's already bound to the folder
- Eliminates need to manually select case when folder already has case assignment
- Brings consistency with "Teams archiving only" method behavior

**Why it matters:**

This streamlines the archiving workflow by:
- Reducing redundant data entry when cases are already defined at the folder level
- Ensuring consistency between folder and document case assignments
- Speeding up the archiving process for users
- Matching the behavior users already experience with "Teams archiving only" method

Users save time and avoid potential errors from selecting the wrong case manually.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/388675

**Note to us:** ==Candidate for external release webinar==

---

## FIKS IO Archive Integration

### External Export using FIKS IO Archive Client (Pilot Release - Closed Cases)

**What changed:**

- Pilot release enabling automated export of **closed cases** to external archive systems via FIKS IO Archive protocol
- Automated archiving job runs every 3 minutes checking eligibility based on configurable criteria:
  - SubArchive / Series
  - Case status
  - Document archive
  - Transfer timing (immediate or delayed after last update)
  - Code table mappings with P360/FIKS Schema
  - Queue capacity management (default: 20 items)
- Transferred metadata for cases, documents, and files:
  - **Case metadata** (Saksmappe message):
     - Case type (CL default, CC selectable), titles (official and unofficial)
     - Case status, sub-archive, journal unit
     - Case number (sequence number + year)
     - Organizational unit, responsible person (name, identifier, email)
     - Classification codes (Noark class codes), access codes and paragraphs
     - Case contacts, case references, remarks
  - **Document metadata** (Journalpost message):
     - Titles (official and unofficial), journal type, journal status, document category, document archive
     - Dates: document date, received date, created date, dispatched date, due date, journal date
     - Document sequence number, responsible person
     - Access codes, paragraphs, remarks, document contacts
  - **File metadata and content** (Dokumentbeskrivelse):
     - File content (active version only)
     - Title, sequence number, status, file category
     - Access codes, paragraphs, degrade dates and degrade codes
     - Format, variant format (PR, A), version number
     - Special handling: Proposition documents transfer production format for archive editing
- Update functionality for archived cases and documents (behind feature toggles):
  - **Case updates** (feature toggle "FIKSArchiveEnableCaseUpdate", default enabled): Case status, access code and paragraph, case titles (official and unofficial), case references (add only), classification codes (add and delete)
  - **Document and file updates** (feature toggle "FIKSArchiveEnableDocFileUpdate", default disabled): Document status, document titles, access code and paragraph
  - File access codes don't update when document access codes change
- Secure message exchange via FIKS IO (RabbitMQ) with Maskinporten authentication
- Complete monitoring through internal dependency queue showing request/response for each transfer
- Acknowledgment receipts from archiving system for success or failure
- Support extended to 360 eArchive installation type (in addition to existing Plan&Build and Public360 support)
- Requires separate FIKS IO Account IDs for client and server - same system cannot function as both
- Current limitations:
  - Only one default value set supported at server end
  - Access groups not supported (KS limitation)

**Why it matters:**

This standardized archiving protocol, developed by KS for Norwegian municipalities, enables:
- Seamless integration with external archiving systems without custom development
- Automated, secure case and document archival with encrypted, authenticated communication
- Complete audit trail through dependency queue monitoring
- Flexible configuration to match organizational archiving policies
- Compliance with Norwegian public sector archiving standards
- Reduced manual archiving work for closed cases
- Support across all 360 installation types (Public360, Plan&Build, and eArchive)
- Ongoing case management with ability to update archived metadata as cases progress

Organizations can now automatically transfer closed cases to long-term archiving systems while maintaining full visibility into the archival process. The pilot focuses on closed cases to ensure data stability during transfer. Updates to case status, access codes, titles, and classification codes automatically sync to the archive, ensuring consistency between active and archived data. Special handling for proposition documents preserves editing capability in the archive by transferring production format variants.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/354580 (External Export)
- https://tieto-si.visualstudio.com/360/_workitems/edit/384776 (eArchive support)

**Note to us:** ==Candidate for external release webinar - major integration feature==

---

## AutoSaver Enhancements

### AutoSaver Schema Mapping Experience (Pilot Feature)

**What changed:**

- Introduced new visual mapping experience for AutoSaver schema configuration in 360 web admin under Tools → AutoSaver schema mapping
- Pilot release behind feature toggle `EnableAutoSaverSchemaMappingUI` (default: false)
- Automatic background migration of existing schemas to enable UI-based editing
- Schema export functionality downloads complete mapping configuration as JSON to local computer
- Three-step schema creation workflow:
  - **Create schema mapping**: Configure schema code, description, schema type (XML or JSON), and upload schema content
  - **Fields mapping**: One-to-one node mapping from external schema to 360 parameters (example: `title` → `CreateDocumentParameter.Title`)
  - **Additional configuration**: Set default values for case mapping (case type, status, value sets, access groups, codes, paragraphs), document mapping (category, archive, responsible party, value sets), and search mapping (search filters for case connections)
- Designed to cater to non-technical users
- Target user roles: Public 360 Admin and Technical Administrator
- Built to help Professional Services consultants and customer technical administrators create easy, repeatable setups

**Why it matters:**

AutoSaver is a powerful framework for automating archiving from external systems, but despite positive feedback from users, adoption remains low. The main barrier: complex setup requiring technical expertise and manual JSON/XML editing of schema mappings.

This new mapping experience addresses that challenge directly:

- **Accessible to non-technical users**: Form-based interface replaces manual code editing, making AutoSaver configuration possible for administrators without development backgrounds
- **Reduced human error**: Guided workflows and visual mapping reduce mistakes in complex configuration scenarios
- **Repeatable configurations**: Export schemas as JSON to reuse across different customers with similar needs, the same customer with multiple processes, or between test and production environments
- **Faster deployment**: Consultants can configure AutoSaver integrations more quickly by building on proven templates rather than starting from scratch each time
- **Clear mapping logic**: Visual representation of how external data flows into 360 fields makes configuration easier to understand and validate

For consultants, this means creating one solid configuration and adapting it for multiple similar deployments. For customers, it means gaining independence to maintain and adjust their own AutoSaver setups without deep technical knowledge.

As a pilot feature, early adopters can test the workflow and provide feedback before general release.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/358550 (Epic)
- https://tieto-si.visualstudio.com/360/_workitems/edit/358551 (Implementation)

**Note to us:** ==Candidate for external release webinar - addresses major adoption barrier==

---

## FIKS Integration Quality Improvements

**What changed:**

- Fixed spelling error in CodeTableQueueowner for FIKS Saksfaser (corrected "FiIKS" to "FIKS")
- Added missing CodeTableQueueaction entry for FIKS Saksfaser
- Norwegian language support for code tables displayed in admin UI
- Language support for FIKS job start messages
- Improved localization throughout FIKS configuration interface

**Why it matters:**

These improvements ensure:
- Correct Norwegian terminology throughout the FIKS admin interface
- Complete and accurate queue configuration options
- Better user experience for Norwegian-speaking administrators
- Professional presentation of FIKS functionality
- Consistency in terminology across the system

Users see properly localized, error-free text throughout the FIKS configuration and monitoring interfaces, improving the professional appearance and usability of the system.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/367333

**Note to us:** ==Quality improvement - consider for internal notes only==

---

## Folkeregisteret Integration Enhancements

### Postal Address Support in Infotorg Lookup

**What changed:**

- Added postal address support in Infotorg lookup service, now enabled for all customers
- Introduced status code mapping using the "Contact Status" code table with Feature field set to "InfotorgPerson"
- Extended support to all address types available from the Infotorg service
- When status is enabled in the code table, the address field displays the status description instead

**Why it matters:**

You'll get more accurate and complete contact information from Folkeregisteret. The postal address support means you can now retrieve full address details during lookups, not just residence addresses. The new status code mapping gives you better visibility into special address situations, with clear status descriptions replacing addresses when appropriate.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/394571
- https://tieto-si.visualstudio.com/360/_workitems/edit/378734

**Note to us:** ==Candidate for external release webinar==

---

### Folkeregisteret Update (Vask) Improvements

**What changed:**

- Fixed bugs in the synchronization process
- Updated status text handling to use code table configuration
- Changed how contact statuses are processed and stored
- Fixed issue where postal address was incorrectly set for emigrated contacts

**Why it matters:**

Expect a higher volume of updates during the first synchronization after upgrading. This is normal and expected behavior, not an error. The fixes ensure more accurate contact information going forward, and the code table integration gives you better control over status handling.

**DevOps stories:**
- https://tieto-si.visualstudio.com/360/_workitems/edit/373777
- https://tieto-si.visualstudio.com/360/_workitems/edit/383027
- https://tieto-si.visualstudio.com/360/_workitems/edit/386692

**Note to us:** ==Internal webinar only - technical improvement==

---

### DPI and eFormidling Status Job Improvements

**What changed:**

- Enabled new business logic for DPI status job for all customers (previously only available to select customers)
- Reduced unnecessary logging in both DPI and eFormidling2 status jobs

**Why it matters:**

Your system logs will be cleaner and more focused on actionable information. The reduced logging improves performance and makes it easier to identify real issues when troubleshooting.

**DevOps stories:**
- [Work item IDs pending]

**Note to us:** ==Internal webinar only - technical improvement==

---
