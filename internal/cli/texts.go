package cli

const INFO = `Node Runtime Environment
  - Version 0.2.0
  - See https://github.com/Adrosar/nre
  - Go to help "nre help"`

const HELP = `Main command:
  > nre <NODE_VERSION> <COMMAND>

Examples:
  > nre 8 node -v
  > nre 10 npm run build
  > nre 12 gulp --env dev
  > nre 14 git commit -m "lorem ipsum"

Additional commands:
  > nre help ...... see help
  > nre www ....... go to WWW
  > nre req ....... requirements
  > nre list ...... see all NodeJS versions`

const REQ = `NRE requires NVM:
  - for Windows ... https://github.com/coreybutler/nvm-windows
  - for Linux ..... https://github.com/nvm-sh/nvm`
