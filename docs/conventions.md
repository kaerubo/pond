# Development Conventions

## 📂 Branch Naming Convention

Use the following pattern for naming branches:

```
<type>/<short-description>
```

### Supported Types

| Type      | Description              |
|-----------|--------------------------|
| feat      | New feature              |
| fix       | Bug fix                  |
| refactor  | Code refactoring         |
| docs      | Documentation            |
| chore     | Maintenance / tooling    |
| test      | Adding/modifying tests   |

### Example

- `feat/create-user-endpoint`
- `fix/auth-header-bug`

## Commit Message Convention

We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard.

### Format

```
<type>(<scope>): <short description>
```

### Example

```
feat(user): implement user registration API
```

### Common Types

| Type       | Use case                        |
|------------|---------------------------------|
| feat       | A new feature                   |
| fix        | A bug fix                       |
| refactor   | Code changes that aren’t fixes or features |
| docs       | Documentation changes           |
| chore      | Build process or tooling updates |
| test       | Adding or updating tests        |

---

## Commit Scope Reference

To provide better clarity in commit messages, we recommend using the following scopes based on the project's modular structure:

| Scope        | Area / Directory             | Example Usage                        |
|--------------|------------------------------|--------------------------------------|
| `kero`       | Core entity (`kero`)         | `feat(kero): add create use case`    |
| `auth`       | Authentication logic         | `fix(auth): fix refresh logic`       |
| `user`       | User-related logic           | `refactor(user): extract profile`    |
| `handler`    | HTTP handlers                | `feat(handler): add route`           |
| `service`    | Use case / Interactor layer  | `test(service): add unit test`       |
| `repo`       | Repository / DB access       | `fix(repo): update query logic`      |
| `middleware` | Middleware (e.g., JWT)       | `feat(middleware): add logger`       |
| `docs`       | Documentation                | `docs(docs): update instructions`    |
| `ci`         | CI / GitHub Actions          | `chore(ci): update workflow`         |
| `infra`      | Docker / infrastructure      | `chore(infra): update Dockerfile`    |
| `test`       | Testing (unit/integration)   | `test(test): add e2e case`           |
