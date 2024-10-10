# gitprint

## Flow

- Sign in with Github
- Enter a public repo URL
- Or select a private repo you have access to
- Download repo contents
- Analyize repo contents
- Apply default configurations
  - Languages to include
  - Always skipped files/directories
  - Skip tests
- Let user configure the export
- Apply order
  - README.md
  - root folder
  - chapter per folder
- Generate HTML file
  - apply syntax highlighting
- Generate PDF file from HTML file - https://gotenberg.dev/docs/routes#html-file-into-pdf-route
- Create a secure downloadable link

Download repo contents flow:

- Downlaod all contents of the repo into owner/repo/ref/...

## TODO

Pre Release TODO:
- Logo, Favicon and OG Image
- 3 top examples to download: Neovim...
- per user download/generate rate limit

Post Release TODO:
- global IP rate limit

## Additional information

Domain name: gitprint.me
Github OAuth App: https://github.com/settings/applications/2718171
