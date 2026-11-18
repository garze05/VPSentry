# Instructions for coding agents

## Keep the README accurate

- Read `README.md` before changing code.
- For every code change, check whether the README needs an update. Update it in the same change when affected; documentation is part of completing the work.
- Review changes to CLI flags, defaults, prompts, output, error handling, prerequisites, privileges, supported platforms, build commands, testing, project structure, and security behavior.
- Keep implemented capabilities, known limitations, and planned features clearly separated. When a feature is implemented or a documented limitation is fixed, update the corresponding status and roadmap entries.
- Write README content in English. Keep examples consistent with the actual code.
- Do not claim production readiness, verified access, successful tests, or security guarantees without evidence.
- In the final summary, state which documentation was updated, or explain briefly why no README update was needed.

## Scope and safety

- Preserve user changes and keep edits focused on the requested task.
- Never publish credentials, private keys, local environment files, or scratch artifacts.
- Use a disposable Ubuntu VM for tests that modify users, SSH, firewall rules, packages, or services. Do not make privileged changes to the development host.
- Do not commit or push changes unless the user requests it.
