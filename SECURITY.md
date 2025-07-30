# Security Policy

## Supported Versions

We provide security updates for the following versions of the Mattermost PDF Dekont Parser Plugin:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |

## Reporting a Vulnerability

We take security seriously. If you discover a security vulnerability, please follow these guidelines:

### Do NOT:
- Open a public GitHub issue
- Disclose the vulnerability publicly
- Test the vulnerability on production systems

### DO:
1. **Email us privately** at [security contact email]
2. **Provide detailed information** including:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if available)
3. **Allow reasonable time** for us to investigate and fix the issue

## Security Response Process

1. **Acknowledgment**: We will acknowledge receipt of your report within 48 hours
2. **Investigation**: We will investigate the issue and determine its severity
3. **Fix Development**: We will develop and test a fix
4. **Disclosure**: We will work with you on responsible disclosure timing
5. **Release**: We will release a security patch and advisory

## Security Best Practices

### For Users
- Always download the plugin from official sources
- Keep the plugin updated to the latest version
- Monitor Mattermost server logs for unusual activity
- Restrict plugin installation to trusted administrators
- Review PDF files before uploading if they contain sensitive information

### For Developers
- Follow secure coding practices
- Validate all inputs
- Use parameterized queries/prepared statements
- Implement proper error handling that doesn't leak information
- Keep dependencies updated
- Run security scans regularly

## Security Features

### Input Validation
- PDF file type verification
- File size limits (handled by Mattermost)
- Content sanitization before processing

### Error Handling
- No sensitive information in error messages
- Proper logging without exposing user data
- Graceful failure modes

### Resource Management
- Temporary file cleanup
- Memory usage controls
- Process isolation

### Data Privacy
- No persistent storage of PDF content
- Minimal data extraction (only specified fields)
- No external network requests

## Dependencies Security

We regularly monitor our dependencies for security vulnerabilities:

### Go Dependencies
- `github.com/mattermost/mattermost-server/v6` - Official Mattermost framework
- `github.com/ledongthuc/pdf` - PDF processing library

### Security Scanning
- Automated dependency scanning via GitHub Actions
- Regular vulnerability assessments
- Automated security updates where possible

## Threat Model

### Assets
- PDF documents uploaded to Mattermost
- Extracted transaction data
- Mattermost server resources

### Threats
- Malicious PDF files
- Information disclosure
- Resource exhaustion
- Code injection

### Mitigations
- PDF parsing in isolated process
- Input validation and sanitization
- Resource limits and timeouts
- Minimal privilege requirements

## Security Contact

For security-related inquiries, please contact:
- **Email**: [security@example.com]
- **PGP Key**: [Key ID if available]

## Acknowledgments

We appreciate security researchers and users who help improve the security of our plugin. Contributors who responsibly disclose security issues will be acknowledged in our security advisories (with their permission).

---

**Note**: This security policy is subject to change. Please check back regularly for updates.
