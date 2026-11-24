# Praemunio Go SDK

> Client officiel pour l'API Praemunio — Go

## 📦 Installation

```bash
go get github.com/idcoopp-praemunio/praemunio-go-sdk
```

Ou ajoutez à votre `go.mod` :

```go
require github.com/idcoopp-praemunio/praemunio-go-sdk v1.0.0
```

## 🚀 Quick Start (5 minutes)

### 1. Obtenir vos clés API

Récupérez vos clés depuis le [dashboard Praemunio](https://dashboard.praemunio.io) :
- **Public Key** : `pk_test_...` (sandbox) ou `pk_live_...` (production)
- **Secret Key** : `sk_test_...` (sandbox) ou `sk_live_...` (production)

### 2. Initialiser le client

```go
package main

import (
    "fmt"
    "log"

    praemunio "github.com/idcoopp-praemunio/praemunio-go-sdk"
)

func main() {
    // Mode sandbox (tests)
    client := praemunio.NewPraemunioClientSandbox("pk_test_xxx", "sk_test_xxx")

    // Mode production
    // client := praemunio.NewPraemunioClientLive("pk_live_xxx", "sk_live_xxx")
}
```

### 3. S'authentifier

```go
if err := client.Authenticate(); err != nil {
    log.Fatal(err)
}

fmt.Println("Access token:", client.GetAccessToken())
fmt.Println("Refresh token:", client.GetRefreshToken())
```

### 4. Premier appel API

```go
api := client.CreateAPI()

// Lister les workspaces
result, err := api.WorkspacesDriveGet()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Workspaces:", result)
```

### 5. Exemple complet

```go
package main

import (
    "fmt"
    "log"

    praemunio "github.com/idcoopp-praemunio/praemunio-go-sdk"
)

func main() {
    // Initialiser le client (sandbox)
    client := praemunio.NewPraemunioClientSandbox("pk_test_xxx", "sk_test_xxx")

    // Authentification
    if err := client.Authenticate(); err != nil {
        log.Fatal("Authentication failed:", err)
    }
    fmt.Println("Authentifié avec succès !")

    // Premier appel API
    api := client.CreateAPI()
    workspaces, err := api.WorkspacesDriveGet()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Workspaces:", workspaces)
}
```

## 🔐 Authentification

Le SDK gère automatiquement l'authentification et le refresh des tokens :

```go
// Authentification initiale
err := client.Authenticate()
if err != nil {
    log.Fatal(err)
}

// Le token est renouvelé automatiquement lors des appels API
// Vous pouvez aussi forcer le refresh :
err = client.Refresh()
```

## 🌍 Environnements

| Environnement | URL | Utilisation |
|---------------|-----|-------------|
| **Sandbox** | `https://sandbox-api.praemunio.org/api/v1` | Tests et développement |
| **Live** | `https://api.praemunio.org/api/v1` | Production |

```go
// Sandbox (par défaut)
client := praemunio.NewPraemunioClientSandbox(publicKey, secretKey);

// Production
client := praemunio.NewPraemunioClientLive(publicKey, secretKey);
```

## 📚 Documentation complète

- [Documentation officielle Go](https://docs.praemunio.io/sdk/go)
- [Référence API](https://docs.praemunio.io/api-reference)
- [Guide d'authentification](https://docs.praemunio.io/authentication)
- [Exemples](https://docs.praemunio.io/sdk/go/examples)

## 📋 Endpoints principaux

| Méthode | Endpoint | Description |
|---------|----------|-------------|
| `POST` | `/auth/auth-token` | Authentification |
| `GET` | `/workspaces/drive` | Lister les workspaces |
| `POST` | `/teams/create` | Créer une équipe |
| `GET` | `/folders/{id}` | Détails d'un dossier |
| `POST` | `/documents/archive` | Archiver un document |

## 🤝 Support

- 📖 [Documentation](https://docs.praemunio.io)
- 🐛 [Signaler un bug](https://github.com/praemunio/sdks/issues)
- 💬 [Discord](https://discord.gg/praemunio)
- 📧 [Email](mailto:support@praemunio.org)