package cli

const INFO = `Node Runtime Environment
  - Version 0.3.0-dev
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
  > nre list ...... see all NodeJS versions
  > nre link ...... create link "./node_modules/myapp"
                    to "/workspace/myapp"

Examples:
  > nre link "C:\workspace\my-app"  (Windows)
  > nre link "/workspace/my-app"    (Unix)`

const REQ = `NRE requires NVM:
  - for Windows ... https://github.com/coreybutler/nvm-windows
  - for Linux ..... https://github.com/nvm-sh/nvm`
