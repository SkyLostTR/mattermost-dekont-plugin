# GitHub Pages Setup Guide

This guide will help you enable GitHub Pages for your Mattermost PDF Dekont Parser Plugin documentation.

## 🚀 Quick Setup

### 1. Enable GitHub Pages

1. Go to your repository on GitHub
2. Click on **Settings** tab
3. Scroll down to **Pages** section
4. Under **Source**, select **GitHub Actions**
5. Save the changes

### 2. Commit and Push

```bash
# Add all documentation files
git add docs/

# Commit changes
git commit -m "docs: add GitHub Pages documentation website"

# Push to main branch
git push origin main
```

### 3. Verify Deployment

1. Go to **Actions** tab in your repository
2. Wait for the "Deploy GitHub Pages" workflow to complete
3. Visit your documentation site at: `https://YOUR_USERNAME.github.io/mattermost-dekont-plugin`

## 🔧 Configuration Options

### Custom Domain (Optional)

If you have a custom domain:

1. Create a `CNAME` file in the `docs/` directory:
   ```
   docs.yourdomain.com
   ```

2. Configure DNS records:
   - Add a CNAME record pointing to `YOUR_USERNAME.github.io`

3. Update `_config.yml`:
   ```yaml
   url: "https://docs.yourdomain.com"
   baseurl: ""
   ```

### Google Analytics (Optional)

To add Google Analytics tracking:

1. Get your GA tracking ID
2. Add to `_config.yml`:
   ```yaml
   google_analytics: GA_TRACKING_ID
   ```

### SEO Improvements

The site already includes:
- ✅ Meta tags and Open Graph
- ✅ Structured data
- ✅ Sitemap generation
- ✅ robots.txt
- ✅ Performance optimizations

## 📝 Content Management

### Adding New Pages

1. Create a new `.md` file in `docs/`
2. Add front matter:
   ```yaml
   ---
   layout: default
   title: Page Title
   nav_order: 5
   description: "Page description"
   ---
   ```
3. Add content using Markdown
4. Commit and push changes

### Updating Navigation

Edit `_config.yml` and update the `header_pages` list:
```yaml
header_pages:
  - index.md
  - installation.md
  - your-new-page.md
```

### Styling Changes

Edit `docs/assets/custom.css` to customize:
- Colors and themes
- Typography
- Layout and spacing
- Component styles

## 🔍 Testing Locally

```bash
# Install dependencies
cd docs
bundle install

# Serve locally
bundle exec jekyll serve

# Open http://localhost:4000/mattermost-dekont-plugin
```

## 🐛 Troubleshooting

### Build Failures

Check the Actions tab for error details. Common issues:
- YAML syntax errors in `_config.yml`
- Missing dependencies in Gemfile
- Broken internal links

### Content Not Updating

- Clear browser cache
- Check if changes were committed to main branch
- Verify GitHub Actions completed successfully

### Layout Issues

- Test locally before pushing
- Check CSS syntax
- Verify responsive design on mobile

## 📊 Analytics and Monitoring

### GitHub Insights

Monitor your documentation:
- **Traffic**: Repository insights → Traffic
- **Popular Content**: Page views and referrers
- **User Engagement**: Bounce rate and session duration

### Performance

The site is optimized for:
- Fast loading times (< 3 seconds)
- Mobile responsiveness
- Search engine indexing
- Accessibility standards

## 🆕 Updates and Maintenance

### Regular Updates

- Keep dependencies updated
- Monitor for security alerts
- Update content for new releases
- Review and improve based on user feedback

### Automation

The setup includes automated:
- ✅ Dependency updates
- ✅ Security scanning
- ✅ Build and deployment
- ✅ SEO optimization

## 📚 Resources

- [Jekyll Documentation](https://jekyllrb.com/)
- [GitHub Pages Docs](https://docs.github.com/en/pages)
- [Markdown Guide](https://www.markdownguide.org/)
- [Jekyll Themes](https://jekyllrb.com/docs/themes/)

---

Your documentation website is now ready! Visit it at:
**https://skylosttr.github.io/mattermost-dekont-plugin**
