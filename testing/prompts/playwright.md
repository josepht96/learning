I want to set up Playwright end-to-end tests for this app, copying a setup I've used before. Read the repo first, then come back with what you found and your plan before writing anything.

## Environment
- I work in WSL2. Docker runs there; nothing else is installed on the host (no node). Run every tool in a container.
- I open the app from a Windows browser. Work out where the app actually runs (WSL Docker, natively in WSL, or on Windows) and how a container in WSL reaches it: compose networking, host.docker.internal, or WSL's mirrored networking. Tell me what you find.

## How the setup should look
Everything lives in one `e2e/` folder:
- **Dockerfile** based on `mcr.microsoft.com/playwright:vX.Y.Z-noble`. Its tag must match the `@playwright/test` version pinned in `e2e/package.json`. Install dependencies into `/node_modules`, outside `/app`, so `e2e/` can be bind-mounted over `/app` for live editing.
- **Same URLs as my browser.** If the app is in compose, the browser in the container must use the same URLs I do (e.g. `localhost:<port>`). CORS and OAuth redirect URIs depend on the exact origin. Use socat in the entrypoint to forward those localhost ports to the compose services.
- **`e2e/compose.yml`**, layered on the app's compose file, with an `e2e` profile:
  - `ipc: host`, because Chromium crashes in the default 64MB `/dev/shm`
  - `HOME=/tmp`, because the container runs as my uid
  - results mounted to `logs/e2e/<env>/` on the host
- **Make targets** (or whatever task runner the repo uses), all run as my user so output isn't root-owned:
  - `e2e-test`: headless; `ARGS=` passes Playwright options
  - `e2e-ui`: Playwright UI mode, shown on my Windows desktop through WSLg (pass `DISPLAY` and mount `/tmp/.X11-unix`)
  - `e2e-codegen`: the recorder, on the chosen environment
  - `e2e-shell`
- **Environments.** `ENV=local|<shared test env>|prod` picks the base URL from a `targets.ts` map; local is the default. The project in the config is named after the env, so reports show where tests ran.
- **Prod is read-only:**
  - The config sets `grep: /@readonly/` when the target is prod, so a test only runs there if it's tagged.
  - A `onlyOn('local')` helper skips tests that change data everywhere else.
  - Shared environments only get tests that leave no trace.
- **Fixtures** (`fixtures.ts`): an auto `problems` fixture collects page errors and failed API responses. Leave out expected failures (e.g. the 401 a signed-out "who am I" call returns). Tests that load a page assert it's empty.
- **Failures:** keep a trace, screenshot and video only for failing tests, and write an HTML report that doesn't auto-open.
- A short `e2e/README.md` covering how to run it, plus a table of what's covered.

## OAuth sign-in
The app signs in with OAuth. Find out how: the provider, the flow, and whether the session is a cookie or a token in storage. Then propose one of these:
1. **Real provider login, once per run**, in a Playwright `setup` project. It saves `storageState` to a gitignored `e2e/.auth/<env>-<role>.json`, and the other projects depend on it. This only works with dedicated test accounts that don't use MFA or CAPTCHA.
2. **Get a token directly**, through a test grant on the provider (client credentials, ROPC on a test tenant, etc.). It then gets injected as a header or into storage.
3. **A mock identity provider for local only** (e.g. `mock-oauth2-server` in compose), if the app's OAuth config can point at it.

Rules for sign-in:
- Don't work around MFA or other security controls on real accounts.
- Credentials come only from environment variables. If they aren't set, the test skips with a message saying which variables are needed.
- Never commit or log secrets or the `.auth/` files.
- Support one stored session per role if the app has roles.

## Writing the tests
- **Look at the app first:**
  - Read the routes and the API.
  - Visit each environment in a headless browser to see what's really there; don't assume.
  - Pages I gave you that need a sign-in, visit only after we've sorted out authentication.
- **Coverage.** Group specs by area. Cover:
  - every page loads cleanly (no page errors or failed requests)
  - navigation and layout
  - forms: validation, plus a successful submit (stub the submit with `page.route` outside local)
  - authenticated flows
  - API contract checks, through the `request` fixture
  - mobile width
  - basic accessibility (labels, landmarks, skip link)
- **Locators:** use role and label locators, no CSS selectors. No fixed sleeps; use `expect.poll` and auto-waiting.
- Use unique stamped names for anything a test creates.
- **Run everything against each environment.** When a test fails, decide whether the test or the app is wrong. Fix the test, or list it as an app finding for me; don't hide real bugs.

## Ground rules
- Don't push.
- Don't change the app's code or infrastructure without asking.
- Don't run anything against prod except `@readonly` tests.
- Ask me before acting on anything that hits a shared environment beyond reading it.
- Keep it simple: one way to run each thing, everything through Docker and Make.

Before you use it:
- Shared environment names: it says <shared test env> because I don't know what your work's environments are called. Fill in the real names, or leave it for Claude to find.
- OAuth test accounts: option 1 only works if your identity team will give you test accounts without MFA. If they won't, a token grant (option 2) or a mock provider for local (option 3) is the practical route.
- Copying OTO's files: OTO's e2e/ is gitignored, so only this machine has it. If the work repo is on the same machine, you could add "Use ~/repos/our-town-originals/e2e/ as the reference implementation." Claude can then copy the shape straight from those files.