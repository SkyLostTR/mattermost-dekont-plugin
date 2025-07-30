# Contributing to Mattermost PDF Dekont Parser Plugin

We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests
3. If you've changed APIs, update the documentation
4. Ensure the test suite passes
5. Make sure your code lints
6. Issue that pull request!

### Code Style

* Use `gofmt` for formatting
* Follow Go naming conventions
* Write clear, self-documenting code
* Add comments for complex logic
* Keep functions small and focused

### Commit Messages

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

Examples:
- `feat: add support for multiple bank formats`
- `fix: handle PDF parsing errors gracefully`
- `docs: update installation instructions`

## Setting Up Development Environment

1. **Prerequisites**:
   - Go 1.21+
   - Git
   - VS Code (recommended)

2. **Clone and setup**:
   ```bash
   git clone https://github.com/SkyLostTR/mattermost-dekont-plugin.git
   cd mattermost-dekont-plugin
   go mod download
   ```

3. **Build and test**:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o plugin
   go test ./...
   ```

## Testing

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...
```

### Writing Tests
- Write unit tests for all new functions
- Use table-driven tests where appropriate
- Mock external dependencies
- Test error conditions
- Aim for >80% test coverage

### Test Structure
```go
func TestExtractFields(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "valid bank receipt",
            input:    "ALICI: John Doe\nAÇIKLAMA: Payment\nTUTARI: 100.00 TL",
            expected: "**Açıklama**: Payment\n**Alıcı**: John Doe\n**İşlem Tutarı**: 100.00 TL",
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := extractFields(tt.input)
            if result != tt.expected {
                t.Errorf("extractFields() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## Code Review Process

### For Contributors
1. Ensure all tests pass
2. Update documentation if needed
3. Follow the coding standards
4. Write clear commit messages
5. Respond to review feedback promptly

### For Reviewers
1. Check code functionality and logic
2. Verify test coverage
3. Review documentation updates
4. Ensure coding standards compliance
5. Test the changes locally if possible

## Security

### Reporting Security Issues
Please do not report security vulnerabilities through public GitHub issues. Instead, send an email to [security contact].

### Security Guidelines
- Never commit secrets or API keys
- Validate all user inputs
- Use secure coding practices
- Keep dependencies up to date
- Follow the principle of least privilege

## Documentation

### Required Documentation
- Update README.md for user-facing changes
- Add inline code comments for complex logic
- Update API documentation for interface changes
- Include examples for new features

### Documentation Style
- Write clear, concise explanations
- Use proper Markdown formatting
- Include code examples where helpful
- Keep documentation up to date with code changes

## Issue Reporting

### Bug Reports
Include:
- Clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Environment details (OS, Go version, Mattermost version)
- Relevant logs or error messages

### Feature Requests
Include:
- Clear description of the feature
- Use case and motivation
- Proposed implementation approach
- Any alternative solutions considered

## Release Process

1. Update version in `plugin.json`
2. Update CHANGELOG.md
3. Create release tag
4. GitHub Actions will automatically build and publish

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

## Questions?

Feel free to open an issue for questions or reach out to the maintainers.

---

Thank you for contributing! 🎉
