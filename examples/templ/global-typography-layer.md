# Global Typography Layer
Define typography globally when the same treatment is app-wide, applies to document-level content, or must also style non-templ consumers such as markdown, CMS content, or shared prose blocks.

## Pattern
```css
.prose-content {
  color: var(--text-strong);
}

.prose-content h1,
.prose-content h2 {
  font-family: var(--font-display);
  font-weight: 700;
  letter-spacing: 0.04em;
}

.prose-content p {
  line-height: 1.7;
}
```

```templ
package components

import "github.com/a-h/templ"

templ ArticleBody(content templ.Component) {
  <article class="prose-content">
    @content
  </article>
}
```

The pattern uses a global typography layer because the content tree can come from markdown, CMS, or shared prose outside one templ component family.

## Anti-pattern
```templ
package components

import "github.com/a-h/templ"

css articleRoot() {
  color: var(--text-strong);
}

css articleHeading() {
  font-family: var(--font-display);
  font-weight: 700;
  letter-spacing: 0.04em;
}

css articleParagraph() {
  line-height: 1.7;
}

templ ArticleBody(content templ.Component) {
  <article class={ articleRoot() }>
    @content
  </article>
}
```

The anti-pattern keeps typography local to one templ component even though the inner headings and paragraphs can come from non-templ content that cannot receive those local classes.
