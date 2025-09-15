package cli

const INFO = `Node Runtime Environment
  - Version 0.1.0
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

const REQ = `Requirements:
  NRE requires NVM to be installed and
  the "NVM_HOME" environment variable to be set.

  Check if NVM_HOME is not empty:
    1. Open CMD (Windows Terminal)
    2. Run command:
         echo %NVM_HOME%
    3. Check if the path is displayed.

  Node Version Manager:
    - for Windows ... https://github.com/coreybutler/nvm-windows`
