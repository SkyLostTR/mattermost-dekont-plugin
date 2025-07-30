# Mattermost PDF Dekont Parser Plugin Documentation

This directory contains the GitHub Pages documentation website for the Mattermost PDF Dekont Parser Plugin.

## 🌐 Live Website

Visit the documentation website at: [https://skylosttr.github.io/mattermost-dekont-plugin](https://skylosttr.github.io/mattermost-dekont-plugin)

## 📁 Documentation Structure

```
docs/
├── _config.yml                 # Jekyll configuration
├── _layouts/
│   └── default.html           # Custom page layout
├── assets/
│   └── custom.css             # Custom styling
├── index.md                   # Homepage
├── installation.md            # Installation guide
├── usage.md                   # Usage instructions
├── banks.md                   # Supported banks
├── development.md             # Developer guide
├── api.md                     # API reference
├── contributing.md            # Contributing guidelines
├── changelog.md               # Release history
└── README.md                  # This file
```

## 🚀 Local Development

To run the documentation site locally:

### Prerequisites
- Ruby 3.1+
- Bundler gem

### Setup
```bash
# Navigate to docs directory
cd docs

# Install dependencies
bundle install

# Serve locally
bundle exec jekyll serve

# Open http://localhost:4000/mattermost-dekont-plugin
```

### Local Development with Docker
```bash
# From project root
docker run --rm \
  --volume="$PWD/docs:/srv/jekyll:Z" \
  --volume="$PWD/docs/_site:/srv/jekyll/_site:Z" \
  --publish 4000:4000 \
  jekyll/jekyll:latest \
  jekyll serve --watch --force_polling
```

## 📝 Content Guidelines

### Page Structure
Each documentation page should include:
- Front matter with title, nav_order, and description
- Table of contents (for longer pages)
- Clear headings and subheadings
- Code examples where applicable
- Links to related pages

### Front Matter Example
```yaml
---
layout: default
title: Page Title
nav_order: 1
description: "Brief description for SEO"
---
```

### Markdown Standards
- Use descriptive headings
- Include code examples with syntax highlighting
- Add tables for structured information
- Use callout boxes for important information
- Include emoji icons for visual appeal

### Navigation Order
Pages are ordered by the `nav_order` field:
1. Home (index.md)
2. Installation Guide
3. Usage Guide
4. Supported Banks
5. Development Guide
6. API Reference
7. Contributing
8. Changelog

## 🎨 Styling and Design

### Custom CSS
- Located in `assets/custom.css`
- Extends the default Minima theme
- Includes custom components and utilities
- Responsive design for mobile devices

### Color Scheme
- Primary: #1e74fd (Mattermost blue)
- Secondary: #28a745 (Success green)
- Warning: #ffc107 (Warning yellow)
- Danger: #dc3545 (Error red)

### Components
- Custom buttons with hover effects
- Alert boxes for important information
- Feature cards for highlighting capabilities
- Status badges for bank support levels
- Navigation breadcrumbs and page links

## 🔧 GitHub Pages Configuration

### Automatic Deployment
- Triggered on pushes to `main` branch
- Uses Jekyll for static site generation
- Deployed via GitHub Actions workflow
- Available at custom domain (if configured)

### Workflow File
See `.github/workflows/pages.yml` for the complete CI/CD pipeline.

### Build Process
1. Checkout repository code
2. Setup Ruby and Jekyll environment
3. Install dependencies via Bundler
4. Build static site with Jekyll
5. Deploy to GitHub Pages

## 📊 SEO and Analytics

### Search Engine Optimization
- Structured meta tags
- Open Graph tags for social sharing
- Twitter Card support
- Sitemap generation
- SEO-friendly URLs

### Performance
- Optimized images and assets
- Minified CSS and JavaScript
- Fast loading times
- Mobile-responsive design

## 🤝 Contributing to Documentation

### Making Changes
1. Fork the repository
2. Create a feature branch
3. Edit documentation files
4. Test locally with Jekyll
5. Submit a pull request

### Content Guidelines
- Write clear, concise documentation
- Include practical examples
- Use proper grammar and spelling
- Follow the established style guide
- Add screenshots where helpful

### Review Process
- Documentation changes are reviewed by maintainers
- Ensure accuracy and completeness
- Verify links and references work
- Test local build before submitting

## 📋 Content Checklist

When adding new content:
- [ ] Add appropriate front matter
- [ ] Include in navigation if needed
- [ ] Add to table of contents
- [ ] Include relevant code examples
- [ ] Add links to related pages
- [ ] Test all links and references
- [ ] Verify mobile responsiveness
- [ ] Check spelling and grammar

## 🐛 Issues and Suggestions

- **Documentation Bugs**: [Report here](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues/new?template=documentation.md)
- **Content Suggestions**: [Discuss here](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions)
- **Website Issues**: Check the build status in GitHub Actions

## 📚 Additional Resources

- [Jekyll Documentation](https://jekyllrb.com/docs/)
- [GitHub Pages Documentation](https://docs.github.com/en/pages)
- [Markdown Guide](https://www.markdownguide.org/)
- [Mattermost Brand Guidelines](https://www.mattermost.org/brand-guidelines/)

---

For questions about the documentation website, please [open an issue](https://github.com/SkyLostTR/mattermost-dekont-plugin/issues) or [start a discussion](https://github.com/SkyLostTR/mattermost-dekont-plugin/discussions).
